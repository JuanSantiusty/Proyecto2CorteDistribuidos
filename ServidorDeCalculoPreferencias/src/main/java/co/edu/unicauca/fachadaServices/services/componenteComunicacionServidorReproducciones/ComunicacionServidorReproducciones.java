// ServidorDeCalculoPreferencias/src/main/java/co/edu/unicauca/fachadaServices/services/componenteComunicacionServidorReproducciones/ComunicacionServidorReproducciones.java

package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import feign.Feign;
import feign.jackson.JacksonDecoder;

import java.util.ArrayList;
import java.util.List;

public class ComunicacionServidorReproducciones {

   // ✅ VERIFICAR: URL del ServidorDeReproducciones (puerto 5004)
   private static final String BASE_URL = "http://localhost:5004/";
    private final ReproduccionesRemoteClient client;

    public ComunicacionServidorReproducciones(){
        this.client = Feign.builder()
                    .decoder(new JacksonDecoder())
                    .target(ReproduccionesRemoteClient.class, BASE_URL);
    }

    public List<ReproduccionesDTOEntrada> obtenerReproduccionesRemotas(String idUsuario){
        try{
            System.out.println("📡 Llamando a ServidorDeReproducciones: GET /reproducciones/listar?idUsuario=" + idUsuario);
            List<ReproduccionesDTOEntrada> reproducciones = client.obtenerReproducciones(idUsuario);
            System.out.println("✅ Reproducciones obtenidas: " + (reproducciones != null ? reproducciones.size() : 0));
            return reproducciones != null ? reproducciones : new ArrayList<>();
        } catch (Exception e){
            System.err.println("❌ Error obteniendo reproducciones: " + e.getMessage());
            return new ArrayList<>();
        }
    }
}