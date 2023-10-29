package votos

import (
	"rerepolez/errores"
	TDAPila "rerepolez/pila"
)

type votanteImplementacion struct {
	yaVoto     bool
	dni        int
	votoActual Voto
	voto       TDAPila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.voto = TDAPila.CrearPilaDinamica[Voto]()
	votante.yaVoto = false
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.yaVoto {
		return &errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if alternativa == LISTA_IMPUGNA {
		votante.votoActual.Impugnado = true
	} else {
		votante.votoActual.VotoPorTipo[tipo] = alternativa
	}
	votante.voto.Apilar(votante.votoActual)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	var votoEnBlanco Voto
	if votante.voto.EstaVacia() {
		return &errores.ErrorNoHayVotosAnteriores{}
	} else if votante.yaVoto {
		return &errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	votante.voto.Desapilar()
	if !votante.voto.EstaVacia() {
		votante.votoActual = votante.voto.VerTope()
	} else if votante.voto.EstaVacia() {
		votante.votoActual = votoEnBlanco
	}
	return nil
}
func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.voto.EstaVacia() {
		votante.voto.Apilar(votante.votoActual)
	}
	if votante.yaVoto {
		return votante.voto.VerTope(), &errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	votante.yaVoto = true
	return votante.voto.VerTope(), nil
}
