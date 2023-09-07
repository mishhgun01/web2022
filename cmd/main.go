package main

import (
	"log"
	"web2022/internal/app"
)

func main() {
	err := app.StartServer()

	if err != nil {
		log.Println(err.Error())
		return
	}
}
