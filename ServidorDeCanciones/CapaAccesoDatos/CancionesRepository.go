package CapaAccesoDatos

import (
	"fmt"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
)

type CancionesRepo struct {
	canciones []mo.Cancion
}

func NewCancionesRepo() *CancionesRepo {
	return &CancionesRepo{
		canciones: []mo.Cancion{},
	}
}

func (r *CancionesRepo) AgregarCancion(cancion mo.Cancion) {
	r.canciones = append(r.canciones, cancion)
}

func (r *CancionesRepo) ActualizarCancion(cancionActualizada mo.Cancion) error {
	for i, cancion := range r.canciones {
		if cancion.Id == cancionActualizada.Id {
			r.canciones[i] = cancionActualizada
			return nil
		}
	}
	return fmt.Errorf("canción con ID %d no encontrada", cancionActualizada.Id)
}

// EliminarCancion elimina una canción por ID
func (r *CancionesRepo) EliminarCancion(id int32) error {
	for i, cancion := range r.canciones {
		if cancion.Id == id {
			// Eliminar manteniendo el orden
			r.canciones = append(r.canciones[:i], r.canciones[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("canción con ID %d no encontrada", id)
}

func (r *CancionesRepo) BuscarPorGenero(genero string) []mo.Cancion {
	var cancionesGenero []mo.Cancion
	for _, cancion := range r.canciones {
		if cancion.Genero.Nombre == genero {
			cancionesGenero = append(cancionesGenero, cancion)
		}
	}
	return cancionesGenero
}

func (r *CancionesRepo) ListaCanciones() []mo.Cancion {
	return r.canciones
}

func (r *CancionesRepo) BuscarPorTitulo(nombre string) mo.Cancion {
	for _, cancion := range r.canciones {
		if cancion.Titulo == nombre {
			return cancion
		}
	}
	return mo.Cancion{}
}
