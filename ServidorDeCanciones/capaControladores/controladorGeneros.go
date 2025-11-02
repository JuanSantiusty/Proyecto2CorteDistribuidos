package capaControladores

import (
	"encoding/json"
	"net/http"
	se "servidor/grpc-servidor/capaFachadaServices"
)

type ControladorGenero struct {
	servicioGenero *se.ServiciosGenero
}

func NuevoControladorGenero() *ControladorGenero {
	servicio := se.NuevoServicioGeneros()
	return &ControladorGenero{
		servicioGenero: servicio,
	}
}

// ListarGeneros maneja las peticiones GET para listar todos los géneros
func (c *ControladorGenero) ListarGeneros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Llamar al servicio para obtener la lista de géneros
	generos := c.servicioGenero.ListaGeneros()

	// Crear respuesta
	respuesta := map[string]interface{}{
		"data":    generos,
		"codigo":  200,
		"mensaje": "Géneros obtenidos exitosamente",
	}

	// Enviar respuesta
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}

// ManejarGeneros es un router que maneja todas las rutas de géneros
func (c *ControladorGenero) ManejarGeneros(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case path == "/generos" && r.Method == http.MethodGet:
		c.ListarGeneros(w, r)
	default:
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
	}
}
