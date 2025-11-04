package co.edu.unicauca.fachadaServices.DTO;

public class RespuestaDTO<T> {
    private T data;
    private int codigo;
    private String mensaje;

    // Constructor por defecto
    public RespuestaDTO() {}

    // Constructor parametrizado (equivalente a NewRespuestaDTO)
    public RespuestaDTO(T data, int codigo, String mensaje) {
        this.data = data;
        this.codigo = codigo;
        this.mensaje = mensaje;
    }

    // Getters y Setters
    public T getData() {
        return data;
    }

    public void setData(T data) {
        this.data = data;
    }

    public int getCodigo() {
        return codigo;
    }

    public void setCodigo(int codigo) {
        this.codigo = codigo;
    }

    public String getMensaje() {
        return mensaje;
    }

    public void setMensaje(String mensaje) {
        this.mensaje = mensaje;
    }

    // Método estático equivalente a NewRespuestaDTO
    public static <T> RespuestaDTO<T> newRespuestaDTO(T data, int codigo, String mensaje) {
        return new RespuestaDTO<>(data, codigo, mensaje);
    }

    @Override
    public String toString() {
        return "RespuestaDTO{" +
                "data=" + data +
                ", codigo=" + codigo +
                ", mensaje='" + mensaje + '\'' +
                '}';
    }
}
