package capaFachadaServices

import (
	"fmt"
	"reflect"
	rp "servidor/grpc-servidor/CapaAccesoDatos"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
	dto "servidor/grpc-servidor/capaFachadaServices/DTO"
	cola "servidor/grpc-servidor/conexionCola"
)

type ServiciosCancion struct {
	repo           *rp.CancionesRepo
	servicioGenero *ServiciosGenero
	colaRa         *cola.RabbitPublisher
}

func NuevoServicioCanciones() *ServiciosCancion {
	fmt.Print("Inicializando servicios canción")
	repo := rp.GetRepositorioCanciones()
	servicioGenero := NuevoServicioGeneros()
	colaRa, err := cola.NewRabbitPublisher()
	if err != nil {
		fmt.Printf("Advertencia: No se pudo conectar a RabbitMQ: %v\n", err)
	}
	return &ServiciosCancion{
		repo:           repo,
		servicioGenero: servicioGenero,
		colaRa:         colaRa,
	}
}

func (s *ServiciosCancion) AlmacenarCancion(cancionalmacenardto dto.CancionAlmacenarDTO, data []byte) error {
	cancionNueva := MapearCancionAlmacenarDTOACancion(cancionalmacenardto, s.servicioGenero)
	er := s.repo.AgregarCancion(cancionNueva, data)
	if er != nil {
		go s.colaRa.PublicarNotificacion(cola.NotificacionCancion{
			Titulo:  cancionNueva.Titulo,
			Artista: cancionNueva.Artista_Banda,
			Genero:  cancionNueva.Genero.Nombre,
			Mensaje: "Canción " + cancionNueva.Titulo + "Almacenada sin exito",
		})
	} else {
		go s.colaRa.PublicarNotificacion(cola.NotificacionCancion{
			Titulo:  cancionNueva.Titulo,
			Artista: cancionNueva.Artista_Banda,
			Genero:  cancionNueva.Genero.Nombre,
			Mensaje: "Canción " + cancionNueva.Titulo + "Almacenada con exito",
		})
	}
	return er
}

func (s *ServiciosCancion) BuscarCancion(titulo string) dto.RespuestaDTO[dto.CancionRespuestaDTO] {
	objvacio := mo.Cancion{}
	objCancion := s.repo.BuscarPorTitulo(titulo)

	if !reflect.DeepEqual(objvacio, objCancion) {
		fmt.Printf("Canción encontrada: %s - %s\n", objCancion.Artista_Banda, objCancion.Titulo)
		respuesta := MapearCancionACancionRespuestaDTO(objCancion)
		return dto.NewRespuestaDTO(respuesta, 200, "Canción encontrada")
	} else {
		fmt.Printf("Canción NO encontrada: %s\n", titulo)
		respuesta := dto.NuevaCancionRespuestaDTO("", "", 0, "", "", "")
		return dto.NewRespuestaDTO(respuesta, 404, "Canción no encontrada")
	}
}

func (s *ServiciosCancion) BuscarPorGenero(genero string) dto.RespuestaDTO[[]dto.CancionRespuestaDTO] {
	fmt.Printf("Buscando canciones del género: %s\n", genero)

	canciones := s.repo.BuscarPorGenero(genero)

	if len(canciones) > 0 {
		fmt.Printf("Encontradas %d canciones del género %s\n", len(canciones), genero)
		cancionesMap := MapearListaCancionesAListaCancionRespuestaDTO(canciones)
		return dto.NewRespuestaDTO(cancionesMap, 200, "canciones encontradas")
	} else {
		fmt.Printf("No se encontraron canciones del género %s\n", genero)
		return dto.NewRespuestaDTO([]dto.CancionRespuestaDTO{}, 404, fmt.Sprintf("No se encontraron canciones del género %s", genero))
	}
}

func (s *ServiciosCancion) ListarCanciones() dto.RespuestaDTO[[]dto.CancionRespuestaDTO] {
	canciones := MapearListaCancionesAListaCancionRespuestaDTO(s.repo.ListaCanciones())
	return dto.NewRespuestaDTO(canciones, 200, "Canciones Encontradas")
}
