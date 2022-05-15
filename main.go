package main

import (
	"fmt"

	"github.com/Alexsander-Espindola/API-with-golang/aprendendo"
)

// Criando métodos
type Pessoa struct {
	Nome string
}

func (pessoa Pessoa) comer() {
	fmt.Println(pessoa.Nome, " está comendo")
}

func main() {
	// Iniciando uma struct
	pessoa := Pessoa{
		Nome: "Alexsander",
	}
	pessoa.comer()

	// Structs com ponteiros
	aprendendo.PessoaPonteiro()

	// Structs, Composição e Json
	aprendendo.ClienteStruct()

	// Declara variável e imprime ela
	name := "Alexsander"
	fmt.Printf("Hello %v, seja bem vindo! \n", name)

	// Imprime o tipo da variável
	soma := testeFuncSoma(10, 20)
	fmt.Printf("%T \n", soma)

	// Importando uma função
	retorno := aprendendo.FuncExport()
	fmt.Println(retorno)

	// Funções de tratamento e lançamento de erro
	aprendendo.TratandoErro()
	_, err := aprendendo.LancandoErro(0)
	if err != nil {
		fmt.Println(err)
	}

	// Fução variádica
	fmt.Println(aprendendo.FuncaoVariadica(1, 4, 7, 10, 13))

	// Função anônima
	retornoAnonimo := func(x int, y int) func() int {
		retorno := x * y
		return func() int {
			return retorno / x
		}
	}
	fmt.Println(retornoAnonimo(10, 20)())

	// Ponteiros
	aprendendo.Ponteiros()
}

func testeFuncSoma(a int, b int) int {
	return a + b
}
