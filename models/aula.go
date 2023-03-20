package models

type Aula struct {
	AulaId      uint   `gorm:"primaryKey;autoIncrement"`
	Materia     string `json:"materia"`
	ProfessorID uint   `json:"professorID"`
	AlunosID    []uint `json:"alunosID"`
}
