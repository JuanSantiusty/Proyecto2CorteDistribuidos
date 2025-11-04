package main

import (
	"fmt"
	"net/http"
	controlador "reproducciones/capaControladores"
)

func main() {
	ctrl := controlador.NuevoControladorTendencias()

	http.HandleFunc("/reproducciones/listar", ctrl.ListarReproduccionesPorClienteHandler)
	http.HandleFunc("/reproducciones/almacenar", ctrl.RegistrarReproduccionHandler)

	fmt.Println(" Servicio de reproducciones escuchando en el puerto 5004...")
	if err := http.ListenAndServe(":5004", nil); err != nil {
		fmt.Println("Error iniciando el servidor: ", err)
	}
}
