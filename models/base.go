package models

import (
	"database/sql"
	"review-server/core/database"
)

func GetDB() *sql.DB {
	return database.GetDB()
}
