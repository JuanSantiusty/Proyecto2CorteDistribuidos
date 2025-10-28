package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	rp "servidor/grpc-servidor/CapaAccesoDatos"
	ctrl "servidor/grpc-servidor/capaControladores"
	pb "servidor/grpc-servidor/serviciosCancion"
)

func main() {
	// Configurar el listener
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Error al abrir el puerto: %v", err)
	}

	// Inicializar repositorio
	repo := rp.NewCancionesRepo()

	// Crear controlador con inyección de dependencias
	controladorCanciones := ctrl.NewControladorCanciones(repo)

	// Inicializar datos
	controladorCanciones.InicializarDatos()

	// Crear servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar el servicio usando el controlador
	pb.RegisterServiciosCancionesServer(grpcServer, controladorCanciones)

	// Iniciar el servidor
	log.Println("Servidor gRPC escuchando en puerto 50053...")
	log.Println("Controladores inicializados correctamente")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
