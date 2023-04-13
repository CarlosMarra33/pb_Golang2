package services

import (
	"application/models"
	"application/repositorio"
	"fmt"
	"time"
)

type PresencaService struct {
	repo repositorio.PresencaRepository
}

func NewPresencaService(repo repositorio.PresencaRepository) *PresencaService{
	return &PresencaService{
		repo : repo,
	}
}

func (p *PresencaService)MarcarPresenca(idAula uint, idaluno uint) string  {

	var presenca models.Presenca
	presenca.AlunoId = idaluno
	presenca.AulaId = idAula
	presenca.Tipo = "presente"
	presenca.DataCreate = time.Now()

	fmt.Println("teste service")

	if !p.repo.ChecaPresenca(&presenca.AulaId, &presenca.AlunoId) {
		return "Presensa de hoje já foi marcada"
	}
	p.repo.MarcarPresenca(&presenca)
	return ""
}

func (p *PresencaService)MarcarFalta(idAula uint, idaluno uint) error {

	presenca := models.Presenca{
		AlunoId:    idaluno,
		AulaId:     idAula,
		Tipo:       "falta",
		DataCreate: time.Now(),
	}

	if !p.repo.ChecaPresenca(&presenca.AulaId, &presenca.AlunoId) {
		return nil
	}
	p.repo.MarcarPresenca(&presenca)
	return nil
}

func (p *PresencaService) MarcarAbono(idAula uint, idaluno uint) error {

	presenca := models.Presenca{
		AlunoId:    idaluno,
		AulaId:     idAula,
		Tipo:       "abono",
		DataCreate: time.Now(),
	}
	
	 
	if !p.repo.ChecaPresenca(&presenca.AulaId, &presenca.AlunoId) {
		return nil
	}
	p.repo.MarcarPresenca(&presenca)

	return nil
}

func (p *PresencaService) UpdatePresenca(idPresenca uint, tipo string) error {
	err := p.repo.AtualizarPresenca(tipo, idPresenca)
	
	if err != nil {
		return err
	}

	return nil
}

func (p *PresencaService) GetPresencaAula(idAula uint, idAluno uint) []models.Presenca {
	presenca := p.repo.GetPresencaAula(&idAula, &idAluno)

	return presenca
}
