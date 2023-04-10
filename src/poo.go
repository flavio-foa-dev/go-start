package main

import "fmt"

type People struct {
	firstName string
	lastName  string
	agencia   int32
	conta     int64
}

type ContaCorrente struct {
	nome          string
	numeroAgencia int32
	numeroConta   int32
	saldo         float64
}

func main() {
	contaMaria := ContaCorrente{}
	contaMaria.nome = "Maria"
	contaMaria.numeroAgencia = 10234
	contaMaria.numeroConta = 2876
	contaMaria.saldo = 500

	contaJoao := ContaCorrente{}
	contaJoao.nome = "Joao"
	contaJoao.numeroAgencia = 10234
	contaJoao.numeroConta = 2666
	contaJoao.saldo = 500

	fmt.Println(contaMaria)
	var nome string = "erick"
	fmt.Println(nome)
	fmt.Println(&nome)

	contaMaria.transFerir(100, &contaJoao)
	contaJoao.transFerir(50, &contaMaria)

}

func (c *ContaCorrente) sacar(valor float64) {
	consultaSaldo := valor <= c.saldo && valor > 0

	if consultaSaldo {
		c.saldo -= valor
		fmt.Println("Saque realizado com sucesso")
		fmt.Println("novo saldo conta1", c.saldo)

	} else {
		fmt.Println("Saldo insuficiente")
	}
}

func (c *ContaCorrente) depositar(valor float64) {
	saldoIsValid := valor > 0
	if saldoIsValid {
		c.saldo += valor
		fmt.Println("O Novo saldo apos deposito ", c.saldo)
	} else {
		fmt.Println("O valor passado para deposito invalido")
	}
}

func (c *ContaCorrente) transFerir(valor float64, conta *ContaCorrente) {
	saldoIsValid := valor > 0
	if saldoIsValid {
		consultaSaldo := valor <= c.saldo
		if consultaSaldo {
			c.saldo -= valor
			conta.depositar(valor)
			fmt.Println("O Novo saldo conta1 ", c.saldo)
			fmt.Println("novo saldo conta2", conta.saldo)

		} else {
			fmt.Println("Saldo insuficiente para transferencia")
		}

	} else {
		fmt.Println("O valor passado nao e valido")
	}
}

func teste() {
	flavio := People{"Flavio", "Andrade", 1000, 2222}
	var erick *People
	erick = new(People)
	erick.firstName = "Erick"
	erick.lastName = "Andrade"

	fmt.Println(flavio)
	fmt.Println(&erick)
	fmt.Println(erick)
	fmt.Println(*erick)

	var a *string = nil
	fmt.Println(a)

	conta1 := People{"Flavio", "Andrade", 1000, 2222}
	conta2 := People{"Flavio", "Andrade", 1000, 2222}

	fmt.Println(conta1 == conta2)

	var conta3 *People
	conta3 = new(People)
	conta3.firstName = "John"
	conta3.lastName = "Doe"

	var conta4 *People
	conta4 = new(People)
	conta4.firstName = "John"
	conta4.lastName = "Doe"

	fmt.Println(conta3 == conta4)
	fmt.Println(*conta3 == *conta4)
}
