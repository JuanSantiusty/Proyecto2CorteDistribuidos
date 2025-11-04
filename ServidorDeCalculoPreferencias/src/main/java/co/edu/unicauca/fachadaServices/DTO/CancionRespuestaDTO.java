package co.edu.unicauca.fachadaServices.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CancionRespuestaDTO {

    @JsonProperty("id")
    private Integer id;

    @JsonProperty("titulo")
    private String titulo;
    
    @JsonProperty("artista_banda")
    private String artistaBanda;
    
    @JsonProperty("lanzamiento")
    private int lanzamiento;
    
    @JsonProperty("duracion")
    private String duracion;
    
    @JsonProperty("ruta")
    private String ruta;

     @JsonProperty("idioma")
    private String idioma;
    
    @JsonProperty("genero")
    private String genero;

    // Constructor por defecto
    public CancionRespuestaDTO() {}

    // Constructor completo
    public CancionRespuestaDTO(String titulo, String artistaBanda, int lanzamiento, 
                              String duracion, String ruta, String idioma,String genero) {
        this.titulo = titulo;
        this.artistaBanda = artistaBanda;
        this.lanzamiento = lanzamiento;
        this.duracion = duracion;
        this.ruta = ruta;
        this.idioma = idioma;
        this.genero = genero;
    }

    // Getters y Setters

      public Integer getId() {
        return id;
    }

    public String getTitulo() {
        return titulo;
    }

    public void setTitulo(String titulo) {
        this.titulo = titulo;
    }

    public String getArtistaBanda() {
        return artistaBanda;
    }

    public void setArtistaBanda(String artistaBanda) {
        this.artistaBanda = artistaBanda;
    }

    public int getLanzamiento() {
        return lanzamiento;
    }

    public void setLanzamiento(int lanzamiento) {
        this.lanzamiento = lanzamiento;
    }

     public String getIdioma() {
        return idioma;
    }

    public void setIdioma(String idioma) {
        this.idioma = idioma;
    }

    public String getDuracion() {
        return duracion;
    }

    public void setDuracion(String duracion) {
        this.duracion = duracion;
    }

    public String getRuta() {
        return ruta;
    }

    public void setRuta(String ruta) {
        this.ruta = ruta;
    }

    public String getGenero() {
        return genero;
    }

    public void setGenero(String genero) {
        this.genero = genero;
    }

    // Método estático equivalente a NuevaCancionRespuestaDTO
    public static CancionRespuestaDTO nuevaCancionRespuestaDTO(String titulo, String artistaBanda, 
                                                              int lanzamiento, String duracion, 
                                                              String ruta, String idioma,String genero) {
        return new CancionRespuestaDTO(titulo, artistaBanda, lanzamiento, duracion, ruta, idioma,genero);
    }

    @Override
    public String toString() {
        return "CancionRespuestaDTO{" +
                "titulo='" + titulo + '\'' +
                ", artistaBanda='" + artistaBanda + '\'' +
                ", lanzamiento=" + lanzamiento +
                ", duracion='" + duracion + '\'' +
                ", ruta='" + ruta + '\'' +
                ", genero='" + genero + '\'' +
                '}';
    }
}
