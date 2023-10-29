package main

import (
	"bufio"
	"fmt"
	"os"
	TDACola "rerepolez/cola"
	errores "rerepolez/errores"
	votos "rerepolez/votos"
	"strconv"
	"strings"
)

const (
	INGRESAR     = "ingresar"
	VOTAR        = "votar"
	DESHACER     = "deshacer"
	FIN_VOTAR    = "fin-votar"
	PRESIDENTE   = "Presidente"
	GOBERNADOR   = "Gobernador"
	INTENDENTE   = "Intendente"
	CONFIRMACION = "OK"
)

func merge(a []votos.Votante, b []votos.Votante) []votos.Votante {
	final := []votos.Votante{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].LeerDNI() < b[j].LeerDNI() {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(vec []votos.Votante) []votos.Votante {
	if len(vec) < 2 {
		return vec
	}
	izq := mergeSort(vec[:len(vec)/2])
	der := mergeSort(vec[len(vec)/2:])
	return merge(izq, der)
}

func cargarPadron(archivoPadrones *os.File) []votos.Votante {
	var vectorVotantes []votos.Votante
	scanner := bufio.NewScanner(archivoPadrones)
	for scanner.Scan() {
		linea := scanner.Text()
		numeroPadron, _ := strconv.Atoi(linea)
		votante := votos.CrearVotante(numeroPadron)
		vectorVotantes = append(vectorVotantes, votante)
	}
	aux := mergeSort(vectorVotantes)
	return aux
}

func cargarPartidos(partidos *os.File) []votos.Partido {
	var vec [](votos.Partido)
	partidoBlanco := votos.CrearVotosEnBlanco()
	vec = append(vec, partidoBlanco)
	scanner := bufio.NewScanner(partidos)
	for scanner.Scan() {
		linea := scanner.Text()
		comandos := strings.Split(linea, ",")
		vectorCandidatos := [votos.CANT_VOTACION]string{comandos[1], comandos[2], comandos[3]}
		par := votos.CrearPartido(comandos[0], vectorCandidatos)
		vec = append(vec, par)
	}
	return vec
}

func esDelPadron(vec []votos.Votante, buscado int, inicio int, fin int) int {
	if inicio > fin || len(vec) == 0 {
		return -1
	}
	medio := (inicio + fin) / 2
	if vec[medio].LeerDNI() == buscado {
		return medio
	} else if vec[medio].LeerDNI() < buscado {
		return esDelPadron(vec, buscado, medio+1, fin)
	}
	return esDelPadron(vec, buscado, inicio, medio-1)

}

func esDNIValido(dni string) bool {
	intDni, err := strconv.Atoi(dni)
	return intDni > 0 && err == nil
}

func TipoDeVoto(tipo string, etiquetasCandidatos []string) votos.TipoVoto {
	var i int
	for i < len(etiquetasCandidatos) {
		if tipo == etiquetasCandidatos[i] {
			return votos.TipoVoto(i)
		}
		i++
	}
	return -1
}

func ingresarVotante(dni string, padron []votos.Votante, fila TDACola.Cola[votos.Votante]) error {
	if !esDNIValido(dni) {
		return &errores.DNIError{}
	}
	intDni, _ := strconv.Atoi(dni)
	perteneceAlPadron := esDelPadron(padron, intDni, 0, len(padron)-1)
	if perteneceAlPadron == -1 {
		return &errores.DNIFueraPadron{}
	} else {
		fila.Encolar(padron[perteneceAlPadron])
	}
	return nil
}

func votar(fila TDACola.Cola[votos.Votante], tipo string, lista string, partidos []votos.Partido, etiquetasCandidatos []string) {
	if fila.EstaVacia() {
		fmt.Fprintf(os.Stdout, "%s\n", errores.FilaVacia{}.Error())
		return
	} else if TipoDeVoto(tipo, etiquetasCandidatos) == -1 {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorTipoVoto{}.Error())
		return
	}
	alternativa, err := strconv.Atoi(lista)
	if alternativa < 0 || alternativa >= len(partidos) || err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorAlternativaInvalida{}.Error())
	} else {
		if fila.VerPrimero().Votar(TipoDeVoto(tipo, etiquetasCandidatos), alternativa) == nil {
			fmt.Fprintf(os.Stdout, "%s\n", CONFIRMACION)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", fila.VerPrimero().Votar(TipoDeVoto(tipo, etiquetasCandidatos), alternativa))
			fila.Desencolar()
		}
	}
}

func deshacer(fila TDACola.Cola[votos.Votante]) {
	if fila.EstaVacia() {
		fmt.Fprintf(os.Stdout, "%s\n", errores.FilaVacia{}.Error())
		return
	} else {
		err := fila.VerPrimero().Deshacer()
		if err == nil {
			fmt.Fprintf(os.Stdout, "%s\n", CONFIRMACION)
		} else {
			if (err.Error() == errores.ErrorNoHayVotosAnteriores{}.Error()) {
				fmt.Fprintf(os.Stdout, "%s\n", err)
			} else if (err.Error() == errores.ErrorVotanteFraudulento{Dni: fila.VerPrimero().LeerDNI()}.Error()) {
				fmt.Fprintf(os.Stdout, "%s\n", err)
				fila.Desencolar()
			}
		}
	}
}

func conversorTipoVoto(tipo int) votos.TipoVoto {
	return votos.TipoVoto(tipo)
}

func finVotar(fila TDACola.Cola[votos.Votante], partidos []votos.Partido, etiquetasCandidatos []string, impugnados *int) {
	if fila.EstaVacia() {
		fmt.Fprintf(os.Stdout, "%s\n", errores.FilaVacia{}.Error())
		return
	}
	votos, err := fila.VerPrimero().FinVoto()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		fila.Desencolar()
		return
	}
	if votos.Impugnado {
		*impugnados++
	} else {
		for i := 0; i < 3; i++ {
			partidos[votos.VotoPorTipo[i]].VotadoPara(conversorTipoVoto(i))
		}
	}
	fmt.Fprintf(os.Stdout, "%s\n", CONFIRMACION)
	fila.Desencolar()
}

func ingresarComando(dniVotante *int, comando string, tipo string,
	padron []votos.Votante, fila TDACola.Cola[votos.Votante],
	lista string, partidos []votos.Partido, etiquetasCandidatos []string, impugnados *int) {

	switch comando {
	case INGRESAR:
		if ingresarVotante(tipo, padron, fila) == nil {
			*dniVotante, _ = strconv.Atoi(tipo)
			fmt.Fprintf(os.Stdout, "%s\n", CONFIRMACION)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", ingresarVotante(tipo, padron, fila))
		}
		break
	case VOTAR:
		votar(fila, tipo, lista, partidos, etiquetasCandidatos)
		break

	case DESHACER:
		deshacer(fila)
		break

	case FIN_VOTAR:
		finVotar(fila, partidos, etiquetasCandidatos, impugnados)
		break
	}
}

func imprimirImpugandos(impugnados int) {
	var plural string
	if impugnados == 1 {
		plural = ""
	} else {
		plural = "s"
	}
	fmt.Fprintf(os.Stdout, "Votos Impugnados: %d voto"+plural+"\n", impugnados)
}

func imprimirResultados(partidos []votos.Partido, impugnados int, etiquetasCandidatos []string) {
	for j := conversorTipoVoto(0); j < votos.CANT_VOTACION; j++ {
		fmt.Fprintf(os.Stdout, "%s:\n", etiquetasCandidatos[j])
		for i := 0; i < len(partidos); i++ {
			fmt.Fprintf(os.Stdout, "%s\n", partidos[i].ObtenerResultado(j))
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
	imprimirImpugandos(impugnados)
}

func main() {
	archivos := os.Args[1:]
	if len(archivos) != 2 {
		fmt.Println(errores.ErrorParametros{}.Error())
		return
	}
	partidos, err := os.Open(archivos[0])
	padron, err2 := os.Open(archivos[1])
	if err != nil || err2 != nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		return
	}
	defer padron.Close()
	defer partidos.Close()

	padronVec := cargarPadron(padron)
	partidosVec := cargarPartidos(partidos)
	etiquetasCandidatos := []string{PRESIDENTE, GOBERNADOR, INTENDENTE}
	filaVotantes := TDACola.CrearColaEnlazada[votos.Votante]()
	var dniVotante int
	var impugnados int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := s.Text()
		comandos := strings.Split(entrada, " ")
		for len(comandos) < 3 {
			comandos = append(comandos, " ")
		}
		ingresarComando(&dniVotante, comandos[0], comandos[1], padronVec,
			filaVotantes, comandos[2], partidosVec, etiquetasCandidatos, &impugnados)
	}
	if !filaVotantes.EstaVacia() {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorCiudadanosSinVotar{}.Error())
	}
	imprimirResultados(partidosVec, impugnados, etiquetasCandidatos)
}
