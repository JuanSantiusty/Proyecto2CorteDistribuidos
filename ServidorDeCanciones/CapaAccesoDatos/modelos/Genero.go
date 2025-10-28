package modelos

type Genero struct {
	Id     int32
	Nombre string
}

func NewGenero(id int32, nombre string) Genero {
	objGenero := Genero{Id: id, Nombre: nombre}
	return objGenero
}
