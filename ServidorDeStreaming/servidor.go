// ServidorDeStreaming/servidor.go
package main

import (
	"fmt"
	"log"
	"net"

	capacontroladores "servidor-streaming/capaControladores"
	pb "servidor-streaming/serviciosStreaming"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("🚀 Iniciando Servidor de Streaming...")
	fmt.Println("==========================================")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error escuchando en el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Crear controlador con cliente de reproducciones
	controlador := capacontroladores.NewControladorServidor("http://localhost:5000")

	// Registrar el controlador que ofrece el procedimiento remoto
	pb.RegisterAudioServiceServer(grpcServer, controlador)

	fmt.Println("🌐 Servidor gRPC escuchando en :50051...")
	fmt.Println("🔗 Conectado a Servidor de Reproducciones: http://localhost:5004")
	fmt.Println("==========================================")

	// Iniciar el servidor
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Error al iniciar servidor gRPC: %v", err)
	}
}
