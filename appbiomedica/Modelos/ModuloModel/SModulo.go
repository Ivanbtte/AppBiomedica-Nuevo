package ModuloModel

type Modulo struct {
	Id          string `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status,omitempty"`
}
type Action struct {
	Id          string  `gorm:"PRIMARY_KEY" json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Value       bool    `json:"value,omitempty"`
	ModuloId    string  `json:"moduloId,omitempty"`
	Modulo      *Modulo `json:"modulo,omitempty"`
	Status      bool    `json:"status,omitempty"`
}
