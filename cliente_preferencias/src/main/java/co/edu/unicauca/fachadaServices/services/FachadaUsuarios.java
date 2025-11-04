package co.edu.unicauca.fachadaServices.services;

import java.util.List;

import co.edu.unicauca.capaDeAccesoADatos.entitys.Usuario;
import co.edu.unicauca.capaDeAccesoADatos.repositorios.RepositorioUsuarios;

public class FachadaUsuarios {
    private final RepositorioUsuarios repoUsuarios;

    public FachadaUsuarios() {
        this.repoUsuarios = new RepositorioUsuarios();
    }

    public Usuario loginUsuario(String nickname) {
        return repoUsuarios.buscarPorNickname(nickname);
    }

    public List<Usuario> listarUsuarios() {
        return repoUsuarios.listarUsuarios();
    }

    public String obtenerNicknamePorId(int id) {
        Usuario u = repoUsuarios.buscarPorId(id);
        return (u != null) ? u.getNickname() : null;
    }
}
