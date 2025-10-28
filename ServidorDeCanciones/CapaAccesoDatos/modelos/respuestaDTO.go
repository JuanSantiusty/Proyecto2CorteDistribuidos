package modelos

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
