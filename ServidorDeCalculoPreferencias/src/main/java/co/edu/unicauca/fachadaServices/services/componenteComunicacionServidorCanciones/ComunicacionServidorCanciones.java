// ServidorDeCalculoPreferencias/src/main/java/co/edu/unicauca/fachadaServices/services/componenteComunicacionServidorCanciones/ComunicacionServidorCanciones.java

package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorCanciones;

import co.edu.unicauca.fachadaServices.DTO.RespuestaDTO;
import co.edu.unicauca.fachadaServices.DTO.CancionRespuestaDTO;
import feign.Feign;
import feign.jackson.JacksonDecoder;
import feign.jackson.JacksonEncoder;

import java.util.ArrayList;
import java.util.List;

public class ComunicacionServidorCanciones {

    // ✅ VERIFICAR: URL del ServidorDeCanciones (puerto 5000)
    private static final String BASE_URL = "http://localhost:5000/";
    private final CancionesRemoteClient client;

    public ComunicacionServidorCanciones(){
        this.client = Feign.builder()
                    .encoder(new JacksonEncoder())
                    .decoder(new JacksonDecoder())
                    .target(CancionesRemoteClient.class, BASE_URL);
    }

    public List<CancionRespuestaDTO> obtenerCancionesRemotas(){
        try{
            System.out.println("📡 Llamando a ServidorDeCanciones: GET /canciones");
            
            // Ahora obtenemos RespuestaDTO que contiene la lista de canciones
            RespuestaDTO<List<CancionRespuestaDTO>> respuesta = client.obtenerCanciones();
            
            // Verificamos el código de respuesta
            if (respuesta.getCodigo() == 200) {
                List<CancionRespuestaDTO> canciones = respuesta.getData();
                System.out.println("✅ Canciones obtenidas: " + (canciones != null ? canciones.size() : 0));
                return canciones != null ? canciones : new ArrayList<>();
            } else {
                System.err.println("❌ Error en la respuesta del servidor: Código " + 
                                 respuesta.getCodigo() + " - " + respuesta.getMensaje());
                return new ArrayList<>();
            }
            
        } catch (Exception e){
            System.err.println("❌ Error obteniendo canciones: " + e.getMessage());
            e.printStackTrace();
            return new ArrayList<>();
        }
    }
}