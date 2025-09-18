package main

import (
	"fmt"
	"log"
	"read_files/router"
	"read_files/util"
	"read_files/util/constants"
)

func main() {
	const port = ":3000"
	app := router.InitializeRoutes()
	if err := app.Listen(port); err != nil {
		util.CustomLogger(constants.Error, fmt.Sprintf("Listen: %v", err))
		log.Panicf("Falha ao iniciar o servidor : %v", err)
	}

}
