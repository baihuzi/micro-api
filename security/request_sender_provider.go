package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"regexp"
	"review-server/core"
)

const TOKEN_NAME = "re-token"

type RequestSenderProvider struct {
	pattern string
	secret  string
}

func (rp *RequestSenderProvider) GetCredentialsFromRequest(r *http.Request) (jwt string) {
	if cookie, err := r.Cookie(TOKEN_NAME); err == nil {
		jwt = cookie.Value
	} else if jwt = r.Header.Get(TOKEN_NAME); jwt != "" {

	} else if err = r.ParseForm(); err == nil {
		jwt = r.FormValue(TOKEN_NAME)
	} else {
		jwt = ""
	}
	return
}

func (r *RequestSenderProvider) IfPattern(path string) bool {
	ok, _ := regexp.MatchString(r.pattern, path)

	return ok
}

func (r *RequestSenderProvider) Authenticate(jwtStr string) error {
	if jwtStr == "" {
		return errors.New("jwt is empty")
	}
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(r.secret), nil
	})
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["u"])
	} else {
		return err
	}

	return nil
}

func NewRequestSenderProvider(pattern string) *RequestSenderProvider {
	return &RequestSenderProvider{
		pattern: pattern,
		secret:  core.GetConfig().App.ReLoginSecret}
}
