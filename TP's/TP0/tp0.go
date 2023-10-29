package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {

	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {

	if len(vector) == 0 {
		return -1
	}

	mayor := vector[0]
	posicionMayor := 0

	for i := 0; i < len(vector); i++ {
		if vector[i] > mayor && vector[i] != mayor {
			mayor = vector[i]
			posicionMayor = i
		}
	}
	return posicionMayor
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {

	largoVector1 := len(vector1)
	largoVector2 := len(vector2)

	for i := 0; i < largoVector2 && i < largoVector1; i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	if largoVector1 == largoVector2 {
		return 0
	} else if largoVector1 < largoVector2 {
		return -1
	} else {
		return 1
	}
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {

	ultimaPosicion := len(vector) - 1
	var posicionMayor int
	var slice []int = vector[:]

	for ultimaPosicion > 0 {
		posicionMayor = Maximo(slice)
		Swap(&vector[posicionMayor], &vector[ultimaPosicion])
		slice = vector[:ultimaPosicion]
		ultimaPosicion = ultimaPosicion - 1
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {

	if len(vector) == 0 {
		return 0
	} else {
		return vector[0] + Suma(vector[1:])
	}
}

func EsPalindromoRecursivo(cadena string, inicio int, final int) bool {

	if inicio >= final {
		return true
	}
	if cadena[inicio] == cadena[final] {
		return EsPalindromoRecursivo(cadena, inicio+1, final-1)
	}
	return false
}

// EsPalindromo devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA.
func EsPalindromo(cadena string) bool {
	return EsPalindromoRecursivo(cadena, 0, len(cadena)-1)
}
