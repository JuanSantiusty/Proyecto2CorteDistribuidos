package capaControladores

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	inf "cliente/infraestructura"
	util "cliente/utilidades"
	pbStreaming "servidor-streaming/serviciosStreaming"
	pb "servidor/grpc-servidor/serviciosCancion"
)

// ControladorSpotify maneja toda la lógica del cliente Spotify
type ControladorSpotify struct {
	servicioCancion   inf.ServicioCanciones
	servicioStreaming inf.ServicioStreaming
}

// NewControladorSpotify crea nueva instancia con ambos servicios
func NewControladorSpotify(servicioCancion inf.ServicioCanciones, servicioStreaming inf.ServicioStreaming) *ControladorSpotify {
	return &ControladorSpotify{
		servicioCancion:   servicioCancion,
		servicioStreaming: servicioStreaming,
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

func (ctrl *ControladorSpotify) ObtenerGeneros() []Genero {
	return []Genero{
		{ID: 1, Nombre: "Rock"},
		{ID: 2, Nombre: "Indie"},
		{ID: 3, Nombre: "Salsa"},
	}
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
func (ctrl *ControladorSpotify) ObtenerCancionesPorGenero(generoID int) []CancionDetalle {
	switch generoID {
	case 1: // Rock
		return []CancionDetalle{
			{ID: 1, Titulo: "Francés limón", Artista: "Enanitos Verdes", Album: "Amores lejanos", Anio: 2002, Duracion: "5:17", Genero: "Rock", ArchivoAudio: "frances_limon.mp3"},
			{ID: 2, Titulo: "Para no olvidar", Artista: "Los Rodriguez", Album: "Palabras más, palabras menos", Anio: 1995, Duracion: "3:56", Genero: "Rock", ArchivoAudio: "para_no_olvidar.mp3"},
			{ID: 3, Titulo: "Llueve sobre la ciudad", Artista: "Los Bunkers", Album: "Vida de perros", Anio: 2005, Duracion: "3:56", Genero: "Rock", ArchivoAudio: "llueve_sobre_la_ciudad.mp3"},
		}
	case 2: // Indie
		return []CancionDetalle{
			{ID: 4, Titulo: "Something about you", Artista: "Eyedress", Album: "Mulholland drive", Anio: 2021, Duracion: "2:33", Genero: "Indie", ArchivoAudio: "something_about_you.mp3"},
			{ID: 5, Titulo: "Blood was on my skin", Artista: "Club Hearts", Album: "Blood was on my skin", Anio: 2024, Duracion: "2:42", Genero: "Indie", ArchivoAudio: "blood_was_on_my_skin.mp3"},
			{ID: 6, Titulo: "Bad habit", Artista: "Steve Lacy", Album: "Gemini rights", Anio: 2022, Duracion: "3:53", Genero: "Indie", ArchivoAudio: "bad_habit.mp3"},
		}
	case 3: // Salsa
		return []CancionDetalle{
			{ID: 7, Titulo: "Si no la tengo", Artista: "Diablos Locos", Album: "Con calor y sentimiento", Anio: 2019, Duracion: "5:11", Genero: "Salsa", ArchivoAudio: "si_no_la_tengo.mp3"},
		}
	default:
		return []CancionDetalle{}
	}
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

// BuscarCancionEnServidor busca CUALQUIER canción en el servidor (SIN datos quemados)
func (ctrl *ControladorSpotify) BuscarCancionEnServidor(cancion *CancionDetalle) *RespuestaControlador {
	// Normalizar título para búsqueda (eliminar espacios y convertir a minúsculas)
	tituloNormalizado := ctrl.normalizarTitulo(cancion.Titulo)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	peticion := &pb.PeticionDTO{
		Titulo: tituloNormalizado,
	}

	fmt.Printf("\nLlamando método remoto: BuscarCancion(\"%s\")\n", tituloNormalizado)

	respuesta, err := ctrl.servicioCancion.BuscarCancion(ctx, peticion)
	if err != nil {
		return &RespuestaControlador{
			Exito:   false,
			Mensaje: fmt.Sprintf("Error de comunicación: %v", err),
			Datos:   nil,
		}
	}

	if respuesta.Codigo == 200 {
		// Enriquecer la canción con datos del servidor
		cancionEnriquecida := &CancionDetalle{
			ID:           respuesta.ObjCancion.Id,
			Titulo:       cancion.Titulo, // Mantener formato original con espacios
			Artista:      respuesta.ObjCancion.Artista_Banda,
			Album:        ctrl.obtenerAlbum(tituloNormalizado),
			Anio:         respuesta.ObjCancion.Lanzamiento,
			Duracion:     respuesta.ObjCancion.Duracion,
			Genero:       respuesta.ObjCancion.ObjGenero.Nombre,
			ArchivoAudio: ctrl.obtenerArchivoAudio(tituloNormalizado),
		}

		return &RespuestaControlador{
			Exito:   true,
			Mensaje: "Canción verificada en el servidor",
			Datos:   cancionEnriquecida,
		}
	}

	return &RespuestaControlador{
		Exito:   false,
		Mensaje: respuesta.Mensaje,
		Datos:   nil,
	}
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
func (ctrl *ControladorSpotify) ReproducirCancionConAudio(cancion *CancionDetalle) error {
	fmt.Printf("\nLlamando método remoto: EnviarCancionMedianteStream(\"%s\")\n", cancion.ArchivoAudio)

	// Crear contexto
	ctx := context.Background()

	// Crear petición gRPC
	peticion := &pbStreaming.PeticionDTO{
		Titulo: cancion.ArchivoAudio,
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
