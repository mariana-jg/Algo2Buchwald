package cola_prioridad

const FACTOR_REDIMENSION = 2
const TAMANIO_INICIAL = 11

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func swap[T comparable](dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato2, *dato1
}

func upheap[T comparable](cantidad int, datos []T, posicion int, cmp func(T, T) int) {
	if posicion == 0 {
		return
	}
	posicionPadre := (posicion - 1) / 2
	padre := datos[posicionPadre]
	hijo := datos[posicion]
	if cmp(padre, hijo) > 0 {
		return
	}
	swap(&datos[posicionPadre], &datos[posicion])
	upheap(cantidad, datos, posicionPadre, cmp)
}

func downheap[T comparable](cantidad int, datos []T, posicion int, cmp func(T, T) int) {
	if posicion >= cantidad {
		return
	}
	posicionHijoIzq := 2*posicion + 1
	posicionHijoDer := 2*posicion + 2
	posicionMaximo := posicion
	if posicionHijoIzq < cantidad && cmp(datos[posicionHijoIzq], datos[posicionMaximo]) > 0 {
		posicionMaximo = posicionHijoIzq
	}
	if posicionHijoDer < cantidad && cmp(datos[posicionHijoDer], datos[posicionMaximo]) > 0 {
		posicionMaximo = posicionHijoDer
	}
	if posicion != posicionMaximo {
		swap(&datos[posicion], &datos[posicionMaximo])
		downheap(cantidad, datos, posicionMaximo, cmp)
	}
}

func heapify[T comparable](cantidad int, datos []T, posicion int, cmp func(T, T) int) {
	for i := 0; i < len(datos); i++ {
		upheap(cantidad, datos, i, cmp)
	}
}

func redimensionar[T comparable](heap heap[T], nuevaCapacidad int) []T {
	arrayAuxiliar := make([]T, nuevaCapacidad)
	copy(arrayAuxiliar, heap.datos)
	return arrayAuxiliar
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, TAMANIO_INICIAL)
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, len(arreglo))
	heap.cantidad = len(arreglo)
	heap.cmp = funcion_cmp
	if len(arreglo) > 0 {
		copy(heap.datos, arreglo)
		heapify(heap.cantidad, heap.datos, heap.cantidad, heap.cmp)
	}
	return heap
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(dato T) {
	capacidad := len(heap.datos)
	if capacidad == 0 {
		heap.datos = redimensionar(*heap, TAMANIO_INICIAL)
	} else if capacidad == heap.cantidad {
		heap.datos = redimensionar(*heap, capacidad*FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = dato
	upheap(heap.cantidad, heap.datos, heap.cantidad, heap.cmp)
	heap.cantidad++
}

func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	maximo := heap.VerMax()
	heap.datos[0] = heap.datos[heap.cantidad-1]
	heap.cantidad--
	downheap(heap.cantidad, heap.datos, 0, heap.cmp)
	if heap.cantidad == len(heap.datos)/4 {
		redimensionar(*heap, len(heap.datos)/FACTOR_REDIMENSION)
	}
	return maximo
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(len(elementos), elementos, 0, funcion_cmp)
	posicionDeSwap := len(elementos) - 1
	for i := posicionDeSwap; i >= 0; i-- {
		swap(&elementos[0], &elementos[i])
		downheap(i, elementos, 0, funcion_cmp)
	}
}
