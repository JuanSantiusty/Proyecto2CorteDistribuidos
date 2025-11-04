package co.edu.unicauca.capaDeAccesoADatos.repositorios;

import java.util.ArrayList;
import java.util.List;

import co.edu.unicauca.capaDeAccesoADatos.entitys.Usuario;

public class RepositorioUsuarios {
    private List<Usuario> usuarios;

    public RepositorioUsuarios() {
        usuarios = new ArrayList<>();
        usuarios.add(new Usuario(1, "juan"));
        usuarios.add(new Usuario(2, "maria"));
        usuarios.add(new Usuario(3, "antonio"));
        usuarios.add(new Usuario(4, "miguel"));
        usuarios.add(new Usuario(5, "user1"));
        usuarios.add(new Usuario(6, "user2"));
        usuarios.add(new Usuario(7, "user3"));
    }

    public Usuario buscarPorNickname(String nickname) {
        for (Usuario u : usuarios) {
            if (u.getNickname().equalsIgnoreCase(nickname)) {
                return u;
            }
        }
        return null;
    }

    public List<Usuario> listarUsuarios() {
        return usuarios;
    }

    public Usuario buscarPorId(int id) {
        for (Usuario u : usuarios) {
            if (u.getId() == id) {
                return u;
            }
        }
        return null;
    }
}
