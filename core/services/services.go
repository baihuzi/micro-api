package services

import "sync"

type Services struct {
	DbServices   *DbServices
	HttpServices *HttpServices
	LogServices  *LogServices
}

var ins *Services
var once sync.Once

func GetServices() *Services {
	once.Do(func() {
		ins = &Services{}
		ins.initHttpServices(NewHttpServices())

	})
	return ins
}
func (s *Services) InitDbServices(dbServices *DbServices) {
	s.DbServices = dbServices
}
func (s *Services) initHttpServices(httpServices *HttpServices) {
	s.HttpServices = httpServices
}
func (s *Services) InitLogServices(logServices *LogServices) {
	s.LogServices = logServices
}
