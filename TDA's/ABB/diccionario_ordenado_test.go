package diccionario_test

import (
	TDADiccionario "diccionario"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func mezclarStrings(claves []string, largo int) {

	for i := largo - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		claves[i], claves[j] = claves[j], claves[i]
	}
}

//******************************************************************************************
//* FUNCIONES DE COMPARACIÓN
//******************************************************************************************

func comparacionInt(n1 int, n2 int) int {
	if n1 > n2 {
		return 1
	}
	if n1 < n2 {
		return -1
	}
	return 0
}

func comparacionString(s1 string, s2 string) int {
	if s1 > s2 {
		return 1
	}
	if s1 < s2 {
		return -1
	}
	return 0
}

//******************************************************************************************
//* FUNCIONES AUXILIARES
//******************************************************************************************

func crearIngresos() ([]string, []string) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Delfin"
	clave5 := "Rana"
	clave6 := "Pajaro"
	clave7 := "Ballena"
	clave8 := "Foca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "niiiiii"
	valor5 := "rabbit"
	valor6 := "piopiopio"
	valor7 := "Splash"
	valor8 := "Clap"
	claves := []string{clave1, clave2, clave3, clave4, clave5, clave6, clave7, clave8}
	valores := []string{valor1, valor2, valor3, valor4, valor5, valor6, valor7, valor8}
	return claves, valores
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

//******************************************************************************************
//* PRUEBAS DEL DICCIONARIO
//******************************************************************************************

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(6))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(4) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(2) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](comparacionString)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	claves, valores := crearIngresos()
	dic := TDADiccionario.CrearABB[string, string](comparacionString)

	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	require.EqualValues(t, 8, dic.Cantidad())
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave1 := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](comparacionString)

	dic.Guardar(clave1, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave1))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave1, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave1))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioBorrarUnELementoSinHIjos(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	valor1 := "miau"
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	require.False(t, dic.Pertenece(clave1))
	dic.Guardar(clave1, valor1)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave1))
	require.EqualValues(t, valor1, dic.Borrar(clave1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(clave1) })
}

func TestDiccionarioBorrarUnELementoConUnHijo(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	valor1 := "miau"
	valor2 := "guau"
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	require.EqualValues(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, valor1, dic.Borrar(clave1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(clave1) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
}

func TestDiccionarioBorrarUnELementoConDosHijos(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	require.EqualValues(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
}

func TestArbolComplejo(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	claves, valores := crearIngresos()
	dic := TDADiccionario.CrearABB[string, string](comparacionString)

	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[3], valores[3])
	dic.Guardar(claves[4], valores[4])
	dic.Guardar(claves[5], valores[5])

	require.EqualValues(t, 6, dic.Cantidad())

	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.True(t, dic.Pertenece(claves[3]))
	require.True(t, dic.Pertenece(claves[4]))
	require.True(t, dic.Pertenece(claves[5]))

	require.EqualValues(t, valores[3], dic.Borrar(claves[3]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[3]) })
	require.EqualValues(t, 5, dic.Cantidad())

	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.True(t, dic.Pertenece(claves[4]))
	require.True(t, dic.Pertenece(claves[5]))

	dic.Guardar(claves[3], valores[3])
	require.EqualValues(t, 6, dic.Cantidad())

	require.EqualValues(t, valores[4], dic.Borrar(claves[4]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[4]) })

}

func TestDeVolumenGuardadoYBorrado(t *testing.T) {
	t.Log("Guarda 500 elementos en el diccionario, y se comprueba que en todo momento funciona acorde")

	dic := TDADiccionario.CrearABB[string, *int](comparacionString)

	claves := make([]string, 500)
	valores := make([]int, 500)

	for i := 0; i < 500; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	mezclarStrings(claves, 500)

	for i := 0; i < 500; i++ {
		dic.Guardar(claves[i], &valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	require.EqualValues(t, 500, dic.Cantidad())

	for i := 0; i < 500; i++ {
		require.EqualValues(t, &valores[i], dic.Borrar(claves[i]))
		require.False(t, dic.Pertenece(claves[i]))
	}

	require.EqualValues(t, 0, dic.Cantidad())

	for i := 0; i < 500; i++ {
		dic.Guardar(claves[i], &valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}
	require.EqualValues(t, 500, dic.Cantidad())
}

//******************************************************************************************
//* PRUEBAS DEL ITERADOR INTERNO IN-ORDER
//******************************************************************************************

func TestIteradorVacio(t *testing.T) {
	t.Log("Se itera sobre un diccionario vacío y se comprueba que se comporte como tal.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	contador := 0
	contador_ptr := &contador
	dic.Iterar(func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		return true
	})
	require.EqualValues(t, 0, contador)
}

func TestIteradorInternoSinCorte(t *testing.T) {
	t.Log("Se itera sobre un diccionario con varios elementos y se comprueba que itere sin condición de corte correctamente.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(1, 1)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	dic.Iterar(func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		*acumulador_ptr = *acumulador_ptr + clave
		return true
	})
	require.EqualValues(t, 7, contador)
	require.EqualValues(t, 33, acumulador)
}

func TestIteradorInternoConCorte(t *testing.T) {
	t.Log("Se itera sobre un diccionario con varios elementos y se comprueba que itere con condición de corte correctamente.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(1, 1)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	dic.Iterar(func(clave int, dato int) bool {
		if contador >= 6 {
			return false
		}
		*acumulador_ptr = *acumulador_ptr + clave
		*contador_ptr = *contador_ptr + 1
		return true
	})
	require.EqualValues(t, 24, acumulador)
}

func TestDeVolumenIterInterno(t *testing.T) {
	t.Log("Guarda 500 elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	dic := TDADiccionario.CrearABB[string, *int](comparacionString)

	claves := make([]string, 500)
	valores := make([]int, 500)

	for i := 0; i < 500; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	mezclarStrings(claves, 500)

	for i := 0; i < 500; i++ {
		dic.Guardar(claves[i], &valores[i])

	}
	contador := 0
	contador_ptr := &contador
	dic.Iterar(func(clave string, dato *int) bool {
		*contador_ptr++
		return true
	})
	require.EqualValues(t, 500, contador)

}

// ******************************************************************************************
// * PRUEBAS DEL ITERADOR INTERNO IN-ORDER POR RANGO
// ******************************************************************************************

func TestIterarConRangoEnVacio(t *testing.T) {
	t.Log("Se itera sobre un diccionario vacío y se comprueba que se comporte como tal.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	contador := 0
	contador_ptr := &contador
	desde := 1
	hasta := 9
	dic.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		return true
	})
	require.EqualValues(t, 0, contador)
}

func TestIteradorInternoConRangoCompleto(t *testing.T) {
	t.Log("Se itera sobre un diccionario con varios elementos teniendo un rango que abarque todo el diccionario y sin condición de corte," +
		"esperando que se comporte de manera correcta.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(1, 1)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	desde := 1
	hasta := 9
	dic.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		*acumulador_ptr = *acumulador_ptr + clave
		return true
	})
	require.EqualValues(t, 7, contador)
	require.EqualValues(t, 33, acumulador)

}
func TestIteradorInternoConRangoAcotado(t *testing.T) {
	t.Log("Se itera sobre un diccionario con varios elementos teniendo un rango que abarque parte del diccionario y sin condición de corte," +
		"esperando que se comporte de manera correcta.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(1, 1)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	desde := 2
	hasta := 8
	dic.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		*acumulador_ptr = *acumulador_ptr + clave
		return true
	})
	require.EqualValues(t, 5, contador)
	require.EqualValues(t, 23, acumulador)
}

func TestIteradorInternoDeRangoNil(t *testing.T) {
	t.Log("Se itera sobre un diccionario con varios elementos teniendo un rango que abarque parte del diccionario y sin condición de corte," +
		"esperando que se comporte de manera correcta.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(1, 1)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	dic.IterarRango(nil, nil, func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		*acumulador_ptr = *acumulador_ptr + clave
		return true
	})
	require.EqualValues(t, 7, contador)
	require.EqualValues(t, 33, acumulador)
}

func TestIteradorInternoSinDesde(t *testing.T) {
	t.Log("Se itera sobre un diccionario con varios elementos teniendo un rango que abarque parte del diccionario y sin condición de corte," +
		"esperando que se comporte de manera correcta.")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(1, 1)
	dic.Guardar(15, 15)
	dic.Guardar(4, 4)
	dic.Guardar(20, 20)
	contador := 0
	contador_ptr := &contador
	acumulador := 0
	acumulador_ptr := &acumulador
	hasta := 13
	dic.IterarRango(nil, &hasta, func(clave int, dato int) bool {
		*contador_ptr = *contador_ptr + 1
		*acumulador_ptr = *acumulador_ptr + clave
		return true
	})
	require.EqualValues(t, 8, contador)
	require.EqualValues(t, 37, acumulador)
}

//******************************************************************************************
//* PRUEBAS DEL ITERADOR EXTERNO IN-ORDER
//******************************************************************************************

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.EqualValues(t, primero, iter.Siguiente())
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, segundo, iter.Siguiente())
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	require.EqualValues(t, tercero, iter.Siguiente())

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero := iter3.Siguiente()
	segundo := iter3.Siguiente()
	tercero := iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, los borra y se comprueba que el recorrido que hace el iterador sea correcto")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	require.EqualValues(t, clave1, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
}

func TestRecorridoInorder(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que el recorrido que hace el iterador sea in-order")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves, valores := crearIngresos()
	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}
	orden := []string{claves[6], claves[3], claves[7], claves[0], claves[5], claves[1], claves[4], claves[2]}
	iter := dic.Iterador()
	for i := 0; i < 8; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, orden[i], dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, orden[i], iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDeVolumenIterador(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, *int](comparacionString)

	claves := make([]string, 500)
	valores := make([]int, 500)

	for i := 0; i < 500; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	mezclarStrings(claves, 500)

	for i := 0; i < 500; i++ {
		dic.Guardar(claves[i], &valores[i])
	}

	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < 500; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = 500
		iter.Siguiente()
	}
	require.True(t, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(t, 500, i, "No se recorrió todo el largo")
	require.False(t, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < 500; i++ {
		if valores[i] != 500 {
			ok = false
			break
		}
	}
	require.True(t, ok, "No se cambiaron todos los elementos")
}

//******************************************************************************************
//* PRUEBAS DEL ITERADOR EXTERNO IN-ORDER POR RANGO
//******************************************************************************************

func TestIteradorExternoRangoVacio(t *testing.T) {
	t.Log("Iterar con diccionario vacio")
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	desde := 3
	hasta := 76
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorExternoRangoUnParDeElementos(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que el recorrido que hace el iterador con rango sea in-order")

	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves, valores := crearIngresos()
	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	orden := []string{claves[6], claves[3], claves[7], claves[0]}
	iter := dic.IteradorRango(&claves[6], &claves[0])
	for i := 0; i < 4; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, orden[i], dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, orden[i], iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorExternoRangoVariosElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves, valores := crearIngresos()
	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	orden := []string{claves[6], claves[3], claves[7], claves[0], claves[5], claves[1]}
	iter := dic.IteradorRango(nil, &claves[1])
	for i := 0; i < 6; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, orden[i], dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, orden[i], iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoRangoVariosElementosSinDesde(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves, valores := crearIngresos()
	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	orden := []string{claves[6], claves[3], claves[7], claves[0], claves[5], claves[1], claves[4], claves[2]}
	iter := dic.IteradorRango(nil, &claves[2])
	for i := 0; i < 8; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, orden[i], dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, orden[i], iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestCombinacionI(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(6, 2)
	dic.Guardar(3, 1)
	dic.Guardar(8, 3)
	dic.Guardar(1, 4)
	dic.Guardar(4, 5)
	dic.Guardar(7, 6)
	dic.Guardar(2, 9)
	dic.Guardar(5, 8)
	dic.Guardar(9, 7)
	dic.Guardar(10, 10)
	rango1, rango2 := 2, 5
	iter := dic.IteradorRango(&rango1, &rango2)

	for i := 2; i < 6; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, i, dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, i, iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestCombinacionII(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(4, 4)
	dic.Guardar(3, 3)
	dic.Guardar(1, 1)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)
	rango1, rango2 := 2, 5
	iter := dic.IteradorRango(&rango1, &rango2)

	for i := 2; i < 6; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, i, dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, i, iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}
func TestCombinacionIII(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)
	rango1, rango2 := 1, 4

	iter := dic.IteradorRango(&rango1, &rango2)
	for i := 1; i < 5; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, i, dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, i, iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestIteradorRangoSinHasta(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves, valores := crearIngresos()
	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	orden := []string{claves[7], claves[0], claves[5], claves[1], claves[4], claves[2]}
	iter := dic.IteradorRango(&claves[7], nil)
	for i := 0; i < 6; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, orden[i], dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, orden[i], iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorSinDesdeYSinHastaEsNormal(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que el iterador sin desde y sin hasta se comporte como iterador normal.")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	claves, valores := crearIngresos()
	for i := 0; i < 8; i++ {
		dic.Guardar(claves[i], valores[i])
		require.True(t, dic.Pertenece(claves[i]))
	}

	orden := []string{claves[6], claves[3], claves[7], claves[0], claves[5], claves[1], claves[4], claves[2]}
	iter := dic.IteradorRango(nil, nil)
	for i := 0; i < 8; i++ {
		dato, _ := iter.VerActual()
		require.EqualValues(t, orden[i], dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, orden[i], iter.Siguiente())
	}

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
