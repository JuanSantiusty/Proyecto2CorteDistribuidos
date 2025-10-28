package infraestructura

import (
	"context"
	pb "servidor-streaming/serviciosStreaming"
)

// ServicioStreaming interfaz para el servicio de streaming
type ServicioStreaming interface {
	// Método que retorna el stream directamente
	EnviarCancionMedianteStream(ctx context.Context, req *pb.PeticionDTO) (pb.AudioService_EnviarCancionMedianteStreamClient, error)
}
