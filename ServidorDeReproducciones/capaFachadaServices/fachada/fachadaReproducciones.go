package capafachada

import (
	capaaccesoadatos "reproducciones/capaAccesoADatos/repositorios"
	dtos "reproducciones/capaFachadaServices/DTOs"
)

type FachadaReproducciones struct {
	repo *capaaccesoadatos.RepositorioReproducciones
}

func NuevaFachadaTendencias() *FachadaReproducciones {
	return &FachadaReproducciones{
		repo: capaaccesoadatos.GetRepositorio(),
	}
}
func (f *FachadaReproducciones) RegistrarReproduccion(cancionId int, titulo string, usuarioId string) {
	f.repo.AgregarReproduccion(cancionId, titulo, usuarioId)
}

func (f *FachadaReproducciones) ObtenerReproduccionesPorCliente(usuarioId string) []dtos.ReproduccionResponseDTOOutput {
	reprosEntity := f.repo.ListarPorCliente(usuarioId)

	var reprosDTO []dtos.ReproduccionResponseDTOOutput
	for _, rep := range reprosEntity {
		reprosDTO = append(reprosDTO, dtos.ReproduccionResponseDTOOutput{
			Id:        rep.ID,
			IdUsuario: rep.UsuarioId,
			IdCancion: rep.CancionId,
			Titulo:    rep.Titulo,
			FechaHora: rep.FechaHora,
		})
	}
	return reprosDTO
}
