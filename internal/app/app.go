package app

import (
	"golang-core/internal/controller"
	"golang-core/internal/delivery/http"
	"golang-core/internal/repository"
	"golang-core/internal/server"
	"golang-core/internal/service"
	"golang-core/pkg/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func Run(configPath string) {

	InitEnv()
	
	
	// cfg, err := config.Init(configPath)

	// if err != nil {
	// 	log.Fatalf("ERROR: config file cannot load")
	// }

	db := database.Init()

	
	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repos:repos,
	})

	controllers := controller.NewControllers(controller.Deps{
		Services:services,
	})
	handlers := http.NewHandler(controllers)   
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	server := server.NewServer(handlers.Init(host, port))
	if err := server.Run(); err != nil {
		log.Fatalf("ERROR: error occurred while running http server: %s\n", err.Error())
	}
	// go func(){
	// 	if err := server.Run(); err != nil {
	// 		log.Fatalf("ERROR: error occurred while running http server: %s\n", err.Error())
	// 	}
	// }()
	log.Fatalf("SUCCESS: Project run with ->  %s%s", host,port)

}

func InitEnv() {
	err := godotenv.Load("./.env")
	if err != nil{
		log.Fatalf("ERROR: .env file cannot load")
	}
}