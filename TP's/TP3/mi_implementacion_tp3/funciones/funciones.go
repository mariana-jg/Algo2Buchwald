package funciones

import (
	TDACola "vamosmoshi/cola"
	TDAGrafo "vamosmoshi/grafo"
	TDAHash "vamosmoshi/hash"
	TDAHeap "vamosmoshi/heap"
	TDALista "vamosmoshi/lista"
)

type Direccion struct {
	Desde string
	Hasta string
}

type Aristas struct {
	Desde string
	Hasta string
	Peso  int
}

func comparacionAristas(a Aristas, b Aristas) int {
	if a.Peso < b.Peso {
		return 1
	}
	if a.Peso > b.Peso {
		return -1
	}
	return 0

}

func cargarGrafo(grafo TDAGrafo.Grafo[string]) TDAGrafo.Grafo[string] {
	nuevoGrafo := TDAGrafo.CrearGrafo[string](false)
	iter := grafo.IterarVertices()
	for iter.HaySiguiente() {
		nuevoGrafo.AgregarVertice(iter.VerActual())
		iter.Siguiente()
	}
	return nuevoGrafo
}

func CalcularGrado(grafo TDAGrafo.Grafo[string]) TDAHash.Diccionario[string, int] {
	grados := TDAHash.CrearHash[string, int]()
	iter := grafo.IterarVertices()
	for iter.HaySiguiente() {
		adyacentes := grafo.ObtenerAdyacentes(iter.VerActual())
		grados.Guardar(iter.VerActual(), adyacentes.Largo())
		iter.Siguiente()
	}
	return grados
}

func CalcularGradoEntrada(grafo TDAGrafo.Grafo[string]) TDAHash.Diccionario[string, int] {
	grados := TDAHash.CrearHash[string, int]()
	iter1 := grafo.IterarVertices()
	for iter1.HaySiguiente() {
		v := iter1.VerActual()
		grados.Guardar(v, 0)
		iter1.Siguiente()
	}
	iter2 := grafo.IterarVertices()
	for iter2.HaySiguiente() {
		v := iter2.VerActual()
		adyacentes := grafo.ObtenerAdyacentes(v)
		iterAdyacentes := adyacentes.Iterador()
		for iterAdyacentes.HaySiguiente() {
			w := iterAdyacentes.VerActual()
			if !grados.Pertenece(w) {
				grados.Guardar(w, 1)
			} else {
				gradoW := grados.Obtener(w)
				grados.Guardar(w, gradoW+1)
			}
			iterAdyacentes.Siguiente()
		}
		iter2.Siguiente()
	}
	return grados
}

func ObtenerTodasLasAristas(grafo TDAGrafo.Grafo[string]) TDAHash.Diccionario[Direccion, bool] {
	aristas := TDAHash.CrearHash[Direccion, bool]()
	iter := grafo.IterarVertices()
	for iter.HaySiguiente() {
		actual := iter.VerActual()
		adyacentes := grafo.ObtenerAdyacentes(actual)
		iterAdyacentes := adyacentes.Iterador()

		for iterAdyacentes.HaySiguiente() {
			aristas.Guardar(Direccion{actual, iterAdyacentes.VerActual()}, true)
			iterAdyacentes.Siguiente()
		}
		iter.Siguiente()
	}
	return aristas
}

func ExisteCiclo(grados TDAHash.Diccionario[string, int]) bool {
	iter := grados.Iterador()
	for iter.HaySiguiente() {
		_, V := iter.VerActual()
		if V%2 != 0 {
			return false
		}
		iter.Siguiente()
	}
	return true
}

func cargarNuevoCamino(camino TDALista.Lista[string], caminoSec TDALista.Lista[string], U string) {
	iter := camino.Iterador()
	contador := 0
	for iter.HaySiguiente() {
		if iter.VerActual() == U && contador == 0 {
			iterSec := caminoSec.Iterador()
			for iterSec.HaySiguiente() {
				iter.Insertar(iterSec.VerActual())
				iterSec.Siguiente()
			}
			contador++
		}
		iter.Siguiente()
	}
}

func cicloC(grafo TDAGrafo.Grafo[string], aristas TDAHash.Diccionario[Direccion, bool], camino TDALista.Lista[string]) {
	iter := camino.Iterador()
	for iter.HaySiguiente() {
		adyacentes := grafo.ObtenerAdyacentes(iter.VerActual())
		iterAdyacentes := adyacentes.Iterador()
		for iterAdyacentes.HaySiguiente() {
			if aristas.Pertenece(Direccion{iter.VerActual(), iterAdyacentes.VerActual()}) {
				caminoSec := TDALista.CrearListaEnlazada[string]()
				Hierholzer(iter.VerActual(), grafo, aristas, caminoSec, iter.VerActual())
				cargarNuevoCamino(camino, caminoSec, iter.VerActual())
			}
			iterAdyacentes.Siguiente()
		}
		iter.Siguiente()
	}
}

func Hierholzer(verticeInicial string, grafo TDAGrafo.Grafo[string], aristas TDAHash.Diccionario[Direccion, bool],
	camino TDALista.Lista[string], salida string) {
	adyacentes := grafo.ObtenerAdyacentes(verticeInicial)
	iter := adyacentes.Iterador()
	for iter.HaySiguiente() {
		if aristas.Pertenece(Direccion{verticeInicial, iter.VerActual()}) {
			aristas.Borrar(Direccion{verticeInicial, iter.VerActual()})
			aristas.Borrar(Direccion{iter.VerActual(), verticeInicial})
			camino.InsertarUltimo(iter.VerActual())
			if iter.VerActual() == salida {
				cicloC(grafo, aristas, camino)
				return
			} else {
				Hierholzer(iter.VerActual(), grafo, aristas, camino, salida)
			}
		}
		iter.Siguiente()
	}
}

func Dfs(grafo TDAGrafo.Grafo[string], inicio string, visitados TDAHash.Diccionario[string, int]) {
	visitados.Guardar(inicio, 0)
	adyacentes := grafo.ObtenerAdyacentes(inicio)
	iterA := adyacentes.Iterador()
	for iterA.HaySiguiente() {
		if !visitados.Pertenece(iterA.VerActual()) {
			Dfs(grafo, iterA.VerActual(), visitados)
		}
		iterA.Siguiente()
	}
}

func EsConexo(grafo TDAGrafo.Grafo[string]) bool {
	vertices := grafo.ObtenerVertices()
	visitados := TDAHash.CrearHash[string, int]()
	componentes := 0
	iterV := vertices.Iterador()
	for iterV.HaySiguiente() {
		if !visitados.Pertenece(iterV.VerActual()) {
			componentes++
			v := iterV.VerActual()
			Dfs(grafo, v, visitados)
		}
		iterV.Siguiente()
	}
	return componentes <= 1
}

func OrdenTopologico(grafo TDAGrafo.Grafo[string]) TDALista.Lista[string] {
	grados := CalcularGradoEntrada(grafo)
	cola := TDACola.CrearColaEnlazada[string]()
	iter := grafo.IterarVertices()
	orden := TDALista.CrearListaEnlazada[string]()
	for iter.HaySiguiente() {
		vertice := iter.VerActual()
		if grados.Obtener(vertice) == 0 {
			cola.Encolar(vertice)
		}
		iter.Siguiente()
	}
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		orden.InsertarUltimo(v)
		adyacentes := grafo.ObtenerAdyacentes(v)
		iterA := adyacentes.Iterador()
		for iterA.HaySiguiente() {
			w := iterA.VerActual()
			gradoW := grados.Obtener(w)
			grados.Guardar(w, gradoW-1)
			if grados.Obtener(w) == 0 {
				cola.Encolar(w)
			}
			iterA.Siguiente()
		}
	}
	return orden
}

func CalcularPesoTotal(grafo TDAGrafo.Grafo[string]) int {
	peso := 0
	iter := grafo.IterarVertices()
	for iter.HaySiguiente() {
		adyacentes := grafo.ObtenerAdyacentes(iter.VerActual())
		iterA := adyacentes.Iterador()
		for iterA.HaySiguiente() {
			peso += grafo.PesoUnion(iter.VerActual(), iterA.VerActual())
			iterA.Siguiente()
		}
		iter.Siguiente()
	}
	return peso / 2
}

func TendidoMinimo(grafo TDAGrafo.Grafo[string]) TDAGrafo.Grafo[string] {
	vertice := grafo.ObtenerVerticeRandom()
	visitados := TDAHash.CrearHash[string, bool]()
	visitados.Guardar(vertice, true)
	heap := TDAHeap.CrearHeap(comparacionAristas)
	grafoAux := cargarGrafo(grafo)
	adyacentes := grafo.ObtenerAdyacentes(vertice)
	iter := adyacentes.Iterador()
	for iter.HaySiguiente() {
		w := iter.VerActual()
		arista := Aristas{vertice, w, grafo.PesoUnion(vertice, w)}
		heap.Encolar(arista)
		iter.Siguiente()
	}
	for !heap.EstaVacia() {
		arista := heap.Desencolar()
		if !visitados.Pertenece(arista.Hasta) || !visitados.Pertenece(arista.Desde) {
			grafoAux.AgregarArista(arista.Desde, arista.Hasta, arista.Peso)
			visitados.Guardar(arista.Hasta, true)
		}
		adyacente := grafo.ObtenerAdyacentes(arista.Hasta)
		iter := adyacente.Iterador()
		for iter.HaySiguiente() {
			if !visitados.Pertenece(iter.VerActual()) {
				q := Aristas{arista.Hasta, iter.VerActual(), grafo.PesoUnion(arista.Hasta, iter.VerActual())}
				heap.Encolar(q)
			}
			iter.Siguiente()
		}

	}
	return grafoAux
}

func CaminoMinimo(desde string, hasta string, grafo TDAGrafo.Grafo[string]) (TDALista.Lista[string], int) {
	distancia := TDAHash.CrearHash[string, int]()
	padre := TDAHash.CrearHash[string, string]()
	cola := TDACola.CrearColaEnlazada[string]()
	distancia.Guardar(desde, 0)
	cola.Encolar(desde)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		adyacentes := grafo.ObtenerAdyacentes(v)
		iter := adyacentes.Iterador()
		for iter.HaySiguiente() {
			w := iter.VerActual()
			if (!distancia.Pertenece(w)) ||
				(distancia.Pertenece(w) && distancia.Obtener(v)+grafo.PesoUnion(v, w) < distancia.Obtener(w)) {
				distancia.Guardar(w, distancia.Obtener(v)+grafo.PesoUnion(v, w))
				padre.Guardar(w, v)
				cola.Encolar(w)
			}
			iter.Siguiente()
		}
	}
	camino := TDALista.CrearListaEnlazada[string]()
	actual := hasta
	for padre.Pertenece(actual) {
		camino.InsertarPrimero(actual)
		actual = padre.Obtener(actual)
	}
	camino.InsertarPrimero(actual)
	return camino, distancia.Obtener(hasta)
}

func RecorridoTodasLasAristas(desde string, grafo TDAGrafo.Grafo[string]) TDALista.Lista[string] {
	grados := CalcularGrado(grafo)
	if !ExisteCiclo(grados) {
		return nil
	}
	if !EsConexo(grafo) {
		return nil
	}
	camino := TDALista.CrearListaEnlazada[string]()
	camino.InsertarPrimero(desde)
	aristas := ObtenerTodasLasAristas(grafo)
	Hierholzer(desde, grafo, aristas, camino, desde)
	return camino
}
