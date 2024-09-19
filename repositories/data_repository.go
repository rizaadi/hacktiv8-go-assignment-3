package repositories

type DataRepository interface {
	UpdateData(wind int, water int) error
}
