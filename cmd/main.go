package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/iamengg/ecom/api"
	"github.com/iamengg/ecom/db"
)

func main() {
	dbInstance := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "abc",
		Addr:                 "http://localhost",
		Port:                 ":8080",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            time,
	})
	server := api.NewAPIServer(":8080", dbInstance)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
