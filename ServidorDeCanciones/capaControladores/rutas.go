package capaControladores

import (
	"net/http"
)

func ConfigurarRutas() {
	controladorCancion := NuevoControladorCancion()
	controladorGenero := NuevoControladorGenero()

	// Configurar rutas para canciones
	http.HandleFunc("/canciones", controladorCancion.ManejarCanciones)
	http.HandleFunc("/canciones/buscar", controladorCancion.ManejarCanciones)

	// Configurar rutas para géneros
	http.HandleFunc("/generos", controladorGenero.ManejarGeneros)
}
