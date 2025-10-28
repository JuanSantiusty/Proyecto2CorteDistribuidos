package infraestructura

import (
	"context"
	pb "servidor/grpc-servidor/serviciosCancion"

	"google.golang.org/grpc"
)

// CancionesGRPC implementa ServicioCanciones usando gRPC real
type CancionesGRPC struct {
	canciones pb.ServiciosCancionesClient
}

func NewClienteGRPC(conn *grpc.ClientConn) ServicioCanciones {
	return &CancionesGRPC{
		canciones: pb.NewServiciosCancionesClient(conn),
	}
}

// Implementacion de la funcoin BuscarCancion de la interface
func (c *CancionesGRPC) BuscarCancion(ctx context.Context, request *pb.PeticionDTO) (*pb.RespuestaCancionDTO, error) {
	return c.canciones.BuscarCancion(ctx, request)
}
