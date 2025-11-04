package co.edu.unicauca.capaDeControladores;

import java.rmi.Remote;
import java.rmi.RemoteException;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

//Hereda de la clase renmota, lo cual la convierte en la interfaz remota
public interface ControladorPreferenciasUsuariosInt extends Remote{
    //definicion del primer metodo remoto
    public PreferenciasDTORespuesta getReferencias(String id) throws RemoteException;
    //Cada definicion del metodo debe erspecificar que puede lanzar la excepcion java.rmi.RemoteException
}
