package grafo

import (
	TDAHash "vamosmoshi/hash"
	TDALista "vamosmoshi/lista"
)

type grafo[K comparable] struct {
	vertices   TDAHash.Diccionario[K, TDAHash.Diccionario[K, int]]
	esDirigido bool
}

type iterador[K comparable] struct {
	iterador TDAHash.IterDiccionario[K, TDAHash.Diccionario[K, int]]
}

func CrearGrafo[K comparable](dirigido bool) Grafo[K] {
	grafo := new(grafo[K])
	grafo.esDirigido = dirigido
	grafo.vertices = TDAHash.CrearHash[K, TDAHash.Diccionario[K, int]]()
	return grafo
}

func (grafo *grafo[K]) AgregarVertice(vertice K) {
	if grafo.vertices.Pertenece(vertice) {
		panic("El elemento ya pertenece.")
	}
	adyacentes := TDAHash.CrearHash[K, int]()
	grafo.vertices.Guardar(vertice, adyacentes)
}

func (grafo *grafo[K]) BorrarVertice(vertice K) {
	if !grafo.vertices.Pertenece(vertice) {
		panic("No existe el vertice.")
	}
	grafo.vertices.Borrar(vertice)
}

func (grafo *grafo[K]) AgregarArista(desde K, hasta K, peso int) {
	if !grafo.vertices.Pertenece(desde) || !grafo.vertices.Pertenece(hasta) {
		panic("No se puede agregar una arista si ambos elementos no pertenecen")
	}
	if !grafo.esDirigido {
		adyacentes := grafo.vertices.Obtener(hasta)
		adyacentes.Guardar(desde, peso)
		grafo.vertices.Guardar(hasta, adyacentes)
	}

	adyacentes := grafo.vertices.Obtener(desde)
	adyacentes.Guardar(hasta, peso)
	grafo.vertices.Guardar(desde, adyacentes)
}

func (grafo *grafo[K]) BorrarArista(desde K, hasta K, peso int) {
	if !grafo.vertices.Pertenece(desde) || !grafo.vertices.Pertenece(hasta) {
		panic("No se puede borrar una arista si algún elemento no pertenece")
	}
	if !grafo.esDirigido {
		adyacentes := grafo.vertices.Obtener(hasta)
		adyacentes.Borrar(desde)
	}
	adyacentes := grafo.vertices.Obtener(desde)
	adyacentes.Borrar(hasta)
}

func (grafo grafo[K]) EstanUnidos(desde K, hasta K) bool {
	if !grafo.vertices.Pertenece(desde) || !grafo.vertices.Pertenece(hasta) {
		panic("No existe el vertice")
	}
	adyacentesDesde := grafo.vertices.Obtener(desde)
	return adyacentesDesde.Pertenece(hasta)
}

func (grafo grafo[K]) PesoUnion(desde K, hasta K) int {
	if !grafo.vertices.Pertenece(desde) || !grafo.vertices.Pertenece(hasta) {
		panic("No existe el vertice.")
	}
	adyacentes := grafo.vertices.Obtener(desde)
	if !adyacentes.Pertenece(hasta) {
		panic("No existe la arista.")
	}
	return adyacentes.Obtener(hasta)
}

func (grafo grafo[K]) ExisteID(id K) bool {
	return grafo.vertices.Pertenece(id)
}

func (grafo grafo[K]) ObtenerVerticeRandom() K {
	iter := grafo.vertices.Iterador()
	vertice, _ := iter.VerActual()
	return vertice
}

func (grafo grafo[K]) ObtenerVertices() TDALista.Lista[K] {
	listaVertices := TDALista.CrearListaEnlazada[K]()
	iter := grafo.vertices.Iterador()
	for iter.HaySiguiente() {
		vertice, _ := iter.VerActual()
		listaVertices.InsertarUltimo(vertice)
		iter.Siguiente()
	}
	return listaVertices
}

func (grafo grafo[K]) ObtenerAdyacentes(vertice K) TDALista.Lista[K] {
	listaAdyacentes := TDALista.CrearListaEnlazada[K]()
	adyacentes := grafo.vertices.Obtener(vertice)
	iter := adyacentes.Iterador()
	for iter.HaySiguiente() {
		adyacente, _ := iter.VerActual()
		listaAdyacentes.InsertarUltimo(adyacente)
		iter.Siguiente()
	}
	return listaAdyacentes
}

func (grafo grafo[K]) IterarVertices() Iterador[K] {
	nuevoIterador := new(iterador[K])
	nuevoIterador.iterador = grafo.vertices.Iterador()
	return nuevoIterador
}

func (iter iterador[K]) HaySiguiente() bool {
	return iter.iterador.HaySiguiente()
}

func (iter *iterador[K]) Siguiente() {
	iter.iterador.Siguiente()
}

func (iter *iterador[K]) VerActual() K {
	vertice, _ := iter.iterador.VerActual()
	return vertice
}
