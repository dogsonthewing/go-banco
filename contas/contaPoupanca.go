package contas

import (
	"banco/clientes"
	"banco/globalfuncs"
)

// A primeira letra deve ser maiúscula para variável globais e minúscula para locais
type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

// tipo FALSE = depósito  ||  tipo TRUE = saque
func (c *ContaPoupanca) AlteraSaldo(valorDeAlteracao float64, tipoOperacao bool) (string, float64, bool) {
	objvalor := globalfuncs.ValorParaPositivar{valorDeAlteracao}
	valorDeAlteracao = objvalor.TornaPositivo(valorDeAlteracao)
	if tipoOperacao && valorDeAlteracao <= c.saldo {
		c.saldo -= valorDeAlteracao
		return "Saque realizado com sucesso", c.saldo, true
	} else if !tipoOperacao {
		c.saldo += valorDeAlteracao
		return "Depósito realizado com sucesso", c.saldo, true
	} else {
		return "Valor inválido", valorDeAlteracao, false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
