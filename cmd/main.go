package main

import (
	"database/sql"
	"github.com/Kisesaiyajin/testApi/cmd/api"
	"github.com/Kisesaiyajin/testApi/config"
	postgres "github.com/Kisesaiyajin/testApi/db"
	"log"
)

func main() {
	db, error := postgres.NewPostgreSQLDB(config.Config{
		Host:     config.Envs.Host,
		Port:     config.Envs.Port,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	})
	if error != nil {
		log.Fatalf("error in connecting to db: %s", error.Error())
	}
	initStorage(db)
	server := api.NewApiServer(":8080", db)
	if error := server.Run(); error != nil {
		log.Fatal(error)
	}

}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to storage")
}
