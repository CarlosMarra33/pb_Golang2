package services

import (
	"application/models"
	"application/repositorio"
	"fmt"
	"time"
)

func MarcarPresenca(idAula uint, idaluno uint) string  {

	var p models.Presenca
	p.AlunoId = idaluno
	p.AulaId = idAula
	p.Tipo = "presente"
	p.DataCreate = time.Now()

	fmt.Println("teste service")

	if !repositorio.ChecaPresenca(&p.AulaId, &p.AlunoId) {
		return "Presensa de hoje já foi marcada"
	}
	repositorio.MarcarPresenca(&p)
	return ""
}

func MarcarFalta(idAula uint, idaluno uint) error {

	p := models.Presenca{
		AlunoId:    idaluno,
		AulaId:     idAula,
		Tipo:       "falta",
		DataCreate: time.Now(),
	}

	if !repositorio.ChecaPresenca(&p.AulaId, &p.AlunoId) {
		return nil
	}
	repositorio.MarcarPresenca(&p)
	return nil
}

func MarcarAbono(idAula uint, idaluno uint) error {

	p := models.Presenca{
		AlunoId:    idaluno,
		AulaId:     idAula,
		Tipo:       "abono",
		DataCreate: time.Now(),
	}

	if !repositorio.ChecaPresenca(&p.AulaId, &p.AlunoId) {
		return nil
	}
	repositorio.MarcarPresenca(&p)

	return nil
}

func UpdatePresenca(idPresenca uint, tipo string) error {
	err := repositorio.AtualizarPresenca(tipo, idPresenca)
	
	if err != nil {
		return err
	}

	return nil
}

func GetPresencaAula(idAula uint, idAluno uint) []models.Presenca {
	presenca := repositorio.GetPresencaAula(&idAula, &idAluno)

	return presenca
}
