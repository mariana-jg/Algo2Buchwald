package diccionario

import (
	"fmt"
)

type estadoClave int

const CAPACIDAD_INICIAL = 23
const FACTOR_DE_REDIMENSION = 2
const PORCENTAJE_DE_REDIMENSION_ALTO = 70
const MINIMO_FACTOR_DE_CARGA = 5

const (
	VACIO estadoClave = iota
	OCUPADO
	BORRADO
)

type hashCampo[K comparable, V any] struct {
	clave  K
	valor  V
	estado estadoClave
}

type hash[K comparable, V any] struct {
	capacidad int
	cantidad  int
	borrados  int
	tabla     []hashCampo[K, V]
}

type iterador[K comparable, V any] struct {
	hash     hash[K, V]
	actual   hashCampo[K, V]
	posicion int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

//******************************************************************************************
//* FUNCION DE HASHING
//******************************************************************************************

const (
	uint64Offset uint64 = 0xcbf29ce484222325
	uint64Prime  uint64 = 0x00000100000001b3
)

func hashing(data []byte, capacidad int) int {
	h := uint64Offset
	cap := uint64(capacidad)

	for _, b := range data {
		h ^= uint64(b)
		h *= uint64Prime
	}
	return (int(h % cap))
}

//******************************************************************************************
// Esta funcion de hashing fue extraida de:
// https://golangprojectstructure.com/hash-functions-go-code/#using-the-fnv-hash-function-in-your-own-go-code
//******************************************************************************************

/*
buscarIndice devuelve una posición válida para guardar.
Posicion valida implica

	>> Ocupada con misma clave
	>> Vacia

Posicion no valida implica

	>> Ocupada con distinta clave o borrada.
*/
func buscarIndice[K comparable, V any](hash hash[K, V], clave K) int {
	indice := hashing(convertirABytes(clave), hash.capacidad)
	for hash.tabla[indice].estado != VACIO && (hash.tabla[indice].estado == BORRADO || hash.tabla[indice].clave != clave) {
		indice = (indice + 1) % hash.capacidad
	}
	return indice
}

func llenarVacio[K comparable, V any](tabla []hashCampo[K, V], capacidad int) {
	for i := 0; i < capacidad; i++ {
		tabla[i].estado = VACIO
	}
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	nuevoHash := new(hash[K, V])
	nuevoHash.capacidad = CAPACIDAD_INICIAL
	nuevoHash.cantidad = 0
	tabla := make([]hashCampo[K, V], CAPACIDAD_INICIAL)
	nuevoHash.tabla = tabla
	llenarVacio(nuevoHash.tabla, nuevoHash.capacidad)
	return nuevoHash
}

func (hash *hash[K, V]) redimensionarHash(nuevoTam int) {
	anteriorCapacidad := hash.capacidad
	anteriorTabla := hash.tabla
	nuevaTabla := make([]hashCampo[K, V], nuevoTam)
	hash.capacidad = nuevoTam
	hash.tabla = nuevaTabla
	llenarVacio(hash.tabla, hash.capacidad)
	for i := 0; i < anteriorCapacidad; i++ {
		if anteriorTabla[i].estado == OCUPADO {
			indice := buscarIndice(*hash, anteriorTabla[i].clave)
			hash.tabla[indice] = anteriorTabla[i]
		}
	}
}

//******************************************************************************************
//* PRIMITIVAS DEL DICCIONARIO
//******************************************************************************************

func (hash hash[K, V]) Pertenece(clave K) bool {
	indice := buscarIndice(hash, clave)
	return hash.tabla[indice].estado == OCUPADO
}

func (hash hash[K, V]) Cantidad() int {
	return hash.cantidad
}

func verificarPertenencia[K comparable, V any](hash hash[K, V], clave K, indice int) {
	if hash.tabla[indice].estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
}

func (hash hash[K, V]) Obtener(clave K) V {
	indice := buscarIndice(hash, clave)
	verificarPertenencia(hash, clave, indice)
	return hash.tabla[indice].valor
}

func (hash *hash[K, V]) Guardar(clave K, valor V) {
	indice := buscarIndice(*hash, clave)
	if hash.tabla[indice].estado == VACIO {
		hash.cantidad++
	}
	hash.tabla[indice].clave = clave
	hash.tabla[indice].valor = valor
	hash.tabla[indice].estado = OCUPADO
	factorDeCarga := ((hash.cantidad + hash.borrados) * 100) / hash.capacidad
	if factorDeCarga >= PORCENTAJE_DE_REDIMENSION_ALTO {
		hash.redimensionarHash(FACTOR_DE_REDIMENSION * hash.capacidad)
	}
}

func (hash *hash[K, V]) Borrar(clave K) V {
	indice := buscarIndice(*hash, clave)
	verificarPertenencia(*hash, clave, indice)
	hash.tabla[indice].estado = BORRADO
	hash.cantidad--
	hash.borrados++
	factorDeCarga := ((hash.cantidad + hash.borrados) * 100) / hash.capacidad
	if factorDeCarga <= MINIMO_FACTOR_DE_CARGA && hash.capacidad >= FACTOR_DE_REDIMENSION*CAPACIDAD_INICIAL {
		hash.redimensionarHash(hash.capacidad / FACTOR_DE_REDIMENSION)
	}
	return hash.tabla[indice].valor
}

//******************************************************************************************
//* PRIMITIVA DEL ITERADOR INTERNO
//******************************************************************************************

func (hash hash[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	for i := 0; i < hash.capacidad; i++ {
		if hash.tabla[i].estado == OCUPADO {
			if !visitar(hash.tabla[i].clave, hash.tabla[i].valor) {
				return
			}
		}
	}
}

//******************************************************************************************
//* PRIMITIVAS DEL ITERADOR EXTERNO
//******************************************************************************************

func (hash hash[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIterador := new(iterador[K, V])
	nuevoIterador.hash = hash
	i := 0
	for hash.tabla[i].estado != OCUPADO && hash.Cantidad() != 0 {
		i++
	}
	nuevoIterador.actual = hash.tabla[i]
	nuevoIterador.posicion = i
	return nuevoIterador
}

func (iter iterador[K, V]) HaySiguiente() bool {
	return !(iter.posicion >= iter.hash.capacidad || iter.hash.cantidad == 0)
}

func (iter iterador[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.clave, iter.actual.valor
}

func (iter *iterador[K, V]) Siguiente() K {
	if !iter.HaySiguiente() || iter.hash.Cantidad() == 0 {
		panic("El iterador termino de iterar")
	}
	aux, _ := iter.VerActual()
	iter.posicion++
	for iter.posicion < iter.hash.capacidad && iter.hash.tabla[iter.posicion].estado != OCUPADO {
		iter.posicion++
	}
	if iter.posicion < iter.hash.capacidad {
		iter.actual = iter.hash.tabla[iter.posicion]
	}
	return aux
}
