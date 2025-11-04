package infraestructura

import (
	dto "cliente/infraestructura/DTO"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// FachadaCanciones proporciona métodos para consumir los servicios de canciones
type FachadaCanciones struct {
	baseURL string
	client  *http.Client
}

// NewFachadaCanciones crea una nueva instancia de la fachada de canciones
func NewFachadaCanciones(baseURL string) *FachadaCanciones {
	return &FachadaCanciones{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		client:  &http.Client{},
	}
}

// BuscarCancionesPorGenero busca canciones por género
func (f *FachadaCanciones) BuscarCancionesPorGenero(genero string) ([]dto.CancionRespuestaDTO, error) {
	// Codificar el parámetro de género
	params := url.Values{}
	params.Add("genero", genero)

	url := fmt.Sprintf("%s/canciones/buscar?%s", f.baseURL, params.Encode())

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

	var respuesta dto.RespuestaDTO[[]dto.CancionRespuestaDTO]
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if respuesta.Codigo != 200 {
		return nil, fmt.Errorf("error en el servicio: %s", respuesta.Mensaje)
	}

	return respuesta.Data, nil
}

// BuscarCancionPorTitulo busca una canción específica por título
func (f *FachadaCanciones) BuscarCancionPorTitulo(titulo string) (*dto.RespuestaDTO[dto.CancionRespuestaDTO], error) {
	// Codificar el parámetro de título
	params := url.Values{}
	params.Add("titulo", titulo)

	url := fmt.Sprintf("%s/canciones/buscar?%s", f.baseURL, params.Encode())

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

	var respuesta dto.RespuestaDTO[dto.CancionRespuestaDTO]
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if respuesta.Codigo != 200 {
		return nil, fmt.Errorf("error en el servicio: %s", respuesta.Mensaje)
	}

	return &respuesta, nil
}

// ListarTodasLasCanciones obtiene todas las canciones disponibles
func (f *FachadaCanciones) ListarTodasLasCanciones() ([]dto.CancionRespuestaDTO, error) {
	url := fmt.Sprintf("%s/canciones", f.baseURL)

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

	var respuesta dto.RespuestaDTO[[]dto.CancionRespuestaDTO]
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if respuesta.Codigo != 200 {
		return nil, fmt.Errorf("error en el servicio: %s", respuesta.Mensaje)
	}

	return respuesta.Data, nil
}
