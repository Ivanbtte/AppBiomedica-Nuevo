package FechaLimiteModel

type FechaLimite struct {
	Id          string `gorm:"PRIMARY_KEY"`
	FechaLimite string
	FechaInicio string
	Estado      bool
}
