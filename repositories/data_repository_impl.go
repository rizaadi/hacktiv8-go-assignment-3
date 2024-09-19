package repositories

import (
	"gorm.io/gorm"
	"hacktiv8-go-assignment-3/models"
)

type DataRepositoryImpl struct {
	DB *gorm.DB
}

func NewDataRepository(db *gorm.DB) DataRepository {
	return &DataRepositoryImpl{DB: db}
}

// UpdateData implements DataRepository.
func (r *DataRepositoryImpl) UpdateData(wind int, water int) error {
	data := models.Model{Water: water, Wind: wind}
	result := r.DB.Create(&data)
	return result.Error
}
