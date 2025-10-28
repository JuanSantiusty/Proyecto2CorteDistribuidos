package capaControladores

import (
	"context"
	"log"

	rp "servidor/grpc-servidor/CapaAccesoDatos"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
	se "servidor/grpc-servidor/capaFachadaServices"
	pb "servidor/grpc-servidor/serviciosCancion"
)

// ControladorCanciones maneja las peticiones relacionadas con canciones
type ControladorCanciones struct {
	pb.UnimplementedServiciosCancionesServer
	repo *rp.CancionesRepo
}

// NewControladorCanciones crea una nueva instancia del controlador
func NewControladorCanciones(repo *rp.CancionesRepo) *ControladorCanciones {
	return &ControladorCanciones{
		repo: repo,
	}
}

// BuscarCancion maneja la petición de búsqueda de canción
func (c *ControladorCanciones) BuscarCancion(ctx context.Context, req *pb.PeticionDTO) (*pb.RespuestaCancionDTO, error) {
	// Validar entrada
	if req == nil {
		log.Println("Error: Petición nula recibida")
		return c.crearRespuestaError("Petición inválida", 400), nil
	}

	titulo := req.GetTitulo()
	if titulo == "" {
		log.Println("Error: Título vacío en la petición")
		return c.crearRespuestaError("Título requerido", 400), nil
	}

	// Log de la operación
	log.Printf("Buscando canción con título: %s", titulo)

	// Llamar al servicio de la capa de fachada
	resp := se.BuscarCancion(titulo, c.repo)

	// Convertir la respuesta del modelo interno a DTO de protobuf
	respuestaDTO := c.convertirAProtobufDTO(resp)

	log.Printf("Búsqueda completada. Código de respuesta: %d", respuestaDTO.Codigo)
	return respuestaDTO, nil
}

// BuscarPorGenero maneja la búsqueda de canciones por género
func (c *ControladorCanciones) BuscarPorGenero(ctx context.Context, req *pb.PeticionGeneroDTO) (*pb.RespuestaListaCancionesDTO, error) {
	// Validar entrada
	if req == nil {
		log.Println("Error: Petición nula recibida para búsqueda por género")
		return c.crearRespuestaListaError("Petición inválida", 400), nil
	}

	genero := req.GetGenero()
	if genero == "" {
		log.Println("Error: Género vacío en la petición")
		return c.crearRespuestaListaError("Género requerido", 400), nil
	}

	log.Printf("Buscando canciones del género: %s", genero)

	// Llamar al servicio de búsqueda por género
	resp := se.BuscarPorGenero(genero, c.repo)

	// Convertir respuesta a DTO de protobuf
	respuestaDTO := c.convertirAProtobufListaDTO(resp)

	log.Printf("Búsqueda por género completada. Código: %d", respuestaDTO.Codigo)
	return respuestaDTO, nil
}

// convertirAProtobufDTO convierte la respuesta interna a formato protobuf
func (c *ControladorCanciones) convertirAProtobufDTO(resp mo.RespuestaDTO[mo.Cancion]) *pb.RespuestaCancionDTO {
	respuesta := &pb.RespuestaCancionDTO{
		Codigo:  resp.Codigo,
		Mensaje: resp.Mensaje,
	}

	// Solo llenar los datos si la respuesta fue exitosa
	if resp.Codigo == 200 && resp.Data.Id != 0 {
		respuesta.ObjCancion = &pb.Cancion{
			Id:            resp.Data.Id,
			Titulo:        resp.Data.Titulo,
			Artista_Banda: resp.Data.Artista_Banda,
			Lanzamiento:   resp.Data.Lanzamiento,
			Duracion:      resp.Data.Duracion,
			ObjGenero: &pb.Genero{
				Id:     resp.Data.Genero.Id,
				Nombre: resp.Data.Genero.Nombre,
			},
		}
	}

	return respuesta
}

// convertirAProtobufListaDTO convierte lista de canciones a formato protobuf
func (c *ControladorCanciones) convertirAProtobufListaDTO(resp mo.RespuestaDTO[[]mo.Cancion]) *pb.RespuestaListaCancionesDTO {
	respuesta := &pb.RespuestaListaCancionesDTO{
		Codigo:  resp.Codigo,
		Mensaje: resp.Mensaje,
	}

	// Solo llenar los datos si la respuesta fue exitosa
	if resp.Codigo == 200 && len(resp.Data) > 0 {
		respuesta.Canciones = make([]*pb.Cancion, len(resp.Data))

		for i, cancion := range resp.Data {
			respuesta.Canciones[i] = &pb.Cancion{
				Id:            cancion.Id,
				Titulo:        cancion.Titulo,
				Artista_Banda: cancion.Artista_Banda,
				Lanzamiento:   cancion.Lanzamiento,
				Duracion:      cancion.Duracion,
				ObjGenero: &pb.Genero{
					Id:     cancion.Genero.Id,
					Nombre: cancion.Genero.Nombre,
				},
			}
		}
	}

	return respuesta
}

// crearRespuestaError crea una respuesta de error estándar
func (c *ControladorCanciones) crearRespuestaError(mensaje string, codigo int32) *pb.RespuestaCancionDTO {
	return &pb.RespuestaCancionDTO{
		Codigo:     codigo,
		Mensaje:    mensaje,
		ObjCancion: nil,
	}
}

// crearRespuestaListaError crea una respuesta de error para listas
func (c *ControladorCanciones) crearRespuestaListaError(mensaje string, codigo int32) *pb.RespuestaListaCancionesDTO {
	return &pb.RespuestaListaCancionesDTO{
		Codigo:    codigo,
		Mensaje:   mensaje,
		Canciones: nil,
	}
}

// InicializarDatos carga los datos iniciales
func (c *ControladorCanciones) InicializarDatos() {
	log.Println("Inicializando datos de canciones...")
	se.CargarCanciones(c.repo)
	log.Println("Datos de canciones cargados exitosamente")
}
