package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorCanciones;

import co.edu.unicauca.fachadaServices.DTO.*;
import feign.Headers;
import feign.RequestLine;

import java.util.List;


public interface CancionesRemoteClient {
    @RequestLine("GET /canciones")
    @Headers("Accept: application/json")
    RespuestaDTO<List<CancionRespuestaDTO>> obtenerCanciones();
}
