package usuarios

import (
	TDAHeap "algogram/heap"
	posts "algogram/posteos"
)

type posteosHeap struct {
	post               *posts.Post
	lineaUsuarioActual int
}

type usuarioImplementacion struct {
	nombre string
	orden  int
	feed   TDAHeap.ColaPrioridad[posteosHeap]
}

func distancia(actual int, insertado int) int {
	distancia := actual - insertado
	if distancia < 0 {
		return distancia * (-1)
	}
	return distancia
}

func contratoAfinidad[T comparable](usuario1 posteosHeap, usuario2 posteosHeap) int {
	post1 := *(usuario1.post)
	post2 := *(usuario2.post)

	distancia1 := distancia(usuario1.lineaUsuarioActual, post1.ObtenerLineaUsuario())
	distancia2 := distancia(usuario2.lineaUsuarioActual, post2.ObtenerLineaUsuario())

	if distancia1 == distancia2 {
		if post1.ObtenerIDPost() < post2.ObtenerIDPost() {
			return 1
		}
		return -1
	}
	if distancia1 < distancia2 {
		return 1
	}
	return -1
}

func CrearUsuario[T comparable](nombre string, orden int) Usuario {
	usuario := new(usuarioImplementacion)
	usuario.nombre = nombre
	usuario.orden = orden
	usuario.feed = TDAHeap.CrearHeap(contratoAfinidad[T])
	return usuario
}

func (usuario usuarioImplementacion) ObtenerNombreUsuario() string {
	return usuario.nombre
}

func (usuario usuarioImplementacion) ObtenerOrdenUsuario() int {
	return usuario.orden
}

func (usuario *usuarioImplementacion) GuardarEnFeed(posteo posts.Post) {
	aux := new(posteosHeap)
	aux.post = &posteo
	aux.lineaUsuarioActual = usuario.orden
	usuario.feed.Encolar(*aux)
}

func (usuario *usuarioImplementacion) VerProximo() *posts.Post {
	if usuario.feed.EstaVacia() {
		return nil
	}
	proximoPosteo := usuario.feed.Desencolar()
	return proximoPosteo.post
}
