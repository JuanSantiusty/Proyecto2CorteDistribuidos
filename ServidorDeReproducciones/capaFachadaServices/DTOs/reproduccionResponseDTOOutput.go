package dtos

// DTO que se devuelve al cliente REST después de listar una reproduccion
type ReproduccionResponseDTOOutput struct {
	Id        int    `json:"id"`
	IdUsuario string `json:"idUsuario"`
	IdCancion int    `json:"idCancion"`
	Titulo    string `json:"titulo"`
	FechaHora string `json:"fechaHora"`
}
