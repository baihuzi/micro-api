package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var dbConn *sql.DB

type Mysql struct {
	host     string
	port     string
	user     string
	password string
	database string
	charset  string
	url      string
}

func (m *Mysql) Init(host, port, user, password, database, charset string) {
	m.host = host
	m.port = port
	m.user = user
	m.password = password
	m.database = database
	m.charset = charset
	m.url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, port, database, charset)
}

func (m *Mysql) Conn() {
	db, err := sql.Open("mysql", m.url)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(100 * time.Second)
	dbConn = db
}

func GetDB() (*sql.DB) {
	return dbConn
}
