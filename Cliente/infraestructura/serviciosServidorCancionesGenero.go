package infraestructura

import (
	dto "cliente/infraestructura/DTO"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// FachadaGeneros proporciona métodos para consumir los servicios de géneros
type FachadaGeneros struct {
	baseURL string
	client  *http.Client
}

// NewFachadaGeneros crea una nueva instancia de la fachada de géneros
func NewFachadaGeneros(baseURL string) *FachadaGeneros {
	return &FachadaGeneros{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		client:  &http.Client{},
	}
}

// ListarGeneros obtiene todos los géneros disponibles
func (f *FachadaGeneros) ListarGeneros() ([]dto.Genero, error) {
	url := fmt.Sprintf("%s/generos", f.baseURL)

	resp, err := f.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error del servidor: código %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	var respuesta dto.RespuestaDTO[[]dto.Genero]
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if respuesta.Codigo != 200 {
		return nil, fmt.Errorf("error en el servicio: %s", respuesta.Mensaje)
	}

	return respuesta.Data, nil
}
