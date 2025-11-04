package co.edu.unicauca.main;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;
import co.edu.unicauca.configuracion.lector.LectorPropiedadesConfig;
import co.edu.unicauca.configuracion.servicios.ClienteDeObjetos;
import co.edu.unicauca.fachadaServices.services.FachadaPrincipal;
import co.edu.unicauca.vista.Menu;

public class Cliente {
    public static void main(String[] args) {

        int puertoNS = Integer.parseInt(LectorPropiedadesConfig.get("ns.port"));
        String direccionIPNS=LectorPropiedadesConfig.get("ns.host");

        String identificadorObjetoRemoto = "objControladorPreferenciasUsuarios";
        
        ControladorPreferenciasUsuariosInt objRemoto = ClienteDeObjetos.obtenerObjetoRemoto(direccionIPNS, puertoNS, identificadorObjetoRemoto);
        FachadaPrincipal objFachada = new FachadaPrincipal(objRemoto);
        Menu objMenu = new Menu(objFachada);
        objMenu.ejecutarMenuPrincipal();
        
    }
}


