package services

import (
	"application/models"
	"application/repositorio"
)

func CreateNewAula(profID uint, materia string, alunos []uint) (string, error) {
	var newAula models.Aula

	newAula.AlunosID = alunos
	newAula.Materia = materia
	newAula.ProfessorID = profID

	resp, err := repositorio.Create(newAula)

	return resp, err
}
