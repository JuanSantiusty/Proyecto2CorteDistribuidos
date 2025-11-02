package main

import (
	"fmt"
	"log"
	"net/http"
	con "servidor/grpc-servidor/capaControladores"
)

func main() {
	// Configurar todas las rutas
	con.ConfigurarRutas()

	// Iniciar el servidor
	puerto := ":5000"
	fmt.Printf("Servidor iniciado en http://localhost%s\n", puerto)
	fmt.Println("Endpoints disponibles:")
	fmt.Println("  POST   /canciones")
	fmt.Println("  GET    /canciones")
	fmt.Println("  GET    /canciones/buscar?titulo=...")
	fmt.Println("  GET    /canciones/buscar?genero=...")
	fmt.Println("  GET    /generos")

	log.Fatal(http.ListenAndServe(puerto, nil))
}
