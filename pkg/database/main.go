package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)


var db *gorm.DB
var err error


func Init() *gorm.DB{
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable binary_parametrs=yes", 
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"),
	)

	db,err = gorm.Open("postgres", dbInfo)

	if err != nil{
		log.Println("ERROR: cannot connect to database")
		panic(err)
	}

	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(5* time.Minute)
	log.Println("The database connected")
	return db
}

func GetDB() *gorm.DB{	
		return db
}