package infraestructura

import (
	"context"
	pb "servidor/grpc-servidor/serviciosCancion"
)

// Interfaz para el servicio de canciones
type ServicioCanciones interface {
	BuscarCancion(ctx context.Context, request *pb.PeticionDTO) (*pb.RespuestaCancionDTO, error)
}
