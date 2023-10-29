package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACola "vamosmoshi/cola"
	Funciones "vamosmoshi/funciones"
	TDAGrafo "vamosmoshi/grafo"
	TDAHash "vamosmoshi/hash"
	TDALista "vamosmoshi/lista"
)

type coordenadas struct {
	x string
	y string
}

const (
	encabezadoKML = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	inicioKML     = `<kml xmlns="http://earth.google.com/kml/2.1">` + "\n\t" + `<Document>` + "\n"
	finKML        = "\t" + `</Document>` + "\n" + `</kml>`
)

func cargarCiudades(archivoCiudades *os.File) (TDAHash.Diccionario[string, coordenadas], TDAGrafo.Grafo[string], TDACola.Cola[string]) {
	ciudades := TDAHash.CrearHash[string, coordenadas]()
	conexiones := TDAGrafo.CrearGrafo[string](false)
	cola := TDACola.CrearColaEnlazada[string]()
	scanner := bufio.NewScanner(archivoCiudades)
	i := 0
	j := 0
	for scanner.Scan() {
		if i == 0 {
			linea := scanner.Text()
			cantidadCiudades, _ := strconv.Atoi(linea)
			j = cantidadCiudades
		} else if i <= j && i != 0 {
			linea := scanner.Text()
			comandos := strings.Split(linea, ",")
			ciudades.Guardar(comandos[0], coordenadas{comandos[1], comandos[2]})
			cola.Encolar(comandos[0])
			conexiones.AgregarVertice(comandos[0])
		} else {
			if i > j+1 {
				linea := scanner.Text()
				comandos := strings.Split(linea, ",")
				tiempo, _ := strconv.Atoi(comandos[2])
				conexiones.AgregarArista(comandos[0], comandos[1], tiempo)
			}
		}
		i++
	}
	return ciudades, conexiones, cola
}

func lineaRecorrido(lista TDALista.Lista[string]) string {
	iter := lista.Iterador()
	var camino []string
	for iter.HaySiguiente() {
		camino = append(camino, iter.VerActual())
		iter.Siguiente()
	}
	recorrido := strings.Join(camino, " -> ")
	return recorrido
}

func escribirInicioMapa(mapa *os.File, desde string, hasta string) {
	mapa.WriteString(encabezadoKML)
	mapa.WriteString(inicioKML)
	mapa.WriteString("\t\t" + `<name>` + "Camino desde " + desde + " hacia " + hasta + `</name>` + "\n\n")
}

func escribirCoordenadas(mapa *os.File, ciudad string, coordenada coordenadas) {
	mapa.WriteString("\t\t" + `<Placemark>` + "\n" + "\t\t\t" + `<name>` + ciudad + `</name>` + "\n" + "\t\t\t" + `<Point>` + "\n")
	mapa.WriteString("\t\t\t\t" + `<coordinates>` + coordenada.x + ", " + coordenada.y + `</coordinates>` + "\n")
	mapa.WriteString("\t\t\t" + `</Point>` + "\n" + "\t\t" + `</Placemark>` + "\n")
}

func escribirUnionesCiudades(mapa *os.File, lista TDALista.Lista[string], ciudades TDAHash.Diccionario[string, coordenadas]) {
	iter := lista.Iterador()
	actual := iter.Siguiente()
	for iter.HaySiguiente() {
		coordenadasActual := ciudades.Obtener(actual)
		coordenadasSig := ciudades.Obtener(iter.VerActual())
		mapa.WriteString("\t\t" + `<Placemark>` + "\n")
		mapa.WriteString("\t\t\t" + `<LineString>` + "\n")
		mapa.WriteString("\t\t\t\t" + `<coordinates>` + coordenadasActual.x + ", " + coordenadasActual.y + " " + coordenadasSig.x +
			", " + coordenadasSig.y + `</coordinates>` + "\n")
		mapa.WriteString("\t\t\t" + `</LineString>` + "\n" + "\t\t" + `</Placemark>` + "\n")
		actual = iter.VerActual()
		iter.Siguiente()
	}
}

func escribirMapaIrDesde(mapa *os.File, ciudades TDAHash.Diccionario[string, coordenadas], desde string, hasta string,
	lista TDALista.Lista[string]) {
	escribirInicioMapa(mapa, desde, hasta)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		ciudad := iter.VerActual()
		coordenada := ciudades.Obtener(ciudad)
		escribirCoordenadas(mapa, ciudad, coordenada)
		iter.Siguiente()
	}
	mapa.WriteString("\n")
	escribirUnionesCiudades(mapa, lista, ciudades)
	mapa.WriteString(finKML)
}

func escribirMapaViaje(resultado *os.File, desde string, ciudades TDAHash.Diccionario[string, coordenadas], lista TDALista.Lista[string]) {
	escribirInicioMapa(resultado, desde, desde)
	iter := ciudades.Iterador()
	for iter.HaySiguiente() {
		ciudad, coordenada := iter.VerActual()
		escribirCoordenadas(resultado, ciudad, coordenada)
		iter.Siguiente()
	}
	resultado.WriteString("\n")
	escribirUnionesCiudades(resultado, lista, ciudades)
	resultado.WriteString(finKML)
}

func tiempoDeViaje(grafo TDAGrafo.Grafo[string], lista TDALista.Lista[string]) int {
	tiempo := 0
	iter := lista.Iterador()
	actual := iter.Siguiente()
	for iter.HaySiguiente() {
		tiempo += grafo.PesoUnion(actual, iter.VerActual())
		actual = iter.VerActual()
		iter.Siguiente()
	}
	return tiempo
}

func cargarDestinos(destino *os.File, cantVertices string, cantAristas string, dicCiudades TDAHash.Diccionario[string, coordenadas],
	aristas TDAHash.Diccionario[Funciones.Direccion, int], cola TDACola.Cola[string]) {
	destino.WriteString(cantVertices + "\n")
	for !cola.EstaVacia() {
		ciudad := cola.Desencolar()
		coordenada := dicCiudades.Obtener(ciudad)
		destino.WriteString(ciudad + "," + coordenada.x + "," + coordenada.y + "\n")
	}
	destino.WriteString(cantAristas + "\n")
	iterAristas := aristas.Iterador()
	for iterAristas.HaySiguiente() {
		direccionActual, peso := iterAristas.VerActual()
		stringPeso := strconv.Itoa(peso)
		destino.WriteString(direccionActual.Desde + "," + direccionActual.Hasta + "," + stringPeso + "\n")
		iterAristas.Siguiente()
	}
}

func ir(conexiones TDAGrafo.Grafo[string], comandos2 []string, comandos []string, ciudades TDAHash.Diccionario[string, coordenadas]) {
	mapa, err := os.Create(comandos[2])
	if err != nil {
		fmt.Printf("No se pudo crear el archivo %s", comandos[2])
		return
	}
	defer mapa.Close()

	desde := strings.Join(comandos2[1:], " ")
	hasta := comandos[1]
	if !conexiones.ExisteID(desde) || !conexiones.ExisteID(hasta) {
		fmt.Println("No se encontro recorrido")
		return
	}
	lista, tiempo := Funciones.CaminoMinimo(desde, hasta, conexiones)
	if lista.EstaVacia() {
		fmt.Println("No se encontro recorrido")
		return
	}
	escribirMapaIrDesde(mapa, ciudades, desde, hasta, lista)
	recorrido := lineaRecorrido(lista)
	fmt.Printf("%s\nTiempo total: %d\n", recorrido, tiempo)
}

func itinerario(comandos2 []string, ciudades TDAHash.Diccionario[string, coordenadas]) {
	archivoOrden, err := os.Open(comandos2[1])
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	defer archivoOrden.Close()

	grafo := TDAGrafo.CrearGrafo[string](true)
	iter := ciudades.Iterador()
	for iter.HaySiguiente() {
		ciudad, _ := iter.VerActual()
		grafo.AgregarVertice(ciudad)
		iter.Siguiente()
	}
	scanner := bufio.NewScanner(archivoOrden)
	for scanner.Scan() {
		linea := scanner.Text()
		vertices := strings.Split(linea, ",")
		grafo.AgregarArista(vertices[0], vertices[1], 0)
	}
	lista := Funciones.OrdenTopologico(grafo)
	if lista.EstaVacia() {
		fmt.Println("No se encontro recorrido")
		return
	}
	linea := lineaRecorrido(lista)
	fmt.Printf("%s\n", linea)
}

func viaje(comandos []string, grafo TDAGrafo.Grafo[string], ciudades TDAHash.Diccionario[string, coordenadas], comand []string) {
	resultado, err := os.Create(comand[1])
	if err != nil {
		fmt.Printf("No se pudo crear el archivo %s", comandos[2])
		return
	}
	defer resultado.Close()

	desde := strings.Join(comandos[1:], " ")
	if !ciudades.Pertenece(desde) {
		fmt.Println("No se encontro recorrido")
		return
	}
	lista := Funciones.RecorridoTodasLasAristas(desde, grafo)
	if lista == nil {
		fmt.Println("No se encontro recorrido")
		return
	}
	escribirMapaViaje(resultado, desde, ciudades, lista)
	linea := lineaRecorrido(lista)
	fmt.Printf("%s\nTiempo total: %d\n", linea, tiempoDeViaje(grafo, lista))
}

func reducirCaminos(comandos []string, conexiones TDAGrafo.Grafo[string], dicCiudades TDAHash.Diccionario[string, coordenadas], cola TDACola.Cola[string]) {
	destino, err := os.Create(comandos[1])
	if err != nil {
		fmt.Printf("No se pudo crear el archivo %s", comandos[2])
		return
	}
	defer destino.Close()

	reducidos := Funciones.TendidoMinimo(conexiones)
	cantVertices := 0
	cantAristas := 0
	aristas := TDAHash.CrearHash[Funciones.Direccion, int]()
	iter := reducidos.IterarVertices()
	for iter.HaySiguiente() {
		cantVertices++
		adyacentes := reducidos.ObtenerAdyacentes(iter.VerActual())
		iterA := adyacentes.Iterador()
		origen := iter.VerActual()
		for iterA.HaySiguiente() {
			destino := iterA.VerActual()
			if !aristas.Pertenece(Funciones.Direccion{Desde: origen, Hasta: destino}) && !aristas.Pertenece(Funciones.Direccion{Desde: destino, Hasta: origen}) {
				aristas.Guardar(Funciones.Direccion{Desde: origen, Hasta: destino}, reducidos.PesoUnion(origen, destino))
				cantAristas++
			}
			iterA.Siguiente()
		}
		iter.Siguiente()
	}
	stringVertices := strconv.Itoa(cantVertices)
	stringAristas := strconv.Itoa(cantAristas)
	cargarDestinos(destino, stringVertices, stringAristas, dicCiudades, aristas, cola)
	fmt.Printf("Peso total: %d\n", Funciones.CalcularPesoTotal(reducidos))
}

func ingresarComando(comandos []string, ciudades TDAHash.Diccionario[string, coordenadas], conexiones TDAGrafo.Grafo[string],
	cola TDACola.Cola[string]) {
	comandos2 := strings.Split(comandos[0], " ")
	switch comandos2[0] {
	case "ir":
		ir(conexiones, comandos2, comandos, ciudades)
	case "itinerario":
		itinerario(comandos2, ciudades)
	case "viaje":
		viaje(comandos2, conexiones, ciudades, comandos)
	case "reducir_caminos":
		reducirCaminos(comandos2, conexiones, ciudades, cola)
	}
}

func main() {
	archivos := os.Args[1:]
	if len(archivos) != 1 {
		fmt.Println("Incorrecta cantidad de argumentos")
	}
	archivoCiudades, err := os.Open(archivos[0])
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	defer archivoCiudades.Close()

	dicCiudades, conexiones, cola := cargarCiudades(archivoCiudades)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := s.Text()
		comandos := strings.Split(entrada, ", ")
		ingresarComando(comandos, dicCiudades, conexiones, cola)
	}
}
