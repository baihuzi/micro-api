package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "review-server/core/http/middlewares"
	requestBody "review-server/core/http/request"
	"review-server/core/http/response"
	"review-server/core/http/router"
)

type ViewMiddleware struct {
}

func (v *ViewMiddleware) Before(r *http.Request) {
	//parse request data
	requestBody.GetRequestBody().ParseFun(r)
}

func (v *ViewMiddleware) After(result interface{}, writer http.ResponseWriter, r *http.Request) {
	res := result.(*response.Result)

	//api server ,content-type is json
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Max-Age", "86400")
	writer.Header().Add("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	writer.Header().Add("Access-Control-Allow-Headers","Origin, Content-Type, re-token")

	//res status code
	if res.StatusCode != http.StatusOK {
		//todo
		//writer.WriteHeader(res.StatusCode)
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusInternalServerError)
	} else {

		fmt.Fprintf(writer, string(json))
	}
}

func (v *ViewMiddleware) GetBeforePriority() int8 {
	return router.MIDDLEWARE_EARLY_EVENT
}

func (v *ViewMiddleware) GetAfterPriority() int8 {
	return router.MIDDLEWARE_LATE_EVENT
}

func NewViewMiddleware() *ViewMiddleware {
	return &ViewMiddleware{}
}
