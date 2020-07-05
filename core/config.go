package core

import (
	"github.com/Unknwon/goconfig"
	"review-server/core/database"
	"review-server/core/http"
	"sync"
)

type App struct {
	ReLoginSecret string
}

type Config struct {
	Mysql database.Mysql
	Http  http.Http
	App   App
}

var ins *Config
var once sync.Once

func loadConfig() {
	cfg, err := goconfig.LoadConfigFile("config/config.ini")
	if err != nil {
		panic("Load config.ini file fail." + err.Error())
	}

	appConf, err := cfg.GetSection("app")
	if err != nil {
		panic("Load app config file fail." + err.Error())
	}

	mysqlConf, err := cfg.GetSection("mysql")
	if err != nil {
		panic("Load mysql config file fail." + err.Error())
	}

	httpConf, err := cfg.GetSection("http")
	if err != nil {
		panic("Load http config file fail." + err.Error())
	}
	once.Do(func() {
		ins = &Config{
			Mysql: database.Mysql{},
			Http:  http.Http{},
			App:   App{},
		}
		ins.Mysql.Init(mysqlConf["host"], mysqlConf["port"], mysqlConf["user"], mysqlConf["password"], mysqlConf["database"], mysqlConf["charset"])
		ins.Http.Init(httpConf["host"], httpConf["port"])
		ins.App.ReLoginSecret = appConf["re_login_secret"]
	})
}

func GetConfig() *Config {
	return ins
}
