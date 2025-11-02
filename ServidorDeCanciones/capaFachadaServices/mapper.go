package capaFachadaServices

import (
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
	dto "servidor/grpc-servidor/capaFachadaServices/DTO"
)

// MapearCancionAlmacenarDTOACancion convierte un CancionAlmacenarDTO a Cancion
// usando el servicio de géneros para obtener o crear el género
func MapearCancionAlmacenarDTOACancion(
	cancionDTO dto.CancionAlmacenarDTO,
	servicioGenero *ServiciosGenero,
) mo.Cancion {
	// Buscar o crear el género usando el servicio
	genero := servicioGenero.BuscarOCrearGenero(cancionDTO.Genero)

	// Crear y retornar la canción
	return mo.NewCancion(
		cancionDTO.Titulo,
		cancionDTO.Artista_Banda,
		cancionDTO.Lanzamiento,
		cancionDTO.Duracion,
		genero,
	)
}

// MapearListaCancionesAListaCancionRespuestaDTO convierte una lista de Cancion a lista de CancionRespuestaDTO
func MapearListaCancionesAListaCancionRespuestaDTO(canciones []mo.Cancion) []dto.CancionRespuestaDTO {
	var cancionesDTO []dto.CancionRespuestaDTO

	for _, cancion := range canciones {
		cancionDTO := dto.NuevaCancionRespuestaDTO(
			cancion.Titulo,
			cancion.Artista_Banda,
			cancion.Lanzamiento,
			cancion.Duracion,
			cancion.Ruta,
			cancion.Genero.Nombre,
		)
		cancionesDTO = append(cancionesDTO, cancionDTO)
	}

	return cancionesDTO
}

// MapearCancionRespuestaDTOACancion convierte un CancionRespuestaDTO a Cancion
// usando el servicio de géneros para obtener o crear el género
func MapearCancionRespuestaDTOACancion(
	cancionDTO dto.CancionRespuestaDTO,
	servicioGenero *ServiciosGenero,
) mo.Cancion {
	// Buscar o crear el género usando el servicio
	genero := servicioGenero.BuscarOCrearGenero(cancionDTO.Genero)

	// Crear la canción (nota: el ID no se asigna aquí)
	return mo.Cancion{
		Titulo:        cancionDTO.Titulo,
		Artista_Banda: cancionDTO.Artista_Banda,
		Lanzamiento:   cancionDTO.Lanzamiento,
		Duracion:      cancionDTO.Duracion,
		Ruta:          cancionDTO.Ruta,
		Genero:        genero,
	}
}

// MapearCancionACancionRespuestaDTO convierte una Cancion a CancionRespuestaDTO
func MapearCancionACancionRespuestaDTO(cancion mo.Cancion) dto.CancionRespuestaDTO {
	return dto.NuevaCancionRespuestaDTO(
		cancion.Titulo,
		cancion.Artista_Banda,
		cancion.Lanzamiento,
		cancion.Duracion,
		cancion.Ruta,
		cancion.Genero.Nombre,
	)
}
