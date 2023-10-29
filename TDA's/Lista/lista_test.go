package lista_test

import (
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
#############################################################
#                   TESTEOS LISTA ENLAZADA                  #
#############################################################
*/

func TestListaVacia(t *testing.T) {
	t.Log("Pruebas para lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	require.EqualValues(t, 0, lista.Largo())
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

}

func TestListarElementosAlPrincipio(t *testing.T) {
	t.Log("----------PRUEBAS LISTA LLENANDOSE AL PRINCIPIO----------")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(5)
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, false, lista.EstaVacia())
	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	lista.InsertarPrimero(3)
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, false, lista.EstaVacia())
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	lista.InsertarPrimero(1)
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, false, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	lista.InsertarPrimero(8)
	require.EqualValues(t, 4, lista.Largo())
	require.EqualValues(t, false, lista.EstaVacia())
	require.EqualValues(t, 8, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
}

func TestListarElementosAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	t.Log("----------PRUEBAS LISTA LLENANDOSE AL FINAL----------")
	lista.InsertarUltimo(9)
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 9, lista.VerUltimo())
	require.EqualValues(t, 9, lista.VerPrimero())
	lista.InsertarUltimo(23)
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 23, lista.VerUltimo())
	require.EqualValues(t, 9, lista.VerPrimero())
	lista.InsertarUltimo(7)
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 7, lista.VerUltimo())
	require.EqualValues(t, 9, lista.VerPrimero())
	lista.InsertarUltimo(20)
	require.EqualValues(t, 4, lista.Largo())
	require.EqualValues(t, 20, lista.VerUltimo())
	require.EqualValues(t, 9, lista.VerPrimero())
}

func TestBorrarPrimero(t *testing.T) {
	t.Log("----------PRUEBAS DE BORRADO DE ELEMENTOS----------")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(7)
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, 7, lista.VerPrimero())
	require.EqualValues(t, 7, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 6, lista.VerPrimero())
	require.EqualValues(t, 6, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.BorrarPrimero())
	require.EqualValues(t, 0, lista.Largo())
	require.EqualValues(t, lista.EstaVacia(), true)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestVolumen(t *testing.T) {
	t.Log("----------PRUEBAS DE VOLUMEN----------")
	tam := 1000
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= tam; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
	}
	for i := tam; i >= 0; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}

	require.EqualValues(t, true, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	for i := 0; i <= tam; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
	}

	for i := 0; i <= tam; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}

	require.EqualValues(t, true, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

}

func TestStrings(t *testing.T) {
	t.Log("----------PRUEBAS DE STRINGS----------")
	lista := TDALista.CrearListaEnlazada[string]()
	require.NotNil(t, lista)
	require.EqualValues(t, true, lista.EstaVacia())
	lista.InsertarUltimo("Algoritmos")
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.EqualValues(t, "Algoritmos", lista.VerUltimo())
	lista.InsertarUltimo("Y")
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.EqualValues(t, "Y", lista.VerUltimo())
	lista.InsertarUltimo("Programación")
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.EqualValues(t, "Programación", lista.VerUltimo())
	lista.InsertarUltimo("2")
	require.EqualValues(t, "Algoritmos", lista.VerPrimero())
	require.EqualValues(t, "2", lista.VerUltimo())
	require.EqualValues(t, "Algoritmos", lista.BorrarPrimero())
	require.EqualValues(t, "Y", lista.BorrarPrimero())
	require.EqualValues(t, "Programación", lista.BorrarPrimero())
	require.EqualValues(t, "2", lista.BorrarPrimero())
	require.EqualValues(t, true, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestFloat32(t *testing.T) {
	t.Log("----------PRUEBAS DE FLOAT32----------")
	lista := TDALista.CrearListaEnlazada[float32]()
	require.NotNil(t, lista)
	require.EqualValues(t, true, lista.EstaVacia())
	lista.InsertarUltimo(1.987)
	require.EqualValues(t, 1.987, lista.VerPrimero())
	require.EqualValues(t, 1.987, lista.VerUltimo())
	lista.InsertarUltimo(8.9)
	require.EqualValues(t, 1.987, lista.VerPrimero())
	require.EqualValues(t, 8.9, lista.VerUltimo())
	lista.InsertarUltimo(5.45)
	require.EqualValues(t, 1.987, lista.VerPrimero())
	require.EqualValues(t, 5.45, lista.VerUltimo())
	lista.InsertarUltimo(7.988765)
	require.EqualValues(t, 1.987, lista.VerPrimero())
	require.EqualValues(t, 7.988765, lista.VerUltimo())
	require.EqualValues(t, 1.987, lista.BorrarPrimero())
	require.EqualValues(t, 8.9, lista.BorrarPrimero())
	require.EqualValues(t, 5.45, lista.BorrarPrimero())
	require.EqualValues(t, 7.988765, lista.BorrarPrimero())
	require.EqualValues(t, true, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

/*
#############################################################
#                   TESTEOS ITERADOR EXTERNO                #
#############################################################
*/

func TestIteradorExtVacio(t *testing.T) {
	t.Log("----------PRUEBAS DE INSERTAR ELEMENTOS----------")
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.EqualValues(t, false, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })

}

func TestHaySiguiente(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(8)
	iterador := lista.Iterador()
	require.EqualValues(t, 8, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	require.EqualValues(t, 8, iterador.Siguiente())
	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	require.EqualValues(t, 1, iterador.Siguiente())
	require.EqualValues(t, 3, iterador.VerActual())
	require.EqualValues(t, true, iterador.HaySiguiente())
	require.EqualValues(t, 3, iterador.Siguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

func TestIteradorExtInsertar(t *testing.T) {
	t.Log("----------PRUEBAS DE INSERTAR ELEMENTOS----------")
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(3)
	iterador.Insertar(41)
	iterador.Insertar(15)
	t.Log("----------INSERTAMOS AL PRINCIPIO----------")
	iterador.Insertar(7)
	require.EqualValues(t, 7, lista.VerPrimero())
	iterador.Insertar(89)
	require.EqualValues(t, 89, lista.VerPrimero())

	t.Log("----------INSERTAMOS EN EL MEDIO----------")
	for i := 0; i < 3; i++ {
		iterador.Siguiente()
	}
	iterador.Insertar(10)
	require.EqualValues(t, 10, iterador.VerActual())

	t.Log("----------INSERTAMOS AL FINAL----------")
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(5)
	require.EqualValues(t, 5, iterador.VerActual())
	require.EqualValues(t, 5, lista.VerUltimo())

}

func TestUnSoloElementoEnMedio(t *testing.T) {
	t.Log("----------ELIMINAR UN SOLO ELEMENTO Y DESPUES REVISAR LA LISTA CON OTRO ITERADOR----------")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(41)
	lista.InsertarPrimero(15)
	iterador := lista.Iterador()
	iterador.Siguiente()
	require.EqualValues(t, 41, iterador.Borrar())
	chequeo := lista.Iterador()
	require.EqualValues(t, 15, chequeo.Siguiente())
	require.EqualValues(t, 3, chequeo.VerActual())
	require.EqualValues(t, 3, chequeo.Siguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { chequeo.Siguiente() })
}

func TestELiminarDesdeElInicio(t *testing.T) {
	t.Log("----------PRUEBAS DE ELIMINAR ELEMENTOS DESDE EL INICIO AL FINAL----------")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(41)
	lista.InsertarPrimero(15)
	iterador := lista.Iterador()
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 15, iterador.VerActual())
	require.EqualValues(t, 15, iterador.Borrar())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 41, iterador.VerActual())
	require.EqualValues(t, 41, iterador.Borrar())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 3, iterador.VerActual())
	require.EqualValues(t, 3, iterador.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })

}

func TestBorrarMedio(t *testing.T) {
	t.Log("----------PRUEBAS DE ELIMINAR EN EL MEDIO----------")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(41)
	lista.InsertarPrimero(11)
	lista.InsertarPrimero(76)
	lista.InsertarPrimero(10)
	iterador := lista.Iterador()
	for i := 0; i < 2; i++ {
		iterador.Siguiente()
	}
	require.EqualValues(t, 11, iterador.VerActual())
	require.EqualValues(t, 11, iterador.Borrar())
	require.EqualValues(t, 4, lista.Largo())
	require.EqualValues(t, 41, iterador.VerActual())
	iterador.Siguiente()
	iterador.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })

	chequeo := lista.Iterador()
	require.EqualValues(t, 10, chequeo.Siguiente())
	require.EqualValues(t, 76, chequeo.Siguiente())
	require.EqualValues(t, 41, chequeo.Siguiente())
	require.EqualValues(t, 7, chequeo.VerActual())
	require.EqualValues(t, 7, chequeo.Siguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { chequeo.Siguiente() })

}

func TestELiminarDesdeElFinal(t *testing.T) {
	t.Log("----------PRUEBAS DE ELIMINAR ÚLTIMO ELEMENTO DEL FINAL----------")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(41)
	lista.InsertarPrimero(11)
	iterador := lista.Iterador()
	for iterador.VerActual() != 7 {
		iterador.Siguiente()
	}
	require.EqualValues(t, 7, iterador.VerActual())
	require.EqualValues(t, 7, iterador.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
}

/*
#############################################################
#                  TESTEOS ITERADOR INTERNO                 #
#############################################################
*/

func TestIteradorVacio(t *testing.T) {
	t.Log("----------ITERAR CON LISTA VACIA----------")
	lista := TDALista.CrearListaEnlazada[int]()

	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(elemento int) bool {
		*contador_ptr = *contador_ptr + 1
		return true
	})
	require.EqualValues(t, 0, contador)
}

func TestIteradorSinCorte(t *testing.T) {
	t.Log("----------ITERAR SIN CORTE----------")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(41)
	lista.InsertarPrimero(11)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(elemento int) bool {
		*contador_ptr = *contador_ptr + 1
		return true
	})
	require.EqualValues(t, 3, contador)
	acumulador := 0
	acumulador_ptr := &acumulador
	lista.Iterar(func(elemento int) bool {
		*acumulador_ptr = *acumulador_ptr + elemento
		return true
	})
	require.EqualValues(t, 59, acumulador)
}

func TestIteradorConCorte(t *testing.T) {
	t.Log("----------ITERAR CON CORTE----------")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(41)
	lista.InsertarPrimero(11)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(16)
	lista.InsertarPrimero(13)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	lista.Iterar(func(elemento int) bool {
		*acumulador_ptr = *acumulador_ptr + elemento
		*contador_ptr = *contador_ptr + 1
		return *contador_ptr < 4
	})
	require.EqualValues(t, 45, acumulador)
}

func TestIteradorRecorrido(t *testing.T) {
	t.Log("----------ITERAR CON VOLUMEN EN ORDEN----------")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 1000; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(elemento int) bool {
		require.EqualValues(t, *contador_ptr, elemento)
		*contador_ptr = *contador_ptr + 1
		return true
	})
}
