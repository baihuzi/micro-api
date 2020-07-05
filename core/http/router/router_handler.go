package router

import (
	"net/http"
	"reflect"
	"review-server/core/http/response"
	"review-server/core/services"
	"sort"
	"sync"
)

const (
	LATE_EVENT             = 512
	EARLY_EVENT            = -512
	MIDDLEWARE_LATE_EVENT  = int8(^uint8(0) >> 1)   //127
	MIDDLEWARE_EARLY_EVENT = ^MIDDLEWARE_LATE_EVENT //-128
	MIDDLEWARE_EVENT_P     = int(^uint8(0))         //255
)

type objects struct {
	route  *Route
	r      *http.Request
	w      http.ResponseWriter
	result *response.Result
}

type event func(o *objects)

type RouteHandler struct {
	events map[int][]event
}

var ins *RouteHandler
var once sync.Once

func GetRouteHandler() *RouteHandler {
	once.Do(func() {
		ins = &RouteHandler{
			events: make(map[int][]event),
		}
		ins.methodAssert()
		ins.firewalls()
		ins.beforeMiddlebrows()
		ins.atController()
		ins.afterMiddlebrows()
	})

	return ins
}

func (r *RouteHandler) Create(route *Route) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		o := &objects{
			route:  route,
			r:      request,
			w:      writer,
			result: response.NewResult(),
		}

		r.runEvent(o)
	}

}

func (r *RouteHandler) runEvent(o *objects) {
	var keys []int
	for k := range r.events {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		for k := range r.events[key] {
			r.events[key][k](o)
		}
	}
}

func (r *RouteHandler) addEvents(e event, priority int) {
	if _, ok := r.events[priority]; !ok {
		r.events[priority] = []event{e}
	} else {
		r.events[priority] = append(r.events[priority], e)
	}
}

func (r *RouteHandler) methodAssert() {
	r.addEvents(func(o *objects) {

	}, EARLY_EVENT)
}

func (r *RouteHandler) firewalls() {
	services := services.GetServices()
	for _, f := range services.HttpServices.Security.Firewalls {
		r.addEvents(func(o *objects) {
			if f.IfPattern(o.route.Path) {
				if err := f.Authenticate(f.GetCredentialsFromRequest(o.r)); err != nil {
					o.result.StatusCode = http.StatusForbidden
					o.result.Msg = err.Error()
					o.result.Status = response.StatusError
				}
			}
		}, EARLY_EVENT+1)
	}
}

func (r *RouteHandler) atController() {
	r.addEvents(func(o *objects) {
		if o.result.StatusCode == http.StatusOK {
			//reflect to controller, call action
			v := reflect.ValueOf(o.route.Controller).MethodByName(o.route.Action).Call([]reflect.Value{})
			o.result = v[0].Interface().(*response.Result)
		}
	}, 0)
}

func (r *RouteHandler) beforeMiddlebrows() {
	services := services.GetServices()

	for _, m := range services.HttpServices.Middlewares {
		r.addEvents(func(o *objects) {
			m.Before(o.r)
		}, int(m.GetBeforePriority())-MIDDLEWARE_EVENT_P)
	}
}

func (r *RouteHandler) afterMiddlebrows() {
	services := services.GetServices()

	for _, m := range services.HttpServices.Middlewares {
		r.addEvents(func(o *objects) {
			m.After(o.result, o.w, o.r)
		}, int(m.GetAfterPriority())+MIDDLEWARE_EVENT_P)
	}
}
