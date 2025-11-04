package controlador

import (
	"encoding/json"
	"fmt"
	"net/http"
	dtos "reproducciones/capaFachadaServices/DTOs"
	capafachada "reproducciones/capaFachadaServices/fachada"
)

type ControladorTendencias struct {
	fachada *capafachada.FachadaReproducciones
}

func NuevoControladorTendencias() *ControladorTendencias {
	return &ControladorTendencias{
		fachada: capafachada.NuevaFachadaTendencias(),
	}
}

// Servicio REST POST que recibe una reproduccion en formato json
func (c *ControladorTendencias) RegistrarReproduccionHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ReproduccionDTOInput

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Error al leer el cuerpo de la peticion", http.StatusBadRequest)
		return
	}
	c.fachada.RegistrarReproduccion(dto.CancionId, dto.Titulo, dto.UsuarioId)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Reproduccion registrada correctamente")
}

// Servicio REZST GET que devuelve todas las reproducciones para un nickname en formato json
func (c *ControladorTendencias) ListarReproduccionesPorClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener idUsuario por QueryParam
	idUsuarioStr := r.URL.Query().Get("idUsuario")
	if idUsuarioStr == "" {
		http.Error(w, "Debe proporcionar un idUsuario", http.StatusBadRequest)
		return
	}

	repros := c.fachada.ObtenerReproduccionesPorCliente(idUsuarioStr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repros)
}
