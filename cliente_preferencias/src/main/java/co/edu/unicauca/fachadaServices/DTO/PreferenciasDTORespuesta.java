package co.edu.unicauca.fachadaServices.DTO;

import java.io.Serializable;
import java.util.List;

import lombok.Data;

@Data
public class PreferenciasDTORespuesta implements Serializable{
   private String idUsuario;
   private List<PreferenciaGeneroDTORespuesta> preferenciasGeneros;
   private List<PreferenciaArtistaDTORespuesta> preferenciasArtistas;
   private List<PreferenciaIdiomaDTORespuesta> preferenciasIdiomas;
}
