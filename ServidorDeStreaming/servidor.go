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

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error escuchando en el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Se registra el controlador que ofrece el procedimiento remoto
	pb.RegisterAudioServiceServer(grpcServer, &capacontroladores.ControladorServidor{})

	fmt.Println("Servidor gRPC escuchando en :50051...")

	// Iniciar el servidor
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar servidor gRPC: %v", err)
	}
}
