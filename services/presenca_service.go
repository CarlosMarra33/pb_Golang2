package services

import (
	"application/models"
	"application/repositorio"
	"time"
)

func MarcarPresenca(idAula uint, idaluno uint) (error) {
	var p models.Presenca
	p.AlunoId = idaluno
	p.AulaId = idAula
	p.Tipo = "presente"
	p.DataCreate = time.Now()

	if !repositorio.ChecaPresenca(p.AulaId, p.AlunoId) {
		return nil
	}
	err := repositorio.MarcarPresenca(p)
	if err != nil {
		return err
	}
	return nil

}

func MarcarFalta(idAula uint, idaluno uint)(error){
	var p models.Presenca
	p.AlunoId = idaluno
	p.AulaId = idAula
	p.Tipo = "falta"
	p.DataCreate = time.Now()

	if !repositorio.ChecaPresenca(p.AulaId, p.AlunoId) {
		return nil
	}
	err := repositorio.MarcarPresenca(p)
	if err != nil {
		return err
	}
	return nil
}

func MarcarAbono(idAula uint, idaluno uint)(error){
	var p models.Presenca
	p.AlunoId = idaluno
	p.AulaId = idAula
	p.Tipo = "falta"
	p.DataCreate = time.Now()

	if !repositorio.ChecaPresenca(p.AulaId, p.AlunoId) {
		return nil
	}
	err := repositorio.MarcarPresenca(p)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePresenca(idPresenca uint, tipo string) error{
	
	err := repositorio.AtualizarPresenca(tipo, idPresenca)

	if err != nil {
		return err
	}

	return nil
}


func GetPresencaAula(idAula uint, idAluno uint) []models.Presenca{
	var presenca []models.Presenca

	presenca = repositorio.GetPresencaAula(idAula, idAluno)

	return presenca
}