package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func crearNodo[T any](elem T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = elem
	return nuevoNodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nodo := crearNodo(elem)

	if lista.EstaVacia() {
		lista.primero = nodo
		lista.ultimo = nodo
		lista.largo++
	} else {
		primerNodoActual := lista.primero
		lista.primero = nodo
		lista.primero.siguiente = primerNodoActual
		lista.largo++
	}

}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	if lista.EstaVacia() {
		lista.InsertarPrimero(elem)
	} else {
		nodo := crearNodo(elem)
		lista.ultimo.siguiente = nodo
		lista.ultimo = nodo
		lista.largo++
	}
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	borrado := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--

	if lista.EstaVacia() {
		lista.ultimo = nil
	}
	return borrado
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	nuevoIterador := new(iterListaEnlazada[T])
	nuevoIterador.lista = lista
	nuevoIterador.actual = lista.primero
	nuevoIterador.anterior = nil
	return nuevoIterador
}

func (iter iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter iterListaEnlazada[T]) HaySiguiente() bool {
	if iter.lista.EstaVacia() {
		return false
	}
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
	return iter.anterior.dato
}

func (iter *iterListaEnlazada[T]) Insertar(elem T) {
	if iter.lista.EstaVacia() || iter.anterior == nil {
		iter.lista.InsertarPrimero(elem)
		iter.actual = iter.lista.primero
		return
	}
	nodo := crearNodo(elem)

	iter.anterior.siguiente = nodo
	nodo.siguiente = iter.actual
	iter.actual = nodo
	iter.lista.largo++

	if iter.actual.siguiente == nil {
		iter.lista.ultimo = nodo
	}
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	if iter.anterior == nil {
		borrado := iter.lista.BorrarPrimero()
		iter.actual = iter.lista.primero
		return borrado
	}
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	borrado := iter.actual.dato
	iter.anterior.siguiente = iter.actual.siguiente
	iter.actual = iter.actual.siguiente
	iter.lista.largo--

	return borrado
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}

}
