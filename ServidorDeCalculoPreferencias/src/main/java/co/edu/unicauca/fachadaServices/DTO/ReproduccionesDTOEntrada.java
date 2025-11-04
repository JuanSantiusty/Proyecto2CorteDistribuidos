package co.edu.unicauca.fachadaServices.DTO;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ReproduccionesDTOEntrada {
   private Integer id;
   private String idUsuario;
   private Integer idCancion;
   private String titulo;
   private String fechaHora;

}
