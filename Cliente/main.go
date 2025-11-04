package main

import (
	"fmt"
	"log"

	ctrl "cliente/capaControladores"
	inf "cliente/infraestructura"
	vista "cliente/vistaCliente"

	"google.golang.org/grpc"
)

func main() {
	//fmt.Println("Iniciando Spotify...")

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
	return ctrl.NewControladorSpotify(servicioCancion, serviciosGenero, servicioStreaming)
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
