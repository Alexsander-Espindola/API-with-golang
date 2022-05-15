package aprendendo

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Função exportadvel começa com letra maiúscula no Go
// FuncExport retorna uma string
func FuncExport() string {
	return "Exporte-me"
}

// TratandoErro vai fazer um tratamento de erro
func TratandoErro() {
	_, err := http.Get("http://google.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}
}

// LancandoErro vai fazer um lançamento de erro se x == 0
func LancandoErro(x int) (int, error) {
	if x == 0 {
		return 0, errors.New("Lançando um novo erro")
	}
	return x + 10, nil
}

// RetornoDiferente vai retornar um inteiro
func RetornoDiferente(x int, y int) (soma int) {
	soma = x + y
	return
}

// FuncaoVariadica retorna a soma de todos valores inteiros que for passado de parâmetro
func FuncaoVariadica(x ...int) (somaTudo int) {
	for _, value := range x {
		somaTudo += value
	}
	return
}

// Ponteiros
func Ponteiros() {
	a := 20
	fmt.Println("Variável: ", a, "Endereço de memória da variável: ", &a)

	var b *int = &a
	fmt.Println("Variável: ", *b, "Endereço de memória da variável: ", b)

	*b = 50
	fmt.Println(*b, a)

	c := &a
	*c = 30
	fmt.Println(a, *b, *c)

	ponteiro(&a)
	fmt.Println(a, *b, *c)

}

func ponteiro(pont *int) {
	*pont = 40
}
