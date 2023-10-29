package posteos

import (
	TDAAbb "algogram/abb"
	TDALista "algogram/lista"
)

type postImplementacion struct {
	nombre       string
	lineaUsuario int
	id           int
	likes        TDAAbb.DiccionarioOrdenado[string, int]
	contenido    string
}

func comparacionString(s1 string, s2 string) int {
	if s1 > s2 {
		return 1
	}
	if s1 < s2 {
		return -1
	}
	return 0
}

func CrearPost(usuario string, orden int, id int, contenido string) Post {
	post := new(postImplementacion)
	post.nombre = usuario
	post.lineaUsuario = orden
	post.id = id
	post.contenido = contenido
	post.likes = TDAAbb.CrearABB[string, int](comparacionString)
	return post
}

func (post postImplementacion) ObtenerLineaUsuario() int {
	return post.lineaUsuario
}

func (post postImplementacion) ObtenerIDPost() int {
	return post.id
}

func (post *postImplementacion) LikearPost(user string) {
	post.likes.Guardar(user, 1)
}

func (post postImplementacion) ObtenerLikes() int {
	return post.likes.Cantidad()
}

func (post postImplementacion) ObtenerPosteador() string {
	return post.nombre
}

func (post postImplementacion) ObtenerContenidoDePost() string {
	return post.contenido
}

func (post postImplementacion) ObtenerListaDeLikes() TDALista.Lista[string] {
	likes := post.likes
	listaLikes := TDALista.CrearListaEnlazada[string]()
	iter := likes.Iterador()
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		listaLikes.InsertarUltimo(clave)
		iter.Siguiente()
	}
	return listaLikes
}
