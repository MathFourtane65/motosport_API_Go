package config

import (
	"database/sql"
	//"fmt"
	"log"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	//connexion à la BDD
	connStr := "postgresql://postgres:root@localhost/api_go_mathieu_fourtane?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
