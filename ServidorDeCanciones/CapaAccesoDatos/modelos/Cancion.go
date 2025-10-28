package modelos

type Cancion struct {
	Id            int32
	Titulo        string
	Artista_Banda string
	Lanzamiento   int32
	Duracion      string
	Genero        Genero
}

func NewCancion(id int32, titulo string, artista_banda string, lanzamiento int32, duracion string, genero Genero) Cancion {
	objCancion := Cancion{Id: id, Titulo: titulo, Artista_Banda: artista_banda, Lanzamiento: lanzamiento, Duracion: duracion, Genero: genero}
	return objCancion
}
