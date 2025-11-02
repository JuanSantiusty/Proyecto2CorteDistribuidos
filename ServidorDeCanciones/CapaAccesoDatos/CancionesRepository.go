package CapaAccesoDatos

import (
	"fmt"
	"os"
	"path/filepath"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
	"sync"
)

type CancionesRepo struct {
	canciones []mo.Cancion
	mu        sync.Mutex
}

var (
	instance        *CancionesRepo
	one             sync.Once
	indiceIdCancion int32
)

// Aplicamos patron singleton para el repositorio de canciones
func GetRepositorioCanciones() *CancionesRepo {
	one.Do(func() {
		instance = &CancionesRepo{}
		indiceIdCancion = 0
		indiceIdGenero = 0
	})
	return instance
}

func (r *CancionesRepo) AgregarCancion(cancion mo.Cancion, data []byte) error {
	//Asignar Id de manera automatica
	obtenerIdCancion()
	cancion.Id = indiceIdCancion
	r.mu.Lock()
	defer r.mu.Unlock()
	//Construir carpeta
	os.MkdirAll("../Canciones", os.ModePerm)

	//Costruir archivo
	fileName := fmt.Sprintf("%s_%s_%s.mp3", cancion.Titulo, cancion.Genero.Nombre, cancion.Artista_Banda)
	filePath := filepath.Join("../Canciones", fileName)

	//Guardar archivo fisico
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("Error al guardar el archivo: %v", err)
	}

	cancion.Ruta = "../Canciones/" + fileName
	r.canciones = append(r.canciones, cancion)

	return nil
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

func obtenerIdCancion() int32 {
	indiceIdCancion += 1
	return indiceIdCancion
}
