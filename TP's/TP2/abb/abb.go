package diccionario

import TDAPila "algogram/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterador[K comparable, V any] struct {
	elementos TDAPila.Pila[nodoAbb[K, V]]
}

type iteradorRango[K comparable, V any] struct {
	desde     *K
	hasta     *K
	elementos TDAPila.Pila[nodoAbb[K, V]]
	cmp       func(K, K) int
}

func crearNodo[K comparable, V any](clave K, valor V) nodoAbb[K, V] {
	nuevoNodo := new(nodoAbb[K, V])
	nuevoNodo.clave = clave
	nuevoNodo.dato = valor
	nuevoNodo.izquierdo = nil
	nuevoNodo.derecho = nil
	return *nuevoNodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dic := new(abb[K, V])
	dic.cantidad = 0
	dic.raiz = nil
	dic.cmp = funcion_cmp
	return dic
}

// buscarNodo devuelve nil en caso de no encontrar un nodo que contenga la misma clave. En caso que lo
// encuentre devuelve un puntero a dicho nodo
func buscarNodo[K comparable, V any](comp func(K, K) int, nodoActual *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if nodoActual == nil {
		return nodoActual
	}
	comparacion := comp(clave, nodoActual.clave)
	if comparacion == 0 {
		return nodoActual
	} else if comparacion > 0 {
		return buscarNodo(comp, nodoActual.derecho, clave)
	}
	return buscarNodo(comp, nodoActual.izquierdo, clave)
}

func (arbol abb[K, V]) Pertenece(clave K) bool {
	return buscarNodo(arbol.cmp, arbol.raiz, clave) != nil
}

func (arbol abb[K, V]) Obtener(clave K) V {
	nodo := buscarNodo(arbol.cmp, arbol.raiz, clave)
	if nodo != nil {
		return nodo.dato
	}
	panic("La clave no pertenece al diccionario")
}
func (arbol abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

// encontrarPadre devuelve el padre del nodo con la clave ingresado. En caso de que no lo encuentre devuelve
// un posible padre para guardar un nodo con esa clave a menos que sea la raiz
func encontrarPadre[K comparable, V any](arbol abb[K, V], nodoActual *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if arbol.cmp(nodoActual.clave, clave) == 0 {
		return nil
	}
	if nodoActual.derecho != nil && arbol.cmp(nodoActual.derecho.clave, clave) == 0 ||
		nodoActual.izquierdo != nil && arbol.cmp(nodoActual.izquierdo.clave, clave) == 0 {
		return nodoActual
	}
	comparacion := arbol.cmp(nodoActual.clave, clave)
	if comparacion > 0 && nodoActual.izquierdo != nil {
		if nodoActual.izquierdo == nil {
			return nodoActual
		}
		return encontrarPadre(arbol, nodoActual.izquierdo, clave)
	} else {
		if nodoActual.derecho == nil {
			return nodoActual
		}
		return encontrarPadre(arbol, nodoActual.derecho, clave)
	}
}

func guardarRec[K comparable, V any](arbol *abb[K, V], nodoActual *nodoAbb[K, V], nodoNuevo *nodoAbb[K, V]) {
	if arbol.raiz == nil {
		arbol.raiz = nodoNuevo
		arbol.cantidad++
		return
	} else if arbol.Pertenece(nodoNuevo.clave) {
		nodoAux := buscarNodo(arbol.cmp, arbol.raiz, nodoNuevo.clave)
		nodoAux.dato = nodoNuevo.dato
		return
	}
	if arbol.cmp(nodoNuevo.clave, nodoActual.clave) > 0 && nodoActual.derecho == nil {
		nodoActual.derecho = nodoNuevo
		arbol.cantidad++
		return
	} else if arbol.cmp(nodoNuevo.clave, nodoActual.clave) < 0 && nodoActual.izquierdo == nil {
		nodoActual.izquierdo = nodoNuevo
		arbol.cantidad++
		return
	}

	if arbol.cmp(nodoNuevo.clave, nodoActual.clave) > 0 && nodoActual.derecho != nil {
		guardarRec(arbol, nodoActual.derecho, nodoNuevo)
	} else if arbol.cmp(nodoNuevo.clave, nodoActual.clave) < 0 && nodoActual.izquierdo != nil {
		guardarRec(arbol, nodoActual.izquierdo, nodoNuevo)
	}

}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nodo := crearNodo(clave, dato)
	guardarRec(arbol, arbol.raiz, &nodo)
}

func calcularCantHijos[K comparable, V any](nodo *nodoAbb[K, V]) int {
	if nodo.izquierdo != nil && nodo.derecho != nil {
		return 2
	}
	if (nodo.izquierdo == nil && nodo.derecho != nil) || (nodo.izquierdo != nil && nodo.derecho == nil) {
		return 1
	}
	return 0
}

func buscarReemplazante[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	actual := nodo.izquierdo
	for actual.derecho != nil {
		actual = actual.derecho
	}
	return actual
}

func borrarRec[K comparable, V any](arbol *abb[K, V], nodo *nodoAbb[K, V], padre *nodoAbb[K, V], hijos int) {
	if hijos == 0 {
		if padre == nil {
			arbol.raiz = nil
		} else if padre.derecho != nil && arbol.cmp(padre.derecho.clave, nodo.clave) == 0 {
			padre.derecho = nil
		} else if padre.izquierdo != nil && arbol.cmp(padre.izquierdo.clave, nodo.clave) == 0 {
			padre.izquierdo = nil
		}
	} else if hijos == 1 {
		if nodo.izquierdo != nil {
			if padre == nil {
				arbol.raiz = nodo.izquierdo
			} else if padre.derecho != nil && arbol.cmp(padre.derecho.clave, nodo.clave) == 0 {
				padre.derecho = nodo.izquierdo
			} else if padre.izquierdo != nil && arbol.cmp(padre.izquierdo.clave, nodo.clave) == 0 {
				padre.izquierdo = nodo.izquierdo
			}
		} else if nodo.derecho != nil {
			if padre == nil {
				arbol.raiz = nodo.derecho
			} else if padre.derecho != nil && arbol.cmp(padre.derecho.clave, nodo.clave) == 0 {
				padre.derecho = nodo.derecho
			} else if padre.izquierdo != nil && arbol.cmp(padre.izquierdo.clave, nodo.clave) == 0 {
				padre.izquierdo = nodo.derecho
			}
		}
	} else {
		auxiliar := buscarReemplazante(nodo)
		borrarRec(arbol, auxiliar, encontrarPadre(*arbol, arbol.raiz, auxiliar.clave), calcularCantHijos(auxiliar))
		nodo.dato = auxiliar.dato
		nodo.clave = auxiliar.clave

	}
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	if !arbol.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo := buscarNodo(arbol.cmp, arbol.raiz, clave)
	aux := nodo.dato
	hijos := calcularCantHijos(nodo)
	padre := encontrarPadre(*arbol, arbol.raiz, clave)

	borrarRec(arbol, nodo, padre, hijos)
	arbol.cantidad--
	return aux
}

//******************************************************************************************
//* PRIMITIVA DEL ITERADOR INTERNO IN-ORDER
//******************************************************************************************

func iterar[K comparable, V any](visitar func(K, V) bool, nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	iterar(visitar, nodo.izquierdo)
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	iterar(visitar, nodo.derecho)

}

func (arbol abb[K, V]) Iterar(visitar func(K, V) bool) {
	iterar(visitar, arbol.raiz)
}

//******************************************************************************************
//* PRIMITIVA DEL ITERADOR INTERNO POR RANGO IN-ORDER
//******************************************************************************************

func buscarMenor[K comparable, V any](nodo *nodoAbb[K, V]) *K {
	actual := nodo
	for actual.izquierdo != nil {
		actual = actual.izquierdo
	}
	return &actual.clave
}

func buscarMayor[K comparable, V any](nodo *nodoAbb[K, V]) *K {
	actual := nodo
	for actual.derecho != nil {
		actual = actual.derecho
	}
	return &actual.clave
}

func iterarSobreRango[K comparable, V any](comp func(K, K) int, desde *K, hasta *K,
	visitar func(clave K, dato V) bool, nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}

	if comp(nodo.clave, *desde) >= 0 {
		iterarSobreRango(comp, desde, hasta, visitar, nodo.izquierdo)

	}
	if comp(nodo.clave, *desde) >= 0 && comp(nodo.clave, *hasta) <= 0 {
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}
	if comp(nodo.clave, *hasta) <= 0 {
		iterarSobreRango(comp, desde, hasta, visitar, nodo.derecho)
	}
}

func (arbol abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if desde == nil {
		desde = buscarMenor(arbol.raiz)
	}
	if hasta == nil {
		hasta = buscarMayor(arbol.raiz)
	}
	iterarSobreRango(arbol.cmp, desde, hasta, visitar, arbol.raiz)
}

//******************************************************************************************
//* PRIMITIVA DEL ITERADOR EXTERNO IN-ORDER
//******************************************************************************************

func apilarHijosIzquierdos[K comparable, V any](elementos TDAPila.Pila[nodoAbb[K, V]], actual *nodoAbb[K, V]) {
	for actual.izquierdo != nil {
		actual = actual.izquierdo
		elementos.Apilar(*actual)
	}
}

func (arbol abb[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIterador := new(iterador[K, V])
	nuevoIterador.elementos = TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	if arbol.raiz != nil {
		nuevoIterador.elementos.Apilar(*arbol.raiz)
		apilarHijosIzquierdos(nuevoIterador.elementos, arbol.raiz)
	}
	return nuevoIterador
}
func (iter iterador[K, V]) VerActual() (K, V) {
	if iter.elementos.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	return iter.elementos.VerTope().clave, iter.elementos.VerTope().dato
}

func (iter iterador[K, V]) HaySiguiente() bool {
	return !iter.elementos.EstaVacia()
}

func (iter *iterador[K, V]) Siguiente() K {
	if !iter.HaySiguiente() || iter.elementos.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	aux := iter.elementos.Desapilar()
	if aux.derecho != nil {
		iter.elementos.Apilar(*aux.derecho)
		apilarHijosIzquierdos(iter.elementos, aux.derecho)
	}
	return aux.clave
}

//******************************************************************************************
//* PRIMITIVA DEL ITERADOR EXTERNO IN-ORDER POR RANGO
//******************************************************************************************

func ApiladorSerialRec[K comparable, V any](cmp func(K, K) int, desde *K, elementos TDAPila.Pila[nodoAbb[K, V]], actual *nodoAbb[K, V]) {
	if actual != nil && cmp(actual.clave, *desde) == 0 {
		elementos.Apilar(*actual)
		return
	}
	if actual != nil && cmp(actual.clave, *desde) > 0 {
		elementos.Apilar(*actual)
		for actual.izquierdo != nil && cmp(actual.izquierdo.clave, *desde) >= 0 {
			actual = actual.izquierdo
			elementos.Apilar(*actual)
		}
		if actual.derecho == nil && actual.izquierdo == nil {
			return
		} else if actual.izquierdo != nil {
			ApiladorSerialRec(cmp, desde, elementos, actual.izquierdo)
		} else {
			ApiladorSerialRec(cmp, desde, elementos, actual.derecho)
		}
	}
	if actual != nil && cmp(actual.clave, *desde) < 0 {
		if actual.derecho == nil {
			return
		}

		for actual.derecho != nil {
			actual = actual.derecho

			if actual != nil && cmp(actual.clave, *desde) >= 0 {
				elementos.Apilar(*actual)
				break
			}
		}

		if actual == nil || cmp(actual.clave, *desde) < 0 || actual.izquierdo == nil {
			return
		}
		for actual.izquierdo != nil && cmp(actual.izquierdo.clave, *desde) >= 0 {
			actual = actual.izquierdo
			elementos.Apilar(*actual)
		}
		if actual != nil && cmp(actual.clave, *desde) == 0 {
			return
		}
		if actual.derecho == nil && actual.izquierdo == nil {
			return
		}
		if actual.derecho != nil {
			ApiladorSerialRec(cmp, desde, elementos, actual.derecho)
		} else {
			ApiladorSerialRec(cmp, desde, elementos, actual.izquierdo)
		}
		return
	}
}

func (arbol abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	nuevoIterador := new(iteradorRango[K, V])
	nuevoIterador.elementos = TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	nuevoIterador.desde = desde
	nuevoIterador.hasta = hasta
	nuevoIterador.cmp = arbol.cmp

	if arbol.raiz == nil {
		return nuevoIterador
	}
	if arbol.raiz != nil && desde == nil {
		nuevoIterador.elementos.Apilar(*arbol.raiz)
		apilarHijosIzquierdos(nuevoIterador.elementos, arbol.raiz)
		return nuevoIterador
	}

	ApiladorSerialRec(nuevoIterador.cmp, desde, nuevoIterador.elementos, arbol.raiz)
	return nuevoIterador
}

func (iter iteradorRango[K, V]) VerActual() (K, V) {
	if iter.elementos.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	if iter.hasta != nil && iter.cmp(iter.elementos.VerTope().clave, *iter.hasta) > 0 {
		panic("El iterador termino de iterar")
	}
	return iter.elementos.VerTope().clave, iter.elementos.VerTope().dato
}

func (iter iteradorRango[K, V]) HaySiguiente() bool {
	if iter.elementos.EstaVacia() {
		return false
	}
	if iter.hasta != nil && iter.cmp(iter.elementos.VerTope().clave, *iter.hasta) > 0 {
		return false
	}
	return true
}

func (iter *iteradorRango[K, V]) Siguiente() K {
	if iter.elementos.EstaVacia() || !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	aux := iter.elementos.Desapilar()
	if aux.derecho != nil && iter.desde == nil {
		iter.elementos.Apilar(*aux.derecho)
		apilarHijosIzquierdos(iter.elementos, aux.derecho)
	} else if aux.derecho != nil && iter.desde != nil {
		iter.elementos.Apilar(*aux.derecho)
		actual := aux.derecho
		for actual.izquierdo != nil && iter.cmp(actual.izquierdo.clave, *iter.desde) >= 0 {
			actual = actual.izquierdo
			iter.elementos.Apilar(*actual)
		}
	}
	return aux.clave
}
