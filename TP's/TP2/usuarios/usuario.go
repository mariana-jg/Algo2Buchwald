package usuarios

import (
	posts "algogram/posteos"
)

type Usuario interface {

	//VerProximo devuelve un puntero hacia el próximo posteo a ver.
	VerProximo() *posts.Post

	//GuardarEnFeed guarda en el feed del usuario el posteo enviado.
	GuardarEnFeed(posts.Post)

	//ObtenerNombreUsuario devuelve el nombre del usuario.
	ObtenerNombreUsuario() string

	//ObtenerOrdenUsuario devuelve el orden del usuario.
	ObtenerOrdenUsuario() int
}
