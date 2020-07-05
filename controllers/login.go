package controllers

import (
	"net/http"
	"review-server/core/http/response"
	"review-server/security"
	"time"
)

type LoginController struct {
	ReController
}

type users map[string]string

func (u users) Get(key string) string {
	if value, ok := u[key]; ok {
		return value
	}
	return ""
}

func (l *LoginController) Login() (result *response.Result) {

	users := users{
		"root": "123456",
	}
	name := l.Request().Get("username", "string", "")
	password := l.Request().Get("password", "string", "")
	result = response.NewResult()
	if p := users.Get(name.(string)); p == "" || p != password {
		result.Status = response.StatusError
		result.Msg = "name or password error."
		result.StatusCode = http.StatusForbidden
		return
	}

	result.Data = struct {
		ReToken string `json:"re_token"`

	}{
		security.NewRequestTokenGenerator().Create(name.(string), time.Hour*7*24)}
	return
}
