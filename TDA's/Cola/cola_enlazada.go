package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	(*cola).primero = nil
	(*cola).ultimo = nil
	return cola
}

func crearNodo[T any](elem T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = elem
	nodo.prox = nil
	return nodo
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil

}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return (cola.primero).dato

}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	newNodo := crearNodo(elem)
	if cola.EstaVacia() {
		cola.primero = newNodo
	} else {
		(cola.ultimo).prox = newNodo
	}
	cola.ultimo = newNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	nodoDesencolado := cola.primero
	datoDesencolado := nodoDesencolado.dato
	cola.primero = cola.primero.prox

	if cola.ultimo == nodoDesencolado {
		cola.ultimo = cola.primero
	}

	return datoDesencolado
}
