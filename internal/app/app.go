package app

import (
	"fmt"
	"golang-core/pkg/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func Run(configPath string) {

	InitEnv()
	
}

func InitEnv() {
	err := godotenv.Load("/.env")
	if err != nil{
		log.Fatalf("ERROR: .env file cannot load")
	}
	database.Init()
	port := os.Getenv("POST")
	_=handler.Run(port)
	fmt.Println("Project run with :",port)
}