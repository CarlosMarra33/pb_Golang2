package models

import "github.com/lib/pq"

type Aula struct {
	AulaId      uint          `gorm:"primaryKey;autoIncrement"`
	Materia     string        `json:"materia"`
	ProfessorID uint          `json:"professorID"`
	AlunosID    pq.Int64Array `json:"alunos" gorm:"type:int[]"`
}