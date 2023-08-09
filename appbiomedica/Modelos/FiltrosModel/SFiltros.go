package FiltrosModel

type Tema struct {
	Id   int `gorm:"PRIMARY_KEY"`
	Tema string
}
type GrupoTema struct {
	Grupo  string `gorm:"PRIMARY_KEY"`
	TemaId int
	Tema   *Tema
}
