package models

import "time"

type SolicitudPin struct {
	Correo string `json:"correo"`
	// "entrar" o vacio si es registro
	Modo string `json:"modo"`
}

type SolicitudVerificar struct {
	Correo string `json:"correo"`
	Pin    string `json:"pin"`
}

// el PIN se guarda en memoria mientras el MID este prendido
type CodigoPin struct {
	Correo   string
	Pin      string
	Creado   time.Time
	Vence    time.Time
	Intentos int
	Activo   bool
}

type SolicitudRegistro struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	FechaNacimiento string `json:"fechaNacimiento"`
	Email           string `json:"email"`
	Telefono        string `json:"telefono"`
	Ubicacion       string `json:"ubicacion"`
	Direccion       string `json:"direccion"`
	Barrio          string `json:"barrio"`
	Avatar          string `json:"avatar"`
}

type LoginAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
