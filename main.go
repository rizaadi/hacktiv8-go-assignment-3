package main

import (
	"hacktiv8-go-assignment-3/database"
	"hacktiv8-go-assignment-3/repositories"
	"hacktiv8-go-assignment-3/services"
	"log"
)

func main() {
	database.StartDB()
	db := database.GetDB()
	repo := repositories.NewDataRepository(db)
	update := services.NewDataUpdater(repo)

	go update.UpdateData()
	log.Println("RUN")
	select {}
}
