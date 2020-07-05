package middlewares

import "net/http"

type MiddlewareInterface interface {
	Before(r *http.Request)
	After(result interface{},w http.ResponseWriter, r *http.Request)
	GetBeforePriority() int8
	GetAfterPriority() int8
}
