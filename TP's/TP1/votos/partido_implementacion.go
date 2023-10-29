package votos

import (
	"fmt"
)

type informacionCandidato struct {
	candidato string
	cantidad  int
}

type partidoImplementacion struct {
	nombrePartido string
	candidatos    [CANT_VOTACION]informacionCandidato
}

type partidoEnBlanco struct {
	nombrePartidoEnBlanco string
	candidatos            [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombrePartido = nombre
	for i := 0; i < int(CANT_VOTACION); i++ {
		partido.candidatos[i].candidato = candidatos[i]
	}
	return partido
}

func CrearVotosEnBlanco() Partido {
	partido := new(partidoEnBlanco)
	partido.nombrePartidoEnBlanco = "En blanco"
	return partido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.candidatos[tipo].cantidad++
}

func (partido partidoImplementacion) obtenerDatos(tipo TipoVoto) (string, int) {
	return partido.candidatos[tipo].candidato, partido.candidatos[tipo].cantidad

}

func distincionPlural(cantidad int) string {
	if cantidad == 1 {
		return ""
	} else {
		return "s"
	}
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var linea string
	candidato, cantidad := partido.obtenerDatos(tipo)
	linea = fmt.Sprintf("%s - %s: %d voto"+distincionPlural(cantidad), partido.nombrePartido, candidato, cantidad)
	return linea
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.candidatos[tipo]++
}

func (blanco partidoEnBlanco) obtenerDatosEnBlanco(tipo TipoVoto) int {
	return blanco.candidatos[tipo]
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	cantidad := blanco.obtenerDatosEnBlanco(tipo)
	var linea string
	linea = fmt.Sprintf("%s: %d voto"+distincionPlural(cantidad), "Votos en Blanco", cantidad)
	return linea
}
