package diccionario

import TDAPila "diccionario/pila"

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

type iterRangoABB[K comparable, V any] struct {
	desde     *K
	hasta     *K
	elementos TDAPila.Pila[nodoAbb[K, V]]
	cmp       func(K, K) int
}

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

func guardarRec[K comparable, V any](arbol *abb[K, V], nodoActual *nodoAbb[K, V], nodoNuevo *nodoAbb[K, V]) {
	if arbol.raiz == nil {
		arbol.raiz = nodoNuevo
		arbol.cantidad++
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
	nodo := buscarNodo(arbol.cmp, arbol.raiz, clave)
	if nodo != nil {
		nodo.dato = dato
		return
	}
	nodoNuevo := crearNodo(clave, dato)
	guardarRec(arbol, arbol.raiz, &nodoNuevo)
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

func borrarNodoSinHijos[K comparable, V any](arbol *abb[K, V], padre *nodoAbb[K, V], clave K) {
	if padre == nil {
		arbol.raiz = nil
	} else if padre.derecho != nil && arbol.cmp(padre.derecho.clave, clave) == 0 {
		padre.derecho = nil
	} else if padre.izquierdo != nil && arbol.cmp(padre.izquierdo.clave, clave) == 0 {
		padre.izquierdo = nil
	}
}

func asignarReemplazante[K comparable, V any](arbol *abb[K, V], padre *nodoAbb[K, V], nodoReemplazante *nodoAbb[K, V], clave K) {
	if padre == nil {
		arbol.raiz = nodoReemplazante
	} else if padre.derecho != nil && arbol.cmp(padre.derecho.clave, clave) == 0 {
		padre.derecho = nodoReemplazante
	} else if padre.izquierdo != nil && arbol.cmp(padre.izquierdo.clave, clave) == 0 {
		padre.izquierdo = nodoReemplazante
	}

}

func borrarNodoUnHijo[K comparable, V any](arbol *abb[K, V], padre *nodoAbb[K, V], nodo *nodoAbb[K, V]) {
	if nodo.izquierdo != nil {
		asignarReemplazante(arbol, padre, nodo.izquierdo, nodo.clave)
	} else if nodo.derecho != nil {
		asignarReemplazante(arbol, padre, nodo.derecho, nodo.clave)
	}

}

func borrarRec[K comparable, V any](arbol *abb[K, V], nodo *nodoAbb[K, V], padre *nodoAbb[K, V], hijos int) {
	if hijos == 0 {
		borrarNodoSinHijos(arbol, padre, nodo.clave)
	} else if hijos == 1 {
		borrarNodoUnHijo(arbol, padre, nodo)
	} else {
		auxiliar := buscarReemplazante(nodo)
		borrarRec(arbol, auxiliar, encontrarPadre(*arbol, arbol.raiz, auxiliar.clave), calcularCantHijos(auxiliar))
		nodo.dato = auxiliar.dato
		nodo.clave = auxiliar.clave
	}
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodo := buscarNodo(arbol.cmp, arbol.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	aux := nodo.dato
	hijos := calcularCantHijos(nodo)
	padre := encontrarPadre(*arbol, arbol.raiz, clave)
	borrarRec(arbol, nodo, padre, hijos)
	arbol.cantidad--
	return aux
}

//******************************************************************************************
//* ITERADOR INTERNO IN-ORDER
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

func iterarSobreRango[K comparable, V any](comp func(K, K) int, desde *K, hasta *K, visitar func(clave K, dato V) bool, nodo *nodoAbb[K, V]) {
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

func (arbol abb[K, V]) Iterar(visitar func(K, V) bool) {
	if arbol.raiz == nil {
		return
	}
	desde := buscarMenor(arbol.raiz)
	hasta := buscarMayor(arbol.raiz)
	iterarSobreRango(arbol.cmp, desde, hasta, visitar, arbol.raiz)
}

//******************************************************************************************
//* PRIMITIVA DEL ITERADOR INTERNO POR RANGO IN-ORDER
//******************************************************************************************

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
//* ITERADOR EXTERNO IN-ORDER
//******************************************************************************************

func apilarHijosIzquierdos[K comparable, V any](elementos TDAPila.Pila[nodoAbb[K, V]], actual *nodoAbb[K, V]) {
	for actual.izquierdo != nil {
		actual = actual.izquierdo
		elementos.Apilar(*actual)
	}
}

func (arbol abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

// ******************************************************************************************
// * ITERADOR EXTERNO IN-ORDER POR RANGO
// ******************************************************************************************
func revisarLadoIzquierdo[K comparable, V any](cmp func(K, K) int, desde *K,
	elementos TDAPila.Pila[nodoAbb[K, V]], actual *nodoAbb[K, V]) {
	elementos.Apilar(*actual)
	for actual.izquierdo != nil && cmp(actual.izquierdo.clave, *desde) >= 0 {
		actual = actual.izquierdo
		elementos.Apilar(*actual)
	}
	if actual.derecho == nil && actual.izquierdo == nil {
		return
	} else if actual.izquierdo != nil {
		iteradorRangoRec(cmp, desde, elementos, actual.izquierdo)
	} else {
		iteradorRangoRec(cmp, desde, elementos, actual.derecho)
	}
}

func revisarLadoDerecho[K comparable, V any](cmp func(K, K) int, desde *K,
	elementos TDAPila.Pila[nodoAbb[K, V]], actual *nodoAbb[K, V]) {
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
	if (actual != nil && cmp(actual.clave, *desde) == 0) || (actual.derecho == nil && actual.izquierdo == nil) {
		return
	}
	if actual.derecho != nil {
		iteradorRangoRec(cmp, desde, elementos, actual.derecho)
	} else {
		iteradorRangoRec(cmp, desde, elementos, actual.izquierdo)
	}

}

func iteradorRangoRec[K comparable, V any](cmp func(K, K) int, desde *K, elementos TDAPila.Pila[nodoAbb[K, V]], actual *nodoAbb[K, V]) {
	if actual != nil && cmp(actual.clave, *desde) == 0 {
		elementos.Apilar(*actual)
		return
	}
	if actual != nil && cmp(actual.clave, *desde) > 0 {
		revisarLadoIzquierdo(cmp, desde, elementos, actual)
	}
	if actual != nil && cmp(actual.clave, *desde) < 0 {
		revisarLadoDerecho(cmp, desde, elementos, actual)
	}
}

func (arbol abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	nuevoIterador := new(iterRangoABB[K, V])
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

	iteradorRangoRec(nuevoIterador.cmp, desde, nuevoIterador.elementos, arbol.raiz)
	return nuevoIterador
}

func (iter iterRangoABB[K, V]) VerActual() (K, V) {
	if iter.elementos.EstaVacia() || !iter.HaySiguiente() || (iter.hasta != nil &&
		iter.cmp(iter.elementos.VerTope().clave, *iter.hasta) > 0) {
		panic("El iterador termino de iterar")
	}
	return iter.elementos.VerTope().clave, iter.elementos.VerTope().dato
}

func (iter iterRangoABB[K, V]) HaySiguiente() bool {
	if iter.elementos.EstaVacia() || (iter.hasta != nil && iter.cmp(iter.elementos.VerTope().clave, *iter.hasta) > 0) {
		return false
	}
	return true
}

func (iter *iterRangoABB[K, V]) Siguiente() K {
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
