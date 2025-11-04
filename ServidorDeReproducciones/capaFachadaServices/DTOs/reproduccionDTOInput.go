package dtos

type ReproduccionDTOInput struct {
	CancionId int    `json:"cancionId"`
	Titulo    string `json:"titulo"`
	UsuarioId string `json:"usuarioId"`
}
