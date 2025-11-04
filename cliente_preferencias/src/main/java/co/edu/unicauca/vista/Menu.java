package  co.edu.unicauca.vista;

import co.edu.unicauca.capaDeAccesoADatos.entitys.Usuario;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.services.FachadaPrincipal;
import co.edu.unicauca.utilidades.UtilidadesConsola;

import java.rmi.RemoteException;
import java.util.List;

public class Menu {
    
    private final FachadaPrincipal objFachada;

    
    public Menu(FachadaPrincipal objFachada)
    {
        this.objFachada=objFachada;
    }

    public void ejecutarMenuPrincipal()
    {
        int opcion = 0;
        do
        {
                System.out.println("==Menu==");
                System.out.println("1. Consultar preferencias del Usuario");
                System.out.println("2. Salir");

                opcion = UtilidadesConsola.leerEntero();

            switch(opcion)
            {
                case 1: opcion1(); break;               
                case 2: System.out.println("Gracias por usar el sistema"); break;
                default: System.out.println("Opcion no valida");
            }

        }while(opcion != 2);
    }

    private void opcion1()
    {
        List<Usuario> usuarios=this.objFachada.usuarios.listarUsuarios();
        System.out.println("\nUsuarios disponibles");
            usuarios.forEach(usuario ->{
                System.out.println(usuario.getId()+"-" + usuario.getNickname());
            });
        System.out.println("Ingrese el id del usuario a consultar sus preferencias:");
        String idUsuario = UtilidadesConsola.leerCadena();
        try {
            //invocacion del metodo remoto a travez de la fachada
            PreferenciasDTORespuesta respuesta = this.objFachada.preferencias.getReferencias(idUsuario);
            System.out.println("==Preferencias del usuario==");
            System.out.println("Usuario: "+idUsuario);
            
            System.out.println("\n Generos");
            respuesta.getPreferenciasGeneros().forEach(genero -> {
                System.out.println(genero.getNombreGenero()+ " Cantidad de veces escuchado: "+genero.getNumeroPreferencias());

            });
            System.out.println("\nArtistas");
            respuesta.getPreferenciasArtistas().forEach(artista ->{
                System.out.println(artista.getNombreArtista()+" Cantidad de veces escuchado: " + artista.getNumeroPreferencias());
            });

            System.out.println("\nIdiomas");
            respuesta.getPreferenciasIdiomas().forEach(idioma ->{
                System.out.println(idioma.getNombreIdioma()+" Cantidad de veces escuchado: " + idioma.getNumeroPreferencias());
            });
            
        } catch (RemoteException e) {
            System.out.println("Error al consultar las preferencias: " + e.getMessage());
        }
    }


}
