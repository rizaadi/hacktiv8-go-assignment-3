package services

import (
	"hacktiv8-go-assignment-3/helpers"
	"hacktiv8-go-assignment-3/repositories"
	"log"
	"math/rand"
	"time"
)

type DataUpdater struct {
	Repo repositories.DataRepository
}

func NewDataUpdater(repo repositories.DataRepository) *DataUpdater {
	return &DataUpdater{Repo: repo}
}
func (u *DataUpdater) UpdateData() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		err := u.Repo.UpdateData(water, wind)
		if err != nil {
			log.Println("Error updating data:", err)
			continue
		}
		log.Printf("Water: %d meter, Status: %s", water, helpers.WaterStatus(water))
		log.Printf("Wind: %d meter, Status: %s", wind, helpers.WindStatus(wind))

		time.Sleep(15 * time.Second)
	}
}
