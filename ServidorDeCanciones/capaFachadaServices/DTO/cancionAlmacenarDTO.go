package DTO

type CancionAlmacenarDTO struct {
	Titulo        string `json:"titulo"`
	Artista_Banda string `json:"artista_banda"`
	Lanzamiento   int32  `json:"lanzamiento"`
	Duracion      string `json:"duracion"`
	Idioma        string `json:"idioma"`
	Genero        string `json:"genero"`
}
