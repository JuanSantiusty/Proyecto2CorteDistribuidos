
package co.edu.unicauca.fachadaServices.services;

import java.util.List;

import co.edu.unicauca.fachadaServices.DTO.*;
import co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias.CalculadorPreferencias;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorCanciones.ComunicacionServidorCanciones;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones.ComunicacionServidorReproducciones;

public class PreferenciasServiceImpl implements IPreferenciasService {
	private ComunicacionServidorCanciones comunicacionServidorCanciones;
	private ComunicacionServidorReproducciones comunicacionServidorReproducciones;
	private CalculadorPreferencias calculadorPreferencias;
	

	public PreferenciasServiceImpl() {
		this.comunicacionServidorCanciones = new ComunicacionServidorCanciones();
		this.comunicacionServidorReproducciones = new ComunicacionServidorReproducciones();
		this.calculadorPreferencias = new CalculadorPreferencias();
	}

	@Override
	public PreferenciasDTORespuesta getReferencias(String id) {
		System.out.println("Obteniendo preferencias para el usuario con ID: " + id);
		List<CancionRespuestaDTO> objCanciones = this.comunicacionServidorCanciones.obtenerCancionesRemotas();
		System.out.println("Canciones obtenidas del servidor de canciones: "+objCanciones.size());
		for(CancionRespuestaDTO cancion : objCanciones){
			System.out.println("Cancion Obtenida: " + cancion.getTitulo());
			System.out.println("Genero: " + cancion.getGenero());
			System.out.println("Artista: " + cancion.getArtistaBanda());
		}
		List<ReproduccionesDTOEntrada> reproduccionesUsuario = this.comunicacionServidorReproducciones.obtenerReproduccionesRemotas(id);
		System.out.println("Reproducciones obtenidas del servidor de reproducciones para el usuario "+ id);
		for (ReproduccionesDTOEntrada reproduccion : reproduccionesUsuario) {
			System.out.println(reproduccion.getIdUsuario() + " " + reproduccion.getIdCancion());
		}

		return this.calculadorPreferencias.calcular(id, objCanciones, reproduccionesUsuario);
	}

	
}
