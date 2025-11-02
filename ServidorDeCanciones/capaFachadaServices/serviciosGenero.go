package capaFachadaServices

import (
	rp "servidor/grpc-servidor/CapaAccesoDatos"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
)

type ServiciosGenero struct {
	repo *rp.GeneroRepo
}

func NuevoServicioGeneros() *ServiciosGenero {
	repo := rp.GetRepositorioGeneros()
	return &ServiciosGenero{
		repo: repo,
	}
}

// BuscarOCrearGenero busca un género por nombre, si existe lo retorna, si no existe lo crea
func (s *ServiciosGenero) BuscarOCrearGenero(nombreGenero string) mo.Genero {
	return s.repo.BuscarOCrearGeneroPorNombre(nombreGenero)
}

// ListaGeneros retorna la lista de todos los géneros
func (s *ServiciosGenero) ListaGeneros() []mo.Genero {
	return s.repo.ListaGeneros()
}

// AgregarGenero guarda un nuevo género
func (s *ServiciosGenero) AgregarGenero(genero mo.Genero) mo.Genero {
	return s.repo.AgregarGenero(genero)
}
