package lista

type IteradorLista[T any] interface {

	// Ver actual devuelve el elemento donde esta actualmente el iterador. Si es nil entra en panico con el mensaje
	// "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve true si el elemento actual de la lista es válido, false en caso contrario.
	HaySiguiente() bool

	// Siguiente devuelve el elemento donde se posiciona actualmente el iterador.
	// Si no hay siguiente, entra en pánico con el mensaje "El iterador termino de iterar.
	Siguiente() T

	// Insertar ubica en la lista un elemento.
	// Si actual esta en inicio -> coloca al elemento al inicio.
	// Si actual esta en el medio -> se coloca luego del actual y ahora el actual es el insertado.
	// Si actual esta al final -> coloca al elemento al final.
	Insertar(T)

	// Borrar elimina un elemento de la lista.
	// Si actual está en inicio -> el iterador irá al siguiente y se borrará el primero de la lista.
	// Si actual está al final -> se borrará el último elemento de la lista y el iterador llegara a su fin.
	// Si actual está en el medio -> el iterador irá al siguiente y se eliminará ese elemento.
	// Si la lista esta vacia o actual es una posicion no valida entra en pánico con el mensaje "El iterador termino de iterar".
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// Agrega un elemento al principio de la lista.
	InsertarPrimero(T)

	// Agrega un elemento al final de la lista.
	InsertarUltimo(T)

	// Elimina el primer elemento de la lista. Si la lista tiene elementos, se quita el primero y se
	// devuelve el dato. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Obtiene la cantidad de elementos en la lista.
	Largo() int

	// Iterador crea un iterador para la lista. Solo podra recorrerla de inicio a final, no en sentido contrario.
	Iterador() IteradorLista[T]

	// Iterar permite la implementacion de funciones a partir del recorrido de los elementos de la lista.
	Iterar(visitar func(T) bool)
}
