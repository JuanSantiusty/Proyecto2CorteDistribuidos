package capaFachadaServices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// FachadaBuscarCancion proporciona métodos específicos para buscar canciones
type FachadaBuscarCancion struct {
	baseURL string
	client  *http.Client
}

// NewFachadaBuscarCancion crea una nueva instancia de la fachada
func NewFachadaBuscarCancion(baseURL string) *FachadaBuscarCancion {
	return &FachadaBuscarCancion{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// CancionRespuesta representa la estructura de respuesta de una canción
type CancionRespuesta struct {
	ID           int32  `json:"id"`
	Titulo       string `json:"titulo"`
	ArtistaBanda string `json:"artista_banda"`
	Lanzamiento  int32  `json:"lanzamiento"`
	Duracion     string `json:"duracion"`
	Ruta         string `json:"ruta"`
	Idioma       string `json:"idioma"`
	Genero       string `json:"genero"`
}

// RespuestaServicio representa la respuesta genérica del servicio
type RespuestaServicio struct {
	Data    *CancionRespuesta `json:"data"`
	Codigo  int32             `json:"codigo"`
	Mensaje string            `json:"mensaje"`
}

// BuscarCancionPorTitulo busca una canción específica por título
func (f *FachadaBuscarCancion) BuscarCancionPorTitulo(titulo string) (*CancionRespuesta, error) {
	// Validar que el título no esté vacío
	if titulo == "" {
		return nil, fmt.Errorf("el título no puede estar vacío")
	}

	// Codificar el parámetro de título
	params := url.Values{}
	params.Add("titulo", titulo)

	url := fmt.Sprintf("%s/canciones/buscar?%s", f.baseURL, params.Encode())

	// Realizar la petición GET
	resp, err := f.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la petición: %v", err)
	}
	defer resp.Body.Close()

	// Verificar el código de estado HTTP
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error del servidor: código %d - %s", resp.StatusCode, string(body))
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	// Decodificar la respuesta JSON
	var respuesta RespuestaServicio
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	// Verificar el código de respuesta del servicio
	if respuesta.Codigo != 200 {
		return nil, fmt.Errorf("error en el servicio: %s", respuesta.Mensaje)
	}

	return respuesta.Data, nil
}

// BuscarCancionPorTituloSimplificado versión simplificada que retorna la canción o nil
func (f *FachadaBuscarCancion) BuscarCancionPorTituloSimplificado(titulo string) *CancionRespuesta {
	cancion, err := f.BuscarCancionPorTitulo(titulo)
	if err != nil {
		fmt.Printf("⚠️ No se pudo encontrar la canción '%s': %v\n", titulo, err)
		return nil
	}
	return cancion
}
