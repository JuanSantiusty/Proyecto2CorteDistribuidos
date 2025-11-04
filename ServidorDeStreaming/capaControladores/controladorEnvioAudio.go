// ServidorDeStreaming/capaControladores/controladorEnvioAudio.go
package capaControladores

import (
	"context"
	"fmt"

	"servidor-streaming/capaComunicacionExterna"
	capafachadaservices "servidor-streaming/capaFachadaServices"
	pb "servidor-streaming/serviciosStreaming"

	"google.golang.org/grpc/peer"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
	clienteReproducciones *capaComunicacionExterna.ClienteReproducciones
	servicioCanciones     *capafachadaservices.FachadaBuscarCancion
}

// NewControladorServidor crea una nueva instancia del controlador
func NewControladorServidor(url string) *ControladorServidor {
	return &ControladorServidor{
		clienteReproducciones: capaComunicacionExterna.NewClienteReproducciones(),
		servicioCanciones:     capafachadaservices.NewFachadaBuscarCancion(url),
	}
}

// EnviarCancionMedianteStream implementación del procedimiento remoto
func (s *ControladorServidor) EnviarCancionMedianteStream(req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {
	// Obtener información del cliente
	direccionCliente := obtenerDireccionCliente(stream.Context())

	fmt.Printf("📥 Echo: EnviarCancionMedianteStream - Canción: %s, Cliente: %s\n", req.Titulo, direccionCliente)

	// Extraer datos para registro de reproducción
	titulo := req.Titulo
	can, _ := s.servicioCanciones.BuscarCancionPorTitulo(titulo)
	usuarioId := req.Formato

	// COMUNICACIÓN ASÍNCRONA: Registrar reproducción en el servidor de reproducciones
	fmt.Printf("📡 Enviando reproducción asíncronamente al Servidor de Reproducciones...\n")
	s.clienteReproducciones.RegistrarReproduccionAsincrona(int(can.ID), titulo, usuarioId)

	// Continuar con el streaming sin esperar respuesta
	return capafachadaservices.StreamAudioFile(
		can.Ruta, func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		})
}

// obtenerDireccionCliente extrae la dirección IP del cliente
func obtenerDireccionCliente(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}
	return "Desconocido"
}
