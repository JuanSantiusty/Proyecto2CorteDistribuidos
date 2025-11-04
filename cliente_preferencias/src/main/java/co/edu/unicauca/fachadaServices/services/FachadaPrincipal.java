package co.edu.unicauca.fachadaServices.services;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;

public class FachadaPrincipal {

    public final FachadaUsuarios usuarios;
    public final FachadaGestorUsuariosIml preferencias;


    public FachadaPrincipal(ControladorPreferenciasUsuariosInt objRemotoPreferencias) {

        this.usuarios = new FachadaUsuarios();
        this.preferencias = new FachadaGestorUsuariosIml(objRemotoPreferencias);
    }

}