package services

import "database/sql"

type DbServices struct {
	MysqlDbCoon *sql.DB
}

func NewDbServices() *DbServices {
	return &DbServices{}
}
