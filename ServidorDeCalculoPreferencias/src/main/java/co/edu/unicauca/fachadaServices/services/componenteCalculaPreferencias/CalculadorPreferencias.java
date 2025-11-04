package co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias;

import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

import co.edu.unicauca.fachadaServices.DTO.CancionRespuestaDTO;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaIdiomaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;

public class CalculadorPreferencias {

    
    public PreferenciasDTORespuesta calcular(String idUsuario,
                                              List<CancionRespuestaDTO> canciones,
                                              List<ReproduccionesDTOEntrada> reproducciones) {
        Map<Integer, CancionRespuestaDTO> mapaCanciones = canciones.stream()
            .filter(Objects::nonNull)
            .filter(c -> c.getId() != null)
            .collect(Collectors.toMap(CancionRespuestaDTO::getId, c -> c, (a,b) -> a));
       
        
        Map<String, Integer> contadorGeneros = new HashMap<>();
        Map<String, Integer> contadorArtistas = new HashMap<>();
        Map<String, Integer> contadorIdiomas = new HashMap<>();

        for (ReproduccionesDTOEntrada r : reproducciones) {
            Integer idCancion = r.getIdCancion();
            if (idCancion == null) continue;

            CancionRespuestaDTO c = mapaCanciones.get(idCancion);
            if (c == null) {               
                continue;
            }

            String genero = c.getGenero() == null ? "Desconocido" : c.getGenero();
            String artista = c.getArtistaBanda() == null ? "Desconocido" : c.getArtistaBanda();
            String idioma = c.getIdioma() == null ? "Desconocido" : c.getIdioma();

            contadorGeneros.put(genero, contadorGeneros.getOrDefault(genero, 0) + 1);
            contadorArtistas.put(artista, contadorArtistas.getOrDefault(artista, 0) + 1);
            contadorIdiomas.put(idioma, contadorIdiomas.getOrDefault(artista, 0) + 1);
        }

        
        List<PreferenciaGeneroDTORespuesta> preferenciasGeneros = contadorGeneros.entrySet().stream()
            .map(e -> {
                PreferenciaGeneroDTORespuesta dto = new PreferenciaGeneroDTORespuesta();
                dto.setNombreGenero(e.getKey());
                dto.setNumeroPreferencias(e.getValue());
                return dto;
            })
            .sorted(Comparator.comparingInt(PreferenciaGeneroDTORespuesta::getNumeroPreferencias).reversed()
                    .thenComparing(PreferenciaGeneroDTORespuesta::getNombreGenero))
            .collect(Collectors.toList());

        List<PreferenciaArtistaDTORespuesta> preferenciasArtistas = contadorArtistas.entrySet().stream()
            .map(e -> {
                PreferenciaArtistaDTORespuesta dto = new PreferenciaArtistaDTORespuesta();
                dto.setNombreArtista(e.getKey());
                dto.setNumeroPreferencias(e.getValue());
                return dto;
            })
            .sorted(Comparator.comparingInt(PreferenciaArtistaDTORespuesta::getNumeroPreferencias).reversed()
                    .thenComparing(PreferenciaArtistaDTORespuesta::getNombreArtista))
            .collect(Collectors.toList());

        List<PreferenciaIdiomaDTORespuesta> preferenciasIdiomas = contadorIdiomas.entrySet().stream()
            .map(e -> {
                PreferenciaIdiomaDTORespuesta dto = new PreferenciaIdiomaDTORespuesta();
                dto.setNombreIdioma(e.getKey());
                dto.setNumeroPreferencias(e.getValue());
                return dto;
            })
            .sorted(Comparator.comparingInt(PreferenciaIdiomaDTORespuesta::getNumeroPreferencias).reversed()
                    .thenComparing(PreferenciaIdiomaDTORespuesta::getNombreIdioma))
            .collect(Collectors.toList());
        PreferenciasDTORespuesta respuesta = new PreferenciasDTORespuesta();
        respuesta.setIdUsuario(idUsuario);
        respuesta.setPreferenciasGeneros(preferenciasGeneros);
        respuesta.setPreferenciasArtistas(preferenciasArtistas);
        respuesta.setPreferenciasIdiomas(preferenciasIdiomas);

        return respuesta; 
    }
}
