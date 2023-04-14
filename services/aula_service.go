package services

import (
	"application/models"
	"application/repositorio"
)

type AulaService struct {
	repo repositorio.AulaRepository
}

func NewAulaService(aulaRepo repositorio.AulaRepository) *AulaService {
	return &AulaService{repo: aulaRepo}
}

func (as *AulaService) CreateNewAula(profID uint, materia string, alunos []uint) (string, error) {
	var newAula models.Aula
	valoresInt := make([]int64, len(alunos))

	for i, v := range alunos {
		valoresInt[i] = int64(v)
	}

	newAula.AlunosID = valoresInt
	newAula.Materia = materia
	newAula.ProfessorID = profID

	resp, err := as.repo.Create(newAula)

	return resp, err
}
