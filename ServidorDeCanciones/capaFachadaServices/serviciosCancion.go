package capaFachadaServices

import (
	"fmt"
	"reflect"
	rp "servidor/grpc-servidor/CapaAccesoDatos"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
)

// CargarCanciones carga TODAS las canciones disponibles en el sistema
func CargarCanciones(r *rp.CancionesRepo) {
	// Género Rock
	generoRock := mo.NewGenero(1, "rock")
	r.AgregarCancion(mo.NewCancion(1, "franceslimon", "Enanitos Verdes", 2002, "5:17", generoRock))
	r.AgregarCancion(mo.NewCancion(2, "paranoolvidar", "Los Rodriguez", 1995, "3:56", generoRock))
	r.AgregarCancion(mo.NewCancion(3, "lluevesobrelaciudad", "Los Bunkers", 2005, "3:56", generoRock))

	// Género Indie
	generoIndie := mo.NewGenero(2, "indie")
	r.AgregarCancion(mo.NewCancion(4, "somethingaboutyou", "Eyedress", 2021, "2:33", generoIndie))
	r.AgregarCancion(mo.NewCancion(5, "bloodwasonmyskin", "Club Hearts", 2024, "2:42", generoIndie))
	r.AgregarCancion(mo.NewCancion(6, "badhabit", "Steve Lacy", 2022, "3:53", generoIndie))

	// Género Salsa
	generoSalsa := mo.NewGenero(3, "salsa")
	r.AgregarCancion(mo.NewCancion(7, "sinolatengo", "Diablos Locos", 2019, "5:11", generoSalsa))

	fmt.Println("Canciones cargadas exitosamente")
}

func BuscarCancion(titulo string, r *rp.CancionesRepo) mo.RespuestaDTO[mo.Cancion] {
	objvacio := mo.Cancion{}
	objCancion := r.BuscarPorTitulo(titulo)

	if !reflect.DeepEqual(objvacio, objCancion) {
		fmt.Printf("Canción encontrada: %s - %s\n", objCancion.Artista_Banda, objCancion.Titulo)
		return mo.NewRespuestaDTO(objCancion, 200, "Canción encontrada")
	} else {
		fmt.Printf("Canción NO encontrada: %s\n", titulo)
		return mo.NewRespuestaDTO(objCancion, 404, "Canción no encontrada")
	}
}

func BuscarPorGenero(genero string, r *rp.CancionesRepo) mo.RespuestaDTO[[]mo.Cancion] {
	fmt.Printf("Buscando canciones del género: %s\n", genero)

	canciones := r.BuscarPorGenero(genero)

	if len(canciones) > 0 {
		fmt.Printf("Encontradas %d canciones del género %s\n", len(canciones), genero)
		return mo.NewRespuestaDTO(canciones, 200, fmt.Sprintf("Se encontraron %d canciones del género %s", len(canciones), genero))
	} else {
		fmt.Printf("No se encontraron canciones del género %s\n", genero)
		return mo.NewRespuestaDTO([]mo.Cancion{}, 404, fmt.Sprintf("No se encontraron canciones del género %s", genero))
	}
}
