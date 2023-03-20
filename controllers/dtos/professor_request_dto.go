package dtos

type ProfessoAulaDto struct {
	ProfessorID uint   `json:"professor_id"`
	Materia     string `json:"materia"`
	Alunos      []uint `json:"alunos"`
}
