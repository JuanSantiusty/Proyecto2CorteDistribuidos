package infraestructura

import (
	repo "cliente/repositorios"
	"fmt"
	"os"
)

type FachadaAutenticacion struct {
	repo *repo.RepositorioUsuarios
}

// NewFachadaAutenticacion crea una nueva fachada de autenticación
func NewFachadaAutenticacion(archivoUsuarios string) *FachadaAutenticacion {
	repo := repo.NewRepositorioUsuarios(archivoUsuarios)
	return &FachadaAutenticacion{
		repo: repo,
	}
}

// RespuestaAutenticacion representa la respuesta de las operaciones de autenticación
type RespuestaAutenticacion struct {
	Exito    bool   `json:"exito"`
	Mensaje  string `json:"mensaje"`
	Username string `json:"username,omitempty"`
}

// RegistrarUsuario registra un nuevo usuario
func (f *FachadaAutenticacion) RegistrarUsuario(username, password string) RespuestaAutenticacion {
	// Validaciones básicas
	if username == "" || password == "" {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "Username y password no pueden estar vacíos",
		}
	}

	if len(username) < 3 {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "El username debe tener al menos 3 caracteres",
		}
	}

	if len(password) < 4 {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "El password debe tener al menos 4 caracteres",
		}
	}

	err := f.repo.GuardarUsuario(username, password)
	if err != nil {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: fmt.Sprintf("Error al registrar usuario: %v", err),
		}
	}

	return RespuestaAutenticacion{
		Exito:    true,
		Mensaje:  "Usuario registrado exitosamente",
		Username: username,
	}
}

// IniciarSesion autentica un usuario
func (f *FachadaAutenticacion) IniciarSesion(username, password string) RespuestaAutenticacion {
	// Validaciones básicas
	if username == "" || password == "" {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "Username y password no pueden estar vacíos",
		}
	}

	esValido := f.repo.ValidarCredenciales(username, password)
	if !esValido {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "Credenciales inválidas",
		}
	}

	return RespuestaAutenticacion{
		Exito:    true,
		Mensaje:  "Inicio de sesión exitoso",
		Username: username,
	}
}

// VerificarUsuario verifica si un usuario existe
func (f *FachadaAutenticacion) VerificarUsuario(username string) bool {
	return f.repo.ExisteUsuario(username)
}

// CambiarPassword cambia la contraseña de un usuario
func (f *FachadaAutenticacion) CambiarPassword(username, oldPassword, newPassword string) RespuestaAutenticacion {
	// Primero verificar las credenciales actuales
	if !f.repo.ValidarCredenciales(username, oldPassword) {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "Credenciales actuales inválidas",
		}
	}

	// Validar nueva contraseña
	if len(newPassword) < 4 {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: "La nueva contraseña debe tener al menos 4 caracteres",
		}
	}

	// Para cambiar contraseña, necesitamos recrear el archivo sin el usuario viejo
	// y agregarlo con la nueva contraseña
	err := f.cambiarPasswordEnArchivo(username, newPassword)
	if err != nil {
		return RespuestaAutenticacion{
			Exito:   false,
			Mensaje: fmt.Sprintf("Error al cambiar contraseña: %v", err),
		}
	}

	return RespuestaAutenticacion{
		Exito:    true,
		Mensaje:  "Contraseña cambiada exitosamente",
		Username: username,
	}
}

// cambiarPasswordEnArchivo implementa el cambio de contraseña recreando el archivo
func (f *FachadaAutenticacion) cambiarPasswordEnArchivo(username, newPassword string) error {
	// Obtener todos los usuarios
	usuarios, err := f.repo.ListarUsuarios()
	if err != nil {
		return err
	}

	// Recrear el archivo
	file, err := os.Create(f.repo.Archivo)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir todos los usuarios, actualizando la contraseña del usuario especificado
	for _, usuario := range usuarios {
		if usuario.Username == username {
			usuario.Password = newPassword
		}
		linea := fmt.Sprintf("user:%s/password:%s\n", usuario.Username, usuario.Password)
		_, err := file.WriteString(linea)
		if err != nil {
			return err
		}
	}

	return nil
}
