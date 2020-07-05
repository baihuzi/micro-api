package request

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type RequestBodyInterface interface {
	Get(key string, valueType string, defaultValue interface{}) (value interface{})
}

type requestBody struct {
	Data     url.Values
	Query    url.Values
	FormData url.Values
	ParseFun func(*http.Request)
}

var rb *requestBody

func GetRequestBody() *requestBody {
	if rb == nil {
		rb = &requestBody{
			Data:  make(url.Values),
			Query: make(url.Values),
			ParseFun: func(request *http.Request) {
				rb.DefaultParse(request)
			},
		}
	}
	return rb
}

func (rb *requestBody) DefaultParse(request *http.Request) {
	rb.RemoveAll()
	rb.DefaultParseQuery(request)
	rb.DefaultarseBody(request)
}

func (rb *requestBody) DefaultParseQuery(request *http.Request) {
	rb.Query = request.URL.Query()
}

func (rb *requestBody) DefaultarseBody(request *http.Request) {
	request.ParseForm()
	request.ParseMultipartForm(1024)
	rb.Data = request.PostForm
	rb.FormData = request.Form
}

func (rb *requestBody) RemoveAll() {
	rb.Data = make(url.Values)
	rb.Query = make(url.Values)
	rb.FormData = make(url.Values)
}

func (rb *requestBody) Get(key string, valueType string, defaultValue interface{}) (value interface{}) {
	getValue := func(key string) (string, error) {
		if rb.FormData == nil {
			return "", errors.New("form data is nil")
		}
		vs := rb.FormData[key]
		if len(vs) == 0 {
			return "", errors.New("the key non-existent.")
		}
		return vs[0], nil

	}

	v, err := getValue(key)
	if err != nil {
		return defaultValue
	}

	switch valueType {
	case "int":
		value, err = strconv.Atoi(v)
		if err != nil {
			return defaultValue
		}
	case "array":
		value = rb.FormData[key]
	default:
		value = v
	}

	return
}

func (rb *requestBody) Merge(agr ...url.Values) (data url.Values) {
	data = make(url.Values)

	for _, values := range agr {
		for key, value := range values {
			data[key] = append(data[key], value...)
		}
	}

	return
}
