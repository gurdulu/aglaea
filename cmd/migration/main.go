package main

import (
	"github.com/gurdulu/aglaea/inits"
	"github.com/gurdulu/aglaea/models"
	"log"
)

func init() {
	inits.LoadEnvs()
	inits.ConnectDB()
}

func main() {
	err := inits.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
