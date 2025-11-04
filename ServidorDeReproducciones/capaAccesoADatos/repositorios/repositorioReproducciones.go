package capaaccesoadatos

import (
	"fmt"
	entitys "reproducciones/capaAccesoADatos/entitys"
	"sync"
	"time"
)

type RepositorioReproducciones struct {
	mu             sync.Mutex
	reproducciones []entitys.ReproduccionEntity
	ultimoID       int
}

var (
	instancia *RepositorioReproducciones
	once      sync.Once
)

func GetRepositorio() *RepositorioReproducciones {
	once.Do(func() {
		instancia = &RepositorioReproducciones{}
	})
	return instancia
}
func (r *RepositorioReproducciones) AgregarReproduccion(cancionId int, titulo string, usuarioId string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Generar ID autoincremental
	r.ultimoID++
	idReproduccion := r.ultimoID

	reproduccion := entitys.ReproduccionEntity{
		ID:        idReproduccion,
		CancionId: cancionId,
		Titulo:    titulo,
		UsuarioId: usuarioId,
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
	}

	r.reproducciones = append(r.reproducciones, reproduccion)
	fmt.Printf("Reproducción almacenada: %+v\n", reproduccion)
	r.mostrarReproducciones()
}

func (r *RepositorioReproducciones) ListarPorCliente(usuarioId string) []entitys.ReproduccionEntity {
	r.mu.Lock()
	defer r.mu.Unlock()

	var resultado []entitys.ReproduccionEntity
	for _, rep := range r.reproducciones {
		if rep.UsuarioId == usuarioId {
			resultado = append(resultado, rep)
		}
	}
	return resultado
}

func (r *RepositorioReproducciones) mostrarReproducciones() {
	fmt.Println("==Reproducciones almacenadas==")
	for _, rep := range r.reproducciones {
		fmt.Printf("ID Repro: %d - Cliente: %s - CancionID: %d - Titulo: %s - Fecha: %s\n",
			rep.ID, rep.UsuarioId, rep.CancionId, rep.Titulo, rep.FechaHora)
	}
}
