package capaControladores

import (
	"context"
	"fmt"
	"io"
	"strings"

	inf "cliente/infraestructura"
	dto "cliente/infraestructura/DTO"
	util "cliente/utilidades"
	pbStreaming "servidor-streaming/serviciosStreaming"
)

// ControladorSpotify maneja toda la lógica del cliente Spotify
type ControladorSpotify struct {
	servicioCancion   inf.FachadaCanciones
	serviciosGenero   inf.FachadaGeneros
	servicioStreaming inf.ServicioStreaming
	User              string
}

// NewControladorSpotify crea nueva instancia con ambos servicios
func NewControladorSpotify(servicioCancion inf.FachadaCanciones, serviciosGenero inf.FachadaGeneros, servicioStreaming inf.ServicioStreaming, user string) *ControladorSpotify {
	return &ControladorSpotify{
		servicioCancion:   servicioCancion,
		serviciosGenero:   serviciosGenero,
		servicioStreaming: servicioStreaming,
		User:              user,
	}
}

// Estructuras de datos
type Genero struct {
	ID     int
	Nombre string
}

type CancionDetalle struct {
	ID           int32
	Titulo       string
	Artista      string
	Album        string
	Anio         int32
	Duracion     string
	Genero       string
	ArchivoAudio string // para identificar el archivo
}

type RespuestaControlador struct {
	Exito   bool
	Mensaje string
	Datos   interface{}
}

// ===== FUNCIONES DE MENÚ PRINCIPAL =====

func (ctrl *ControladorSpotify) ValidarOpcionMenuPrincipal(opcion string) (int, error) {
	opcion = strings.TrimSpace(opcion)
	switch opcion {
	case "1":
		return 1, nil
	case "2":
		return 2, nil
	default:
		return 0, fmt.Errorf("opción no válida. Debe ser 1 o 2")
	}
}

func (ctrl *ControladorSpotify) ProcesarOpcionMenuPrincipal(opcion int) string {
	switch opcion {
	case 1:
		return "VER_GENEROS"
	case 2:
		return "SALIR"
	default:
		return "INVALIDO"
	}
}

// ===== FUNCIONES DE GÉNEROS =====

func (ctrl *ControladorSpotify) ObtenerGeneros() []dto.Genero {
	res, er := ctrl.serviciosGenero.ListarGeneros()
	if er != nil {
		fmt.Printf("Error listando geneos")
		return []dto.Genero{}
	}
	return res
}

func (ctrl *ControladorSpotify) ValidarOpcionGenero(opcion string, maxOpciones int) (int, error) {
	opcion = strings.TrimSpace(opcion)

	if opcion == fmt.Sprintf("%d", maxOpciones+1) {
		return maxOpciones + 1, nil
	}

	switch opcion {
	case "1", "2", "3":
		if opcion == "1" {
			return 1, nil
		}
		if opcion == "2" {
			return 2, nil
		}
		if opcion == "3" {
			return 3, nil
		}
		return 0, fmt.Errorf("opción no válida")
	default:
		return 0, fmt.Errorf("opción no válida. Debe ser 1, 2, 3 o 4")
	}
}

// ===== FUNCIONES DE CANCIONES POR GÉNERO =====

// Mapeo de nombres de canciones a archivos de audio
var mapaArchivosAudio = map[string]string{
	"franceslimon":        "frances_limon.mp3",
	"paranoolvidar":       "para_no_olvidar.mp3",
	"lluevesobrelaciudad": "llueve_sobre_la_ciudad.mp3",
	"somethingaboutyou":   "something_about_you.mp3",
	"bloodwasonmyskin":    "blood_was_on_my_skin.mp3",
	"badhabit":            "bad_habit.mp3",
	"sinolatengo":         "si_no_la_tengo.mp3",
}

// Mapeo de títulos de canciones a álbumes
var mapaAlbumes = map[string]string{
	"franceslimon":        "Amores lejanos",
	"paranoolvidar":       "Palabras más, palabras menos",
	"lluevesobrelaciudad": "Vida de perros",
	"somethingaboutyou":   "Mulholland drive",
	"bloodwasonmyskin":    "Blood was on my skin",
	"badhabit":            "Gemini rights",
	"sinolatengo":         "Con calor y sentimiento",
}

// ObtenerCancionesPorGenero retorna canciones según el género
func (ctrl *ControladorSpotify) ObtenerCancionesPorGenero(genero string) []dto.CancionRespuestaDTO {
	res, er := ctrl.servicioCancion.BuscarCancionesPorGenero(genero)
	if er != nil {
		fmt.Printf("Error buscando canciones por genero")
		return []dto.CancionRespuestaDTO{}
	}
	return res
}

func (ctrl *ControladorSpotify) ValidarOpcionCancion(opcion string, maxOpciones int) (int, error) {
	opcion = strings.TrimSpace(opcion)

	if opcion == fmt.Sprintf("%d", maxOpciones+1) {
		return maxOpciones + 1, nil
	}

	for i := 1; i <= maxOpciones; i++ {
		if opcion == fmt.Sprintf("%d", i) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("opción no válida")
}

// ===== FUNCIONES DE DETALLES DE CANCIÓN =====
/*
func (ctrl *ControladorSpotify) ObtenerDetallesCancion(cancionID int32) *CancionDetalle {
	todosGeneros := []int{1, 2, 3}

	for _, generoID := range todosGeneros {
		canciones := ctrl.ObtenerCancionesPorGenero(generoID)
		for _, cancion := range canciones {
			if cancion.ID == cancionID {
				return &cancion
			}
		}
	}

	return nil
}
*/

// BuscarCancionEnServidor busca CUALQUIER canción en el servidor (SIN datos quemados)
func (ctrl *ControladorSpotify) BuscarCancionEnServidor(cancion string) *dto.RespuestaDTO[dto.CancionRespuestaDTO] {
	res, _ := ctrl.servicioCancion.BuscarCancionPorTitulo(cancion)
	return res
}

// normalizarTitulo elimina espacios y convierte a minúsculas
func (ctrl *ControladorSpotify) normalizarTitulo(titulo string) string {
	// Eliminar espacios y convertir a minúsculas
	normalizado := strings.ToLower(titulo)
	normalizado = strings.ReplaceAll(normalizado, " ", "")
	normalizado = strings.ReplaceAll(normalizado, "á", "a")
	normalizado = strings.ReplaceAll(normalizado, "é", "e")
	normalizado = strings.ReplaceAll(normalizado, "í", "i")
	normalizado = strings.ReplaceAll(normalizado, "ó", "o")
	normalizado = strings.ReplaceAll(normalizado, "ú", "u")
	return normalizado
}

// obtenerArchivoAudio obtiene el nombre del archivo MP3
func (ctrl *ControladorSpotify) obtenerArchivoAudio(tituloNormalizado string) string {
	if archivo, existe := mapaArchivosAudio[tituloNormalizado]; existe {
		return archivo
	}
	return tituloNormalizado + ".mp3" // Fallback
}

// obtenerAlbum obtiene el nombre del álbum
func (ctrl *ControladorSpotify) obtenerAlbum(tituloNormalizado string) string {
	if album, existe := mapaAlbumes[tituloNormalizado]; existe {
		return album
	}
	return "Álbum Desconocido" // Fallback
}

func (ctrl *ControladorSpotify) ValidarOpcionDetalle(opcion string) (int, error) {
	opcion = strings.TrimSpace(opcion)
	switch opcion {
	case "1":
		return 1, nil
	case "2":
		return 2, nil
	default:
		return 0, fmt.Errorf("opción no válida. Debe ser 1 o 2")
	}
}

// ===== FUNCIONES DE STREAMING =====

// ReproducirCancionConAudio reproduce la canción con audio real usando pipes
func (ctrl *ControladorSpotify) ReproducirCancionConAudio(cancion string) error {
	fmt.Printf("\nLlamando método remoto: EnviarCancionMedianteStream\n")

	// Crear contexto
	ctx := context.Background()

	// Crear petición gRPC
	peticion := &pbStreaming.PeticionDTO{
		Titulo:  cancion,
		Formato: ctrl.User,
	}

	// Llamar al servidor de streaming
	stream, err := ctrl.servicioStreaming.EnviarCancionMedianteStream(ctx, peticion)
	if err != nil {
		return fmt.Errorf("error al iniciar streaming: %v", err)
	}

	// Crear pipe para comunicación entre recepción y reproducción
	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})

	// Goroutine para decodificar y reproducir audio
	go util.DecodificarReproducir(reader, canalSincronizacion)

	// Recibir fragmentos y escribir en el pipe
	util.RecibirCancion(stream, writer, canalSincronizacion)

	return nil
}

func (ctrl *ControladorSpotify) ValidarOpcionStreaming(opcion string) (int, error) {
	opcion = strings.TrimSpace(opcion)
	if opcion == "1" {
		return 1, nil
	}
	return 0, fmt.Errorf("opción no válida. Debe ser 1")
}
