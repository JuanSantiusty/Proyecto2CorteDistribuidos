// ServidorDeStreaming/capaComunicacionExterna/clienteReproducciones.go
package capaComunicacionExterna

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ClienteReproducciones cliente HTTP para conectarse al servidor de reproducciones
type ClienteReproducciones struct {
	baseURL    string
	httpClient *http.Client
}

// ReproduccionDTO estructura para enviar datos de reproducción
type ReproduccionDTO struct {
	CancionId int    `json:"cancionId"`
	Titulo    string `json:"titulo"`
	UsuarioId string `json:"usuarioId"`
}

// NewClienteReproducciones crea una nueva instancia del cliente
func NewClienteReproducciones() *ClienteReproducciones {
	return &ClienteReproducciones{
		baseURL: "http://localhost:5004", // Puerto del servidor de reproducciones
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// RegistrarReproduccionAsincrona envía la reproducción de manera asíncrona
func (c *ClienteReproducciones) RegistrarReproduccionAsincrona(cancionId int, titulo string, usuarioId string) {
	// Ejecutar en goroutine para no bloquear el streaming
	go func() {
		err := c.registrarReproduccion(cancionId, titulo, usuarioId)
		if err != nil {
			fmt.Printf("⚠️ Error registrando reproducción (asíncrono): %v\n", err)
		} else {
			fmt.Printf("✅ Reproducción registrada asíncronamente: Canción='%s', Usuario=%d\n", titulo, usuarioId)
		}
	}()
}

// registrarReproduccion envía los datos al servidor de reproducciones
func (c *ClienteReproducciones) registrarReproduccion(cancionId int, titulo string, usuarioId string) error {
	url := fmt.Sprintf("%s/reproducciones/almacenar", c.baseURL)

	// Crear DTO
	dto := ReproduccionDTO{
		CancionId: cancionId,
		Titulo:    titulo,
		UsuarioId: usuarioId,
	}

	// Convertir a JSON
	jsonData, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("error al serializar JSON: %v", err)
	}

	// Hacer petición POST
	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error al enviar petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("servidor respondió con código: %d", resp.StatusCode)
	}

	return nil
}
