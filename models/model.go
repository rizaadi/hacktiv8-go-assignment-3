package models

type Model struct {
	Id    int `gorm:"primaryKey" json:"id"`
	Water int `json:"water"`
	Wind  int `json:"widn"`
}
