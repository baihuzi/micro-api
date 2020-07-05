package router

import (
	"net/http"
	"review-server/core/http/controllers"
)

type Route struct {
	Path       string
	Controller controllers.ControllerInterface
	Action     string
	Method     string
	Options    map[string]interface{}
	Handler    func(http.ResponseWriter, *http.Request)
}

var routes []*Route

func Register(path string, c controllers.ControllerInterface, action, method string, o map[string]interface{}) {
	route := &Route{
		Path:       path,
		Controller: c,
		Action:     action,
		Method:     method,
		Options:    o,
	}
	route.Handler = GetRouteHandler().Create(route)

	routes = append(routes, route)
}

func GET(path string, c controllers.ControllerInterface, action string, o map[string]interface{}) {
	Register(path, c, action, "GET", o)
}

func POST(path string, c controllers.ControllerInterface, action string, o map[string]interface{}) {
	Register(path, c, action, "POST", o)
}

func PUT(path string, c controllers.ControllerInterface, action string, o map[string]interface{}) {
	Register(path, c, action, "PUT", o)
}

func DELETE(path string, c controllers.ControllerInterface, action string, o map[string]interface{}) {
	Register(path, c, action, "DELETE", o)
}

func REQUEST(path string, c controllers.ControllerInterface, action string, o map[string]interface{}) {
	Register(path, c, action, "REQUEST", o)
}

func GetAll() []*Route {
	return routes
}
