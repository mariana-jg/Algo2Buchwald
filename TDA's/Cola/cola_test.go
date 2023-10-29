package cola_test

import (
	TDACola "cola"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColaVacia(t *testing.T) {
	t.Log("----------PRUEBAS COLAS VACÍAS----------")
	cola := TDACola.CrearColaEnlazada[int]()
	require.NotNil(t, cola)
	require.EqualValues(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarElementos(t *testing.T) {
	t.Log("----------PRUEBAS COLAS LLENANDOSE----------")
	cola := TDACola.CrearColaEnlazada[int]()
	require.NotNil(t, cola)
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	cola.Encolar(5)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 5, cola.VerPrimero())
	cola.Encolar(3)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 5, cola.VerPrimero())
	cola.Encolar(1)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 5, cola.VerPrimero())
	cola.Encolar(8)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 5, cola.VerPrimero())
}

func TestDesencolarElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.NotNil(t, cola)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(6)
	cola.Encolar(7)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 3, cola.VerPrimero())
	t.Log("----------PRUEBAS COLAS VACIÁNDOSE----------")
	cola.Desencolar()
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 4, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 6, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 7, cola.VerPrimero())
	cola.Desencolar()
	require.EqualValues(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestVolumen(t *testing.T) {
	t.Log("----------PRUEBAS DE VOLUMEN----------")
	tam := 1000
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i <= tam; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}

	for i := tam; i >= 0; i-- {
		require.EqualValues(t, tam-i, cola.VerPrimero())
		cola.Desencolar()
	}
	require.EqualValues(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestStrings(t *testing.T) {
	t.Log("----------PRUEBAS DE STRINGS----------")
	cola := TDACola.CrearColaEnlazada[string]()
	require.NotNil(t, cola)
	cola.Encolar("Algoritmos")
	cola.Encolar("Y")
	cola.Encolar("Programación")
	cola.Encolar("2")
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, "Algoritmos", cola.VerPrimero())
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()
	require.EqualValues(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestFloat32(t *testing.T) {
	t.Log("----------PRUEBAS DE FLOAT32----------")
	cola := TDACola.CrearColaEnlazada[float32]()
	require.NotNil(t, cola)
	cola.Encolar(1.775)
	cola.Encolar(8.9877)
	cola.Encolar(6.5)
	cola.Encolar(18.98)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 1.775, cola.VerPrimero())
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()
	require.EqualValues(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}
