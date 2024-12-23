package main

import (
	"log"

	"catchhotels/bootstrap"
)

// @title Catchhotels Task Documentation
// @version 1.0
// @description Example Dataverse Powerapps DnD app
// @host localhost
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if bootErr := bootstrap.Boot(); bootErr != nil {
		log.Fatalln(bootErr)
	}
}
