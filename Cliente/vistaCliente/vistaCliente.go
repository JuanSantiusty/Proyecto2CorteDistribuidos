package vistaCliente

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			fmt.Printf("%d. %s\n", genero.Id, genero.Nombre)
		}
		fmt.Printf("%d. Atrás\n", len(generos)+1)

		opcion := v.leerEntrada("")
		indice, _ := strconv.Atoi(opcion)

		if indice > len(generos) {
			break
		}

		v.mostrarCancionesPorGenero(generos[indice-1].Nombre)
	}
}

// ===== MENÚ DE CANCIONES POR GÉNERO =====

func (v *VistaSpotify) mostrarCancionesPorGenero(genero string) {
	for {
		canciones := v.controlador.ObtenerCancionesPorGenero(genero)

		fmt.Printf("\nSpotify\n")
		fmt.Printf("Género: %s\n", genero)

		for i, cancion := range canciones {
			fmt.Printf("%d. %s - %s\n", i+1, cancion.Artista_Banda, cancion.Titulo)
		}
		fmt.Printf("%d. Atrás\n", len(canciones)+1)

		opcion := v.leerEntrada("")
		indice, _ := strconv.Atoi(opcion)

		if indice > len(canciones) {
			break
		}

		v.mostrarDetallesCancion(canciones[indice-1].Titulo)
	}
}

// ===== DETALLES DE CANCIÓN =====

func (v *VistaSpotify) mostrarDetallesCancion(cancionTitulo string) {

	//fmt.Println("Verificando canción en el servidor...")
	respuesta := v.controlador.BuscarCancionEnServidor(cancionTitulo)

	if respuesta.Codigo == 200 {
		fmt.Printf(respuesta.Mensaje)
		cancion := respuesta.Data
		for {
			fmt.Printf("\nSpotify\n")
			fmt.Printf("Canción: %s - %s\n", cancion.Artista_Banda, cancion.Titulo)
			fmt.Printf("* Título de la canción: %s\n", cancion.Titulo)
			fmt.Printf("* Artista/Banda: %s\n", cancion.Artista_Banda)
			fmt.Printf("* Año de lanzamiento: %d\n", cancion.Lanzamiento)
			fmt.Printf("* Idioma: %s\n", cancion.Idioma)
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
				v.mostrarStreamingReal(cancion.Titulo)
			case 2:
				return
			}
		}
	} else {
		fmt.Printf(respuesta.Mensaje)
	}
}

// ===== STREAMING REAL =====

func (v *VistaSpotify) mostrarStreamingReal(cancion string) {
	fmt.Printf("\nSpotify\n")
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
