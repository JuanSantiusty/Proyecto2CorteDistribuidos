package DTO

type CancionRespuestaDTO struct {
	Titulo        string `json:"titulo"`
	Artista_Banda string `json:"artista_banda"`
	Lanzamiento   int32  `json:"lanzamiento"`
	Duracion      string `json:"duracion"`
	Ruta          string `json:"ruta"`
	Genero        string `json:"genero"`
}
