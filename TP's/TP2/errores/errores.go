package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type UsuarioLoggeado struct{}

func (e UsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type UsuarioInexistente struct{}

func (e UsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type SinUsuarioEnLinea struct{}

func (e SinUsuarioEnLinea) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorSinUsuarioSinPosteo struct{}

func (e ErrorSinUsuarioSinPosteo) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorSinUsuarioPostInexistente struct{}

func (e ErrorSinUsuarioPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostInexistenteSinLikes struct{}

func (e ErrorPostInexistenteSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
