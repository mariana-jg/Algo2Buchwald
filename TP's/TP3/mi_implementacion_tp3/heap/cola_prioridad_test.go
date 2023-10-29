package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func mixVector(array []int, largo int) {
	for i := largo - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func comparacionInt(n1 int, n2 int) int {
	if n1 > n2 {
		return 1
	}
	if n1 < n2 {
		return -1
	}
	return 0
}

func TestColaVacia(t *testing.T) {
	t.Log("----------PRUEBAS COLA VACÍA----------")
	heap := TDAHeap.CrearHeap(comparacionInt)
	require.NotNil(t, heap)
	require.EqualValues(t, true, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarPocosElementos(t *testing.T) {
	t.Log("----------PRUEBAS DE ENCOLAR POCOS ELEMENTOS----------")
	heap := TDAHeap.CrearHeap(comparacionInt)
	dato := [6]int{14, 34, 243, 244, 252, 678}

	for i := 0; i < len(dato); i++ {
		heap.Encolar(dato[i])
		require.EqualValues(t, false, heap.EstaVacia())
		require.EqualValues(t, i+1, heap.Cantidad())
		require.EqualValues(t, dato[i], heap.VerMax())
	}

}

func TestDesencolarPocosElementos(t *testing.T) {
	t.Log("----------PRUEBAS DE DESENCOLAR POCOS ELEMENTOS----------")
	heap := TDAHeap.CrearHeap(comparacionInt)
	dato1 := 14
	dato2 := 34
	dato3 := 244
	dato4 := 243
	dato5 := 232
	dato6 := 678
	heap.Encolar(dato1)
	heap.Encolar(dato2)
	heap.Encolar(dato3)
	heap.Encolar(dato4)
	heap.Encolar(dato5)
	heap.Encolar(dato6)
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, dato6, heap.Desencolar())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, dato3, heap.Desencolar())
	require.EqualValues(t, dato4, heap.Desencolar())
	require.EqualValues(t, dato5, heap.Desencolar())
	require.EqualValues(t, dato2, heap.Desencolar())
	require.EqualValues(t, dato1, heap.Desencolar())
	require.EqualValues(t, true, heap.EstaVacia())
}

func TestEncolarDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap(comparacionInt)
	dato1 := 14
	dato2 := 244
	dato3 := 10000
	dato4 := 0

	heap.Encolar(dato1)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, dato1, heap.VerMax())
	heap.Encolar(dato4)
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, dato1, heap.VerMax())
	require.EqualValues(t, dato1, heap.Desencolar())
	require.EqualValues(t, 1, heap.Cantidad())
	heap.Encolar(dato3)
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, dato3, heap.VerMax())
	heap.Encolar(dato2)
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, dato3, heap.VerMax())
	require.EqualValues(t, dato3, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())

}

func TestEncolarDesencolarVolumen(t *testing.T) {
	var arreglo []int
	heap := TDAHeap.CrearHeap(comparacionInt)
	for i := 0; i < 10001; i++ {
		arreglo = append(arreglo, i)
	}
	mixVector(arreglo, len(arreglo))
	for i := 0; i < len(arreglo); i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(arreglo[i])
		require.EqualValues(t, false, heap.EstaVacia())
		require.EqualValues(t, i+1, heap.Cantidad())
	}
	require.EqualValues(t, 10000, heap.VerMax())
	for i := len(arreglo) - 1; i >= 0; i-- {
		require.EqualValues(t, i+1, heap.Cantidad())
		require.EqualValues(t, false, heap.EstaVacia())
		require.EqualValues(t, i, heap.Desencolar())
		require.EqualValues(t, i, heap.Cantidad())
	}
	require.EqualValues(t, true, heap.EstaVacia())

}

func TestCrearArrVacio(t *testing.T) {
	var array []int
	heap := TDAHeap.CrearHeapArr(array, comparacionInt)
	require.EqualValues(t, true, heap.EstaVacia())
	heap.Encolar(3)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
}

func TestCrearHeapConArray(t *testing.T) {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mixVector(array, len(array))

	heap := TDAHeap.CrearHeapArr(array, comparacionInt)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 20, heap.Cantidad())
	require.EqualValues(t, 20, heap.VerMax())

	for i := len(array); i > 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.EqualValues(t, true, heap.EstaVacia())
}

func TestCrearHeapConArrayVolumen(t *testing.T) {
	var array []int
	for i := 0; i < 100001; i++ {
		array = append(array, i)
	}
	mixVector(array, len(array))
	heap := TDAHeap.CrearHeapArr(array, comparacionInt)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 100001, heap.Cantidad())
	require.EqualValues(t, 100000, heap.VerMax())

	for i := len(array) - 1; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.EqualValues(t, true, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	var array []int
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	mixVector(array, len(array))
	TDAHeap.HeapSort(array, comparacionInt)

	for i := 0; i < len(array); i++ {
		require.EqualValues(t, i, array[i])
	}
}

func TestHeapSortVolumen(t *testing.T) {
	var array []int
	for i := 0; i < 100000; i++ {
		array = append(array, i)
	}
	mixVector(array, len(array))
	TDAHeap.HeapSort(array, comparacionInt)

	for i := 0; i < len(array); i++ {
		require.EqualValues(t, i, array[i])
	}
}
