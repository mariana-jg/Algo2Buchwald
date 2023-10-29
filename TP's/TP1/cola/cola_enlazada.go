package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	(*cola).primero = nil
	(*cola).ultimo = nil
	return cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	if cola.primero == nil {
		return true
	}
	return false
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() == true {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func crearNodo[T any](elem T) *nodoCola[T] {
	nuevoNodo := new(nodoCola[T])
	nuevoNodo.dato = elem
	return nuevoNodo
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNodo(elem)

	if cola.EstaVacia() == true {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() == true {
		panic("La cola esta vacia")
	}
	desencolado := cola.VerPrimero()
	cola.primero = cola.primero.prox
	return desencolado
}
