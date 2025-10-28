package vistaCliente

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	ctrl "cliente/capaControladores"
)

// VistaSpotify maneja toda la interfaz de usuario
type VistaSpotify struct {
	controlador *ctrl.ControladorSpotify
	reader      *bufio.Reader
}

// NewVistaSpotify crea nueva instancia de la vista
func NewVistaSpotify(controlador *ctrl.ControladorSpotify) *VistaSpotify {
	return &VistaSpotify{
		controlador: controlador,
		reader:      bufio.NewReader(os.Stdin),
	}
}

// IniciarAplicacion punto de entrada principal
func (v *VistaSpotify) IniciarAplicacion() {
	for {
		if !v.mostrarMenuPrincipal() {
			break
		}
	}
}

// ===== MENÚ PRINCIPAL =====

func (v *VistaSpotify) mostrarMenuPrincipal() bool {
	fmt.Println("\nSpotify")
	fmt.Println("1. Ver géneros")
	fmt.Println("2. Salir")

	opcion := v.leerEntrada("")

	numero, err := v.controlador.ValidarOpcionMenuPrincipal(opcion)
	if err != nil {
		v.mostrarError(err.Error())
		return true
	}

	accion := v.controlador.ProcesarOpcionMenuPrincipal(numero)

	switch accion {
	case "VER_GENEROS":
		v.mostrarMenuGeneros()
		return true
	case "SALIR":
		return false
	default:
		v.mostrarError("Opción no válida")
		return true
	}
}

// ===== MENÚ DE GÉNEROS =====

func (v *VistaSpotify) mostrarMenuGeneros() {
	for {
		generos := v.controlador.ObtenerGeneros()

		fmt.Println("\nSpotify")
		for _, genero := range generos {
			fmt.Printf("%d. %s\n", genero.ID, genero.Nombre)
		}
		fmt.Printf("%d. Atrás\n", len(generos)+1)

		opcion := v.leerEntrada("")

		numero, err := v.controlador.ValidarOpcionGenero(opcion, len(generos))
		if err != nil {
			v.mostrarError(err.Error())
			continue
		}

		if numero == len(generos)+1 {
			return
		}

		v.mostrarCancionesPorGenero(numero)
	}
}

// ===== MENÚ DE CANCIONES POR GÉNERO =====

func (v *VistaSpotify) mostrarCancionesPorGenero(generoID int) {
	generos := v.controlador.ObtenerGeneros()
	var nombreGenero string

	for _, g := range generos {
		if g.ID == generoID {
			nombreGenero = g.Nombre
			break
		}
	}

	for {
		canciones := v.controlador.ObtenerCancionesPorGenero(generoID)

		fmt.Printf("\nSpotify\n")
		fmt.Printf("Género: %s\n", nombreGenero)

		for i, cancion := range canciones {
			fmt.Printf("%d. %s - %s\n", i+1, cancion.Artista, cancion.Titulo)
		}
		fmt.Printf("%d. Atrás\n", len(canciones)+1)

		opcion := v.leerEntrada("")

		numero, err := v.controlador.ValidarOpcionCancion(opcion, len(canciones))
		if err != nil {
			v.mostrarError(err.Error())
			continue
		}

		if numero == len(canciones)+1 {
			return
		}

		cancionSeleccionada := canciones[numero-1]
		v.mostrarDetallesCancion(&cancionSeleccionada)
	}
}

// ===== DETALLES DE CANCIÓN =====

func (v *VistaSpotify) mostrarDetallesCancion(cancion *ctrl.CancionDetalle) {

	//fmt.Println("Verificando canción en el servidor...")
	respuesta := v.controlador.BuscarCancionEnServidor(cancion)

	if respuesta.Exito {
		// Actualizar con datos del servidor si se encontró
		if cancionServidor, ok := respuesta.Datos.(*ctrl.CancionDetalle); ok {
			cancion = cancionServidor
			v.mostrarMensaje(" " + respuesta.Mensaje)
		}
	} else {
		// Si no se encuentra en el servidor, usar datos locales
		v.mostrarMensaje("Usando datos locales (servidor no respondió)")
	}

	for {
		fmt.Printf("\nSpotify\n")
		fmt.Printf("Canción: %s - %s\n", cancion.Artista, cancion.Titulo)
		fmt.Printf("* Título de la canción: %s\n", cancion.Titulo)
		fmt.Printf("* Artista/Banda: %s\n", cancion.Artista)
		fmt.Printf("* Álbum: %s\n", cancion.Album)
		fmt.Printf("* Año de lanzamiento: %d\n", cancion.Anio)
		fmt.Printf("* Duración: %s\n", cancion.Duracion)
		fmt.Println("1. Reproducir")
		fmt.Println("2. Salir")

		opcion := v.leerEntrada("")

		numero, err := v.controlador.ValidarOpcionDetalle(opcion)
		if err != nil {
			v.mostrarError(err.Error())
			continue
		}

		switch numero {
		case 1:
			v.mostrarStreamingReal(cancion)
		case 2:
			return
		}
	}
}

// ===== STREAMING REAL =====

func (v *VistaSpotify) mostrarStreamingReal(cancion *ctrl.CancionDetalle) {
	fmt.Printf("\nSpotify\n")
	fmt.Printf("Canción: %s - %s\n", cancion.Artista, cancion.Titulo)
	fmt.Println("Reproduciendo canción...")

	// Iniciar streaming y reproducción real
	err := v.controlador.ReproducirCancionConAudio(cancion)
	if err != nil {
		v.mostrarError(fmt.Sprintf("Error en streaming: %v", err))
		return
	}

	// Preguntar si desea continuar
	fmt.Println("\nPresione Enter para volver...")
	v.leerEntrada("")
}

// mostrarProgreso muestra el progreso del streaming en tiempo real
func (v *VistaSpotify) mostrarProgreso(progresoChan chan int, errorChan chan error) {
	totalBytes := 0
	fragmentos := 0

	fmt.Println("Progreso del streaming:")

	for {
		select {
		case bytes, ok := <-progresoChan:
			if !ok {
				fmt.Printf("\n Streaming completado: %d fragmentos, %d bytes totales\n", fragmentos, totalBytes)
				return
			}
			fragmentos++
			totalBytes += bytes
			v.actualizarBarraProgreso(fragmentos, bytes)

		case err := <-errorChan:
			fmt.Printf("\n Error en streaming: %v\n", err)
			return

		case <-time.After(30 * time.Second):
			fmt.Println("\n Timeout del streaming")
			return
		}
	}
}

// actualizarBarraProgreso muestra una barra de progreso visual
func (v *VistaSpotify) actualizarBarraProgreso(fragmento int, bytes int) {
	// Limpiar línea anterior
	fmt.Print("\r")

	// Mostrar barra de progreso
	barraLength := 30
	progreso := fragmento % barraLength

	barra := "["
	for i := 0; i < barraLength; i++ {
		if i <= progreso {
			barra += "█"
		} else {
			barra += "░"
		}
	}
	barra += "]"

	fmt.Printf(" %s Fragmento #%d (%d bytes) ", barra, fragmento, bytes)
}

// ===== FUNCIONES AUXILIARES =====

func (v *VistaSpotify) leerEntrada(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}
	entrada, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(entrada)
}

func (v *VistaSpotify) mostrarError(mensaje string) {
	fmt.Printf(" %s\n", mensaje)
}

func (v *VistaSpotify) mostrarMensaje(mensaje string) {
	fmt.Printf("%s\n", mensaje)
}
