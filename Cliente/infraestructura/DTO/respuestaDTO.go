package DTO

type RespuestaDTO[T any] struct {
	Data    T
	Codigo  int32
	Mensaje string
}

func NewRespuestaDTO[T any](data T, codigo int32, mensaje string) RespuestaDTO[T] {
	return RespuestaDTO[T]{
		Data:    data,
		Codigo:  codigo,
		Mensaje: mensaje,
	}
}

func NuevaCancionRespuestaDTO(
	id int32,
	titulo string,
	artistaBanda string,
	lanzamiento int32,
	duracion string,
	ruta string,
	idioma string,
	genero string,
) CancionRespuestaDTO {
	return CancionRespuestaDTO{
		Id:            id,
		Titulo:        titulo,
		Artista_Banda: artistaBanda,
		Lanzamiento:   lanzamiento,
		Duracion:      duracion,
		Ruta:          ruta,
		Idioma:        idioma,
		Genero:        genero,
	}
}
