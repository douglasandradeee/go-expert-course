package ocp

import "reflect"

type FuncionarioCLT struct{}
type Estagiario struct{}

func FolhaPagamento[TipoFuncionario FuncionarioCLT | Estagiario](funcionario TipoFuncionario) float64 {
	if reflect.TypeOf(FuncionarioCLT{}) == reflect.TypeOf(funcionario) {
		//calcular salário + benefícios.
		return 0.0
	} else {
		//calcular bolsa auxílio do estagiário.
		return 0.0
	}

}
