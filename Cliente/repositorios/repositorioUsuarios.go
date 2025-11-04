package repositorio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type RepositorioUsuarios struct {
	Archivo string
	mu      sync.RWMutex
}

type Usuario struct {
	Username string
	Password string
}

// NewRepositorioUsuarios crea una nueva instancia del repositorio
func NewRepositorioUsuarios(archivo string) *RepositorioUsuarios {
	return &RepositorioUsuarios{
		Archivo: archivo,
	}
}

// GuardarUsuario guarda un nuevo usuario en el archivo
func (r *RepositorioUsuarios) GuardarUsuario(username, password string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verificar si el usuario ya existe
	if r.existeUsuario(username) {
		return fmt.Errorf("el usuario '%s' ya existe", username)
	}

	// Abrir archivo en modo append
	file, err := os.OpenFile(r.Archivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Escribir en el formato: "user:username/password:password"
	linea := fmt.Sprintf("user:%s/password:%s\n", username, password)

	_, err = file.WriteString(linea)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	return nil
}

// ValidarCredenciales verifica si el usuario y contraseña son correctos
func (r *RepositorioUsuarios) ValidarCredenciales(username, password string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	usuarios, err := r.cargarUsuarios()
	if err != nil {
		return false
	}

	for _, usuario := range usuarios {
		if usuario.Username == username && usuario.Password == password {
			return true
		}
	}

	return false
}

// ExisteUsuario verifica si un usuario existe
func (r *RepositorioUsuarios) ExisteUsuario(username string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.existeUsuario(username)
}

// existeUsuario verifica internamente si un usuario existe
func (r *RepositorioUsuarios) existeUsuario(username string) bool {
	usuarios, err := r.cargarUsuarios()
	if err != nil {
		return false
	}

	for _, usuario := range usuarios {
		if usuario.Username == username {
			return true
		}
	}

	return false
}

// cargarUsuarios lee todos los usuarios del archivo
func (r *RepositorioUsuarios) cargarUsuarios() ([]Usuario, error) {
	var usuarios []Usuario

	file, err := os.Open(r.Archivo)
	if err != nil {
		if os.IsNotExist(err) {
			return usuarios, nil // Archivo no existe, retornar lista vacía
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		usuario, err := r.parsearLinea(linea)
		if err == nil {
			usuarios = append(usuarios, usuario)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

// parsearLinea parsea una línea del formato "user:username/password:password"
func (r *RepositorioUsuarios) parsearLinea(linea string) (Usuario, error) {
	// Dividir por "/" para separar user y password
	partes := strings.Split(linea, "/")
	if len(partes) != 2 {
		return Usuario{}, fmt.Errorf("formato de línea inválido")
	}

	// Extraer username
	userPart := strings.TrimPrefix(partes[0], "user:")
	if userPart == partes[0] {
		return Usuario{}, fmt.Errorf("formato de usuario inválido")
	}

	// Extraer password
	passPart := strings.TrimPrefix(partes[1], "password:")
	if passPart == partes[1] {
		return Usuario{}, fmt.Errorf("formato de password inválido")
	}

	return Usuario{
		Username: userPart,
		Password: passPart,
	}, nil
}

// ListarUsuarios retorna todos los usuarios (útil para debugging)
func (r *RepositorioUsuarios) ListarUsuarios() ([]Usuario, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.cargarUsuarios()
}
