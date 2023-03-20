package dtos

type PresencaDto struct {
	IdPresenca uint   `json:"idPresenca"`
	IdAluno    uint   `json:"alunoId"`
	IdAula     uint   `json:"aulaId"`
	Tipo       string `json:"tipo"`
}
