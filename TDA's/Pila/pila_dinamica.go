package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const CAPACIDAD_INICIAL int = 50
const TAMANIO_MINIMO = 50
const COEFICENTE_CAPACIDAD int = 2
const COEFICIENTE_REDIMENSION int = 4

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	(*pila).datos = make([]T, CAPACIDAD_INICIAL)
	(*pila).cantidad = 0
	return pila
}

func (pila pilaDinamica[T]) redimensionar(nueva_capacidad int) []T {
	arrayAuxiliar := make([]T, nueva_capacidad)
	copy(arrayAuxiliar, pila.datos)
	return arrayAuxiliar
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0

}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	elem := (pila.datos)[(pila.cantidad)-1]
	return elem

}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	capacidad := len(pila.datos)
	if capacidad == pila.cantidad {
		pila.datos = pila.redimensionar((capacidad * COEFICENTE_CAPACIDAD))
	}

	(pila.datos)[pila.cantidad] = elem
	(pila.cantidad)++

}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	elem := (pila.datos)[(pila.cantidad)-1]
	(pila.cantidad)--
	capacidad := len(pila.datos)
	if (pila.cantidad*COEFICIENTE_REDIMENSION <= capacidad) && (capacidad > TAMANIO_MINIMO) {
		pila.datos = pila.redimensionar((capacidad / COEFICENTE_CAPACIDAD))
	}
	return elem
}
