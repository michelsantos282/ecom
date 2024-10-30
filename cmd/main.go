package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/michelsantos282/ecom/cmd/api"
	"github.com/michelsantos282/ecom/config"
	"github.com/michelsantos282/ecom/db"
)

func main() {
	db := connectDatabase()
	server := api.NewApiServer(":8000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func connectDatabase() *sql.DB {
	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("error while initiating database: %v", err)
	}

	initStorage(db)
	return db
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}
