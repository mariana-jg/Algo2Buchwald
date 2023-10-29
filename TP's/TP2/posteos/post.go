package posteos

import (
	TDALista "algogram/lista"
)

type Post interface {

	//ObtenerLineaUsario devuelve la linea del usuario que realiza el posteo.
	ObtenerLineaUsuario() int

	//ObtenerLineaUsario devuelve el ID del posteo.
	ObtenerIDPost() int

	//ObtenerPosteador devuelve el nombre del usuario que realiza el posteo.
	ObtenerPosteador() string

	//ObtenerContenidoDePost devuelve el contenido del posteo.
	ObtenerContenidoDePost() string

	//LikearPost guarda la informacion del usuario que le dio like.
	LikearPost(user string)

	//ObtenerLikes devuelve la cantidad de likes del posteo.
	ObtenerLikes() int

	//ObtenerListaDeLikes devuelve una lista que contiene a todos los usuarios que dieron le dieron like al posteo.
	ObtenerListaDeLikes() TDALista.Lista[string]
}
