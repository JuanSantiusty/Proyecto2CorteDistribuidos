package CapaAccesoDatos

import (
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
	"sync"
)

type GeneroRepo struct {
	generos []mo.Genero
	mu      sync.Mutex
}

var (
	instanceGenero *GeneroRepo
	onceGenero     sync.Once
	indiceIdGenero int32
)

// GetRepositorioGeneros aplica patrón singleton para el repositorio de géneros
func GetRepositorioGeneros() *GeneroRepo {
	onceGenero.Do(func() {
		instanceGenero = &GeneroRepo{}
		indiceIdGenero = 0
	})
	instanceGenero.CargarGeneros()
	return instanceGenero
}

// AgregarGenero guarda un género y asigna el ID de forma automática
func (r *GeneroRepo) AgregarGenero(genero mo.Genero) mo.Genero {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Asignar ID automático
	indiceIdGenero++
	genero.Id = indiceIdGenero

	r.generos = append(r.generos, genero)
	return genero
}

// ListaGeneros retorna la lista de todos los géneros
func (r *GeneroRepo) ListaGeneros() []mo.Genero {
	return r.generos
}

// BuscarOCrearGeneroPorNombre busca un género por nombre, si existe lo retorna, si no existe crea uno nuevo
func (r *GeneroRepo) BuscarOCrearGeneroPorNombre(nombreGenero string) mo.Genero {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Buscar género existente
	for _, genero := range r.generos {
		if genero.Nombre == nombreGenero {
			return genero
		}
	}

	// Crear nuevo género si no existe
	nuevoGenero := mo.NewGenero(indiceIdGenero+1, nombreGenero)
	indiceIdGenero++
	r.generos = append(r.generos, nuevoGenero)

	return nuevoGenero
}

// BuscarGeneroPorNombre busca un género por nombre y retorna si fue encontrado
func (r *GeneroRepo) BuscarGeneroPorNombre(nombreGenero string) (mo.Genero, bool) {
	for _, genero := range r.generos {
		if genero.Nombre == nombreGenero {
			return genero, true
		}
	}
	return mo.Genero{}, false
}

func (r *GeneroRepo) CargarGeneros() {
	r.AgregarGenero(mo.NewGenero(0, "Alternativo"))
	r.AgregarGenero(mo.NewGenero(0, "Rock"))
	r.AgregarGenero(mo.NewGenero(0, "Salsa"))
}
