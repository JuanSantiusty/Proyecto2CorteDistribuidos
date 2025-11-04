package co.edu.unicauca.fachadaServices.DTO;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CancionDTOEntrada {
    private Integer ID;
    private String Titulo;
    private String Artista;
    private String Genero;
    private String Idioma;
}

