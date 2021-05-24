package main

import "golang-core/internal/app"

const configPath = "configs/main_config"

func main() {
	app.Run(configPath)
}