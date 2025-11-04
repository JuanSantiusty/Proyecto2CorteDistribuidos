package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	ctrl "cliente/capaControladores"
	inf "cliente/infraestructura"
	vista "cliente/vistaCliente"

	"google.golang.org/grpc"
)

// Variable global para el usuario autenticado
var usuarioActual string

func main() {
	//fmt.Println("Iniciando Spotify...")

	//Inicio de sesion
	// Crear fachada de autenticación
	fachadaAuth := inf.NewFachadaAutenticacion("usuarios.txt")

	fmt.Println("🎵 BIENVENIDO AL SISTEMA DE MÚSICA 🎵")
	fmt.Println("======================================")

	salir := false
	for !salir {
		// Menú principal
		fmt.Println("\n¿Qué deseas hacer?")
		fmt.Println("1. Registrarse")
		fmt.Println("2. Iniciar sesión")
		fmt.Println("3. Salir")
		fmt.Print("Selecciona una opción (1-3): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			registrarUsuario(fachadaAuth)
		case "2":
			if iniciarSesion(fachadaAuth) {
				salir = true
			}
		case "3":
			fmt.Println("¡Hasta pronto! 👋")
			os.Exit(0)
		default:
			fmt.Println("❌ Opción no válida. Por favor selecciona 1, 2 o 3.")
		}
	}

	if err := inicializarAplicacion(); err != nil {
		log.Fatalf("Error al inicializar la aplicación: %v", err)
	}
}

// inicializarAplicacion configura todas las dependencias e inicia la app
func inicializarAplicacion() error {
	// Establecer conexiones con ambos servidores
	connStreaming, err := establecerConexiones()
	if err != nil {
		return err
	}
	defer connStreaming.Close()

	// Configurar servicios
	servicioCancion, serviciosGenero := configurarServicioCancion("http://localhost:5000")
	servicioStreaming := configurarServicioStreaming(connStreaming)

	// Configurar capas de la aplicación
	controlador := configurarControlador(servicioCancion, serviciosGenero, servicioStreaming)
	vista := configurarVista(controlador)

	// Iniciar aplicación
	ejecutarAplicacion(vista)

	return nil
}

// establecerConexiones establece conexiones con ambos servidores
func establecerConexiones() (*grpc.ClientConn, error) {

	// Conexión al servidor de streaming (puerto 50051)
	//fmt.Println("Conectando al servidor de streaming (localhost:50051)...")
	connStreaming, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("error conectando al servidor de streaming: %v", err)
	}
	//fmt.Println("Conexión al servidor de streaming establecida")

	return connStreaming, nil
}

// configurarServicioCancion configura el servicio de canciones
func configurarServicioCancion(url string) (inf.FachadaCanciones, inf.FachadaGeneros) {
	//fmt.Println("Configurando servicio de canciones...")
	return *inf.NewFachadaCanciones(url), *inf.NewFachadaGeneros(url)
}

// configurarServicioStreaming configura el servicio de streaming
func configurarServicioStreaming(conn *grpc.ClientConn) inf.ServicioStreaming {
	//fmt.Println("Configurando servicio de streaming...")
	return inf.NewClienteStreamingGRPC(conn)
}

// configurarControlador configura la capa de controladores con ambos servicios
func configurarControlador(servicioCancion inf.FachadaCanciones, serviciosGenero inf.FachadaGeneros, servicioStreaming inf.ServicioStreaming) *ctrl.ControladorSpotify {
	//fmt.Println("Configurando controlador de aplicación...")
	return ctrl.NewControladorSpotify(servicioCancion, serviciosGenero, servicioStreaming, usuarioActual)
}

// configurarVista configura la capa de vista
func configurarVista(controlador *ctrl.ControladorSpotify) *vista.VistaSpotify {
	//fmt.Println("Configurando interfaz de usuario...")
	return vista.NewVistaSpotify(controlador)
}

// ejecutarAplicacion inicia el bucle principal de la aplicación
func ejecutarAplicacion(vistaApp *vista.VistaSpotify) {
	fmt.Println("Iniciando Spotify - Sistema de Música")
	/*fmt.Println("=====================================")
	fmt.Println("Servidor de Canciones: Conectado (puerto 50053)")
	fmt.Println("Servidor de Streaming: Conectado (puerto 50051)")
	fmt.Println("=====================================")*/

	// Iniciar el bucle principal de la interfaz
	vistaApp.IniciarAplicacion()

	fmt.Println("Conexiones cerradas correctamente")
}

// registrarUsuario maneja el proceso de registro
func registrarUsuario(fachadaAuth *inf.FachadaAutenticacion) {
	fmt.Println("\n--- REGISTRO DE USUARIO ---")

	scanner := bufio.NewScanner(os.Stdin)

	// Solicitar username
	fmt.Print("Ingresa un nombre de usuario: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	if username == "" {
		fmt.Println("❌ El nombre de usuario no puede estar vacío.")
		return
	}

	// Solicitar contraseña
	fmt.Print("Ingresa una contraseña: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	if len(password) < 4 {
		fmt.Println("❌ La contraseña debe tener al menos 4 caracteres.")
		return
	}
	// Intentar registrar el usuario
	respuesta := fachadaAuth.RegistrarUsuario(username, password)

	if respuesta.Exito {
		fmt.Println("✅ " + respuesta.Mensaje)
		fmt.Printf("🎉 ¡Bienvenido %s! Tu cuenta ha sido creada exitosamente.\n", username)
		return
	} else {
		fmt.Println("❌ " + respuesta.Mensaje)
	}
}

// iniciarSesion maneja el proceso de inicio de sesión
func iniciarSesion(fachadaAuth *inf.FachadaAutenticacion) bool {
	fmt.Println("\n--- INICIO DE SESIÓN ---")

	scanner := bufio.NewScanner(os.Stdin)

	// Solicitar username
	fmt.Print("Usuario: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	if username == "" {
		fmt.Println("❌ El nombre de usuario no puede estar vacío.")
		return false
	}

	// Solicitar contraseña
	fmt.Print("Contraseña: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	// Intentar iniciar sesión
	respuesta := fachadaAuth.IniciarSesion(username, password)

	if respuesta.Exito {
		fmt.Println("\n✅ " + respuesta.Mensaje)
		fmt.Printf("🎵 ¡Bienvenido de nuevo %s!\n", username)

		// Guardar el usuario en la variable global
		usuarioActual = username
	} else {
		fmt.Println("❌ " + respuesta.Mensaje)
	}
	return respuesta.Exito
}
