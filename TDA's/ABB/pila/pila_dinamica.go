package pila

const _TAMANIOINICIAL int = 10

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {

	pila := new(pilaDinamica[T])
	(*pila).datos = make([]T, _TAMANIOINICIAL)
	return pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	if pila.cantidad == 0 {
		return true
	}
	return false
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() == true {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila pilaDinamica[T]) redimensionarPila(nuevoTam int) []T {
	aux := make([]T, nuevoTam)
	copy(aux, pila.datos)
	return aux
}

func (pila *pilaDinamica[T]) Apilar(elem T) {

	if len(pila.datos) == pila.cantidad {
		pila.datos = pila.redimensionarPila(2 * pila.cantidad)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++

}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() == true {
		panic("La pila esta vacia")
	}
	if 4*pila.cantidad <= len(pila.datos) {
		pila.datos = pila.redimensionarPila(len(pila.datos) / 2)
	}
	pila.cantidad--
	return pila.datos[pila.cantidad]
}
