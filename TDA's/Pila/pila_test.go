package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	t.Log("----------PRUEBAS PILAS VACÍAS----------")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.NotNil(t, pila)
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarElementos(t *testing.T) {
	t.Log("----------PRUEBAS PILAS LLENANDOSE----------")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.NotNil(t, pila)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	pila.Apilar(5)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 5, pila.VerTope())
	pila.Apilar(3)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 3, pila.VerTope())
	pila.Apilar(1)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 1, pila.VerTope())
	pila.Apilar(8)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 8, pila.VerTope())
}

func TestDesapilarElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.NotNil(t, pila)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(6)
	pila.Apilar(7)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 7, pila.VerTope())
	t.Log("----------PRUEBAS PILAS VACIÁNDOSE----------")
	pila.Desapilar()
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 6, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 4, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 3, pila.VerTope())
	pila.Desapilar()
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestVolumen(t *testing.T) {
	t.Log("----------PRUEBAS DE VOLUMEN----------")
	tam := 1000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}

	for i := tam; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		pila.Desapilar()
	}
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestStrings(t *testing.T) {
	t.Log("----------PRUEBAS DE STRINGS----------")
	pila := TDAPila.CrearPilaDinamica[string]()
	require.NotNil(t, pila)
	pila.Apilar("Algoritmos")
	pila.Apilar("Y")
	pila.Apilar("Programación")
	pila.Apilar("2")
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, "2", pila.VerTope())
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestFloat32(t *testing.T) {
	t.Log("----------PRUEBAS DE FLOAT32----------")
	pila := TDAPila.CrearPilaDinamica[float32]()
	require.NotNil(t, pila)
	pila.Apilar(1.775)
	pila.Apilar(8.9877)
	pila.Apilar(6.5)
	pila.Apilar(18.98)
	require.EqualValues(t, false, pila.EstaVacia())
	require.EqualValues(t, 18.98, pila.VerTope())
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	require.EqualValues(t, true, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}
