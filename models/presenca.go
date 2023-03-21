package models

type Presenca struct {
	PresencaId uint      `gorm:"primaryKey;autoIncrement"`
	AlunoId    uint      `json:"alunoId"`
	AulaId     uint      `json:"aulaId"`
	Tipo       string    `json:"tipoPresenca"`
	DataCreate int `json:"created"`
	DataUpdate int `json:"updated"`
}
