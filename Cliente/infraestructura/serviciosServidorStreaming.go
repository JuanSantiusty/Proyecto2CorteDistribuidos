package infraestructura

import (
	"context"

	pb "servidor-streaming/serviciosStreaming"

	"google.golang.org/grpc"
)

// StreamingGRPC implementa ServicioStreaming usando gRPC
type StreamingGRPC struct {
	streaming pb.AudioServiceClient
}

// NewClienteStreamingGRPC crea un nuevo cliente gRPC para streaming
func NewClienteStreamingGRPC(conn *grpc.ClientConn) ServicioStreaming {
	return &StreamingGRPC{
		streaming: pb.NewAudioServiceClient(conn),
	}
}

// EnviarCancionMedianteStream realiza la llamada gRPC para streaming
func (s *StreamingGRPC) EnviarCancionMedianteStream(ctx context.Context, req *pb.PeticionDTO) (pb.AudioService_EnviarCancionMedianteStreamClient, error) {
	return s.streaming.EnviarCancionMedianteStream(ctx, req)
}
