package models

import "time"

type Presenca struct {
	PresencaId uint   `gorm:"primaryKey;autoIncrement"`
	AlunoId    uint   `json:"alunoId"`
	AulaId     uint   `json:"aulaId"`
	Tipo       string `json:"tipoPresenca"`
	DataCreate time.Time
	DataUpdate time.Time
}
