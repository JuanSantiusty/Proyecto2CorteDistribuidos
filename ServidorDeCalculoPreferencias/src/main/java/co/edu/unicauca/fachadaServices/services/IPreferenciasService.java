
package co.edu.unicauca.fachadaServices.services;

import java.rmi.RemoteException;

import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;


public interface IPreferenciasService {
    public PreferenciasDTORespuesta getReferencias(String id) throws RemoteException;
	
}


