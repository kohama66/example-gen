package db

import (
	"api/infrastructure/db/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *Conn

type Conn struct {
	*gorm.DB
}

func init() {
	dsn := "example-gen-db:example-gen-db@tcp(db:3306)/example-gen-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	conn = &Conn{db}
}

func Init() {
	query.SetDefault(conn.DB)
}

func New() *Conn {
	return conn
}
