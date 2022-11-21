package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


var db *sqlx.DB
var err error


func Init() *sqlx.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable binary_parameters=yes", 
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"),
	)
	
	db,err = sqlx.Open("postgres", dbInfo)
	err = db.Ping()
	if err != nil {
		log.Println("ERROR: cannot connect to database")
		panic(err)
	}
	
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(5* time.Minute)
	log.Println("The database is connected")
	return db
}

func GetDB() *sqlx.DB{
	return db
}