// caso o go não ache o packege execute esse comando no terminal "go mod init banco" || solução para o erro "main.go:4:2: package banco/contas is not in GOROOT"

package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.AlteraSaldo(valorBoleto, true)
}

type verificarConta interface {
	AlteraSaldo(valorDeAlteracao float64, tipoOperacao bool) (string, float64, bool)
}

func main() {
	contaDoPaulo := contas.ContaPoupanca{}
	contaDoPaulo.AlteraSaldo(100, false)
	PagarBoleto(&contaDoPaulo, 60)

	fmt.Println(contaDoPaulo.ObterSaldo())

	contaDaLua := contas.ContaCorrente{}
	contaDaLua.AlteraSaldo(500, false)
	PagarBoleto(&contaDaLua, 400)

	fmt.Println(contaDaLua.ObterSaldo())
}
