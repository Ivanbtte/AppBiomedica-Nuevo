package EstatusModel

type Estatus struct {
	Id     int `gorm:"PRIMARY_KEY"`
	Status string
}
