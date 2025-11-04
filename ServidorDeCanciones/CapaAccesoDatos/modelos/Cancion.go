package modelos

type Cancion struct {
	Id            int32
	Titulo        string
	Artista_Banda string
	Lanzamiento   int32
	Duracion      string
	Ruta          string
	Idioma        string
	Genero        Genero
}

func NewCancion(titulo string, artista_banda string, lanzamiento int32, duracion string, idioma string, genero Genero) Cancion {
	objCancion := Cancion{Titulo: titulo, Artista_Banda: artista_banda, Lanzamiento: lanzamiento, Duracion: duracion, Idioma: idioma, Genero: genero}
	return objCancion
}

func NewCancionC(titulo string, artista_banda string, lanzamiento int32, duracion string, ruta string, idioma string, genero Genero) Cancion {
	objCancion := Cancion{Titulo: titulo, Artista_Banda: artista_banda, Lanzamiento: lanzamiento, Duracion: duracion, Ruta: ruta, Idioma: idioma, Genero: genero}
	return objCancion
}
