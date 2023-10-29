package grafo

import (
	TDALista "vamosmoshi/lista"
)

type Grafo[K comparable] interface {
	AgregarVertice(vertice K) //listo

	BorrarVertice(vertice K) //listo

	AgregarArista(desde K, hasta K, peso int) //listo

	BorrarArista(desde K, hasta K, peso int) //listo

	EstanUnidos(desde K, hasta K) bool //listo

	PesoUnion(desde K, hasta K) int //listo

	ExisteID(id K) bool //listo

	ObtenerVerticeRandom() K //listo

	ObtenerVertices() TDALista.Lista[K] //listo

	ObtenerAdyacentes(vertice K) TDALista.Lista[K] //
	IterarVertices() Iterador[K]
}

type Iterador[K comparable] interface {
	HaySiguiente() bool

	Siguiente()
	VerActual() K
}
