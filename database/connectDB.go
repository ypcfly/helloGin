package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func GetDataBase() *sql.DB {
	url := "postgres://postgres:123456@localhost/gosql?sslmode=disable"
	log.Println(">>>> get database connection action start <<<<")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	// 返回数据库连接
	return db
}
