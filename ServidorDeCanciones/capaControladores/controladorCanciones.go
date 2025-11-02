package capaControladores

import (
	"encoding/json"
	"net/http"
	se "servidor/grpc-servidor/capaFachadaServices"
	dto "servidor/grpc-servidor/capaFachadaServices/DTO"
	"strconv"
	"strings"
)

type ControladorCancion struct {
	servicioCancion *se.ServiciosCancion
}

func NuevoControladorCancion() *ControladorCancion {
	servicio := se.NuevoServicioCanciones()
	return &ControladorCancion{
		servicioCancion: servicio,
	}
}

// AlmacenarCancion maneja las peticiones POST para almacenar una nueva canción
func (c *ControladorCancion) AlmacenarCancion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parsear el formulario multipart para obtener archivo y datos
	err := r.ParseMultipartForm(32 << 20) // 32 MB máximo
	if err != nil {
		http.Error(w, "Error al procesar el formulario: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener el archivo de audio
	file, handler, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "Error al obtener el archivo de audio: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Leer los datos del archivo
	fileBytes := make([]byte, handler.Size)
	_, err = file.Read(fileBytes)
	if err != nil {
		http.Error(w, "Error al leer el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener los datos del formulario
	titulo := r.FormValue("titulo")
	artistaBanda := r.FormValue("artista_banda")
	lanzamientoStr := r.FormValue("lanzamiento")
	duracion := r.FormValue("duracion")
	genero := r.FormValue("genero")

	// Validar campos obligatorios
	if titulo == "" || artistaBanda == "" || lanzamientoStr == "" || duracion == "" || genero == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	// Convertir lanzamiento a int32
	lanzamiento, err := strconv.Atoi(lanzamientoStr)
	if err != nil {
		http.Error(w, "El año de lanzamiento debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Crear DTO
	cancionDTO := dto.CancionAlmacenarDTO{
		Titulo:        titulo,
		Artista_Banda: artistaBanda,
		Lanzamiento:   int32(lanzamiento),
		Duracion:      duracion,
		Genero:        genero,
	}

	// Llamar al servicio
	err = c.servicioCancion.AlmacenarCancion(cancionDTO, fileBytes)
	if err != nil {
		http.Error(w, "Error al almacenar la canción: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	respuesta := dto.NewRespuestaDTO(cancionDTO, 200, "Canción almacenada exitosamente")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

// BuscarCancion maneja las peticiones GET para buscar una canción por título
func (c *ControladorCancion) BuscarCancion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener el título de los query parameters
	titulo := r.URL.Query().Get("titulo")
	if titulo == "" {
		http.Error(w, "El parámetro 'titulo' es requerido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio
	respuesta := c.servicioCancion.BuscarCancion(titulo)

	// Enviar respuesta
	w.WriteHeader(int(respuesta.Codigo))
	json.NewEncoder(w).Encode(respuesta)
}

// BuscarPorGenero maneja las peticiones GET para buscar canciones por género
func (c *ControladorCancion) BuscarPorGenero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener el género de los query parameters
	genero := r.URL.Query().Get("genero")
	if genero == "" {
		http.Error(w, "El parámetro 'genero' es requerido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio
	respuesta := c.servicioCancion.BuscarPorGenero(genero)

	// Enviar respuesta
	w.WriteHeader(int(respuesta.Codigo))
	json.NewEncoder(w).Encode(respuesta)
}

// ListarCanciones maneja las peticiones GET para listar todas las canciones
func (c *ControladorCancion) ListarCanciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Llamar al servicio
	respuesta := c.servicioCancion.ListarCanciones()

	// Enviar respuesta
	w.WriteHeader(int(respuesta.Codigo))
	json.NewEncoder(w).Encode(respuesta)
}

// ManejarCanciones es un router que maneja todas las rutas de canciones
func (c *ControladorCancion) ManejarCanciones(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/canciones")

	switch {
	case path == "" && r.Method == http.MethodPost:
		c.AlmacenarCancion(w, r)
	case path == "" && r.Method == http.MethodGet:
		c.ListarCanciones(w, r)
	case strings.Contains(path, "/buscar") && r.Method == http.MethodGet:
		// Extraer parámetros de query para búsqueda
		if r.URL.Query().Get("titulo") != "" {
			c.BuscarCancion(w, r)
		} else if r.URL.Query().Get("genero") != "" {
			c.BuscarPorGenero(w, r)
		} else {
			http.Error(w, "Parámetro de búsqueda no válido", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Ruta no encontrada", http.StatusNotFound)
	}
}
