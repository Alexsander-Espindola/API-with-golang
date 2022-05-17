package aprendendo

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome string
}
type PessoaCpf string

type Cliente struct {
	Nome  string
	Cpf   string
	Email string
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func (pessoa Pessoa) comer() {
	pessoa.Nome = "Outro nome"
	fmt.Println(pessoa.Nome, "está comendo")
}

// Ponteiros com structs
func (pessoa *Pessoa) andar() {
	pessoa.Nome = "Alexs"
	fmt.Println(pessoa.Nome, "está andando")
}

func PessoaPonteiro() {
	pessoa := Pessoa{
		Nome: "Alexsander",
	}
	pessoa.comer()
	fmt.Println(pessoa.Nome)

	pessoa.andar()
	fmt.Println(pessoa.Nome)
}

func ClienteStruct() {
	cliente := Cliente{
		Nome:  "Alexsander",
		Cpf:   "000.000.000.00",
		Email: "Alexsander@gmail.com",
	}
	// ou cliente := Cliente{"Alexs", "000.000.000.00", "Alexs@gmail.com"}

	cliente2 := ClienteInternacional{
		Cliente: Cliente{"Alexs", "000.000.000.00", "Alexs@gmail.com"},
		Pais:    "Canada",
	}

	fmt.Println(cliente, cliente2)
	fmt.Printf("Nome: %s, Email: %s, Pais: %s\n", cliente2.Nome, cliente2.Email, cliente2.Pais)

	// json.Marshal retorna o valor em um slice de bytes
	clienteJson, err := json.Marshal(cliente2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(clienteJson)
	fmt.Println(string(clienteJson))

	jsonCliente := `{"Nome": "A", "Email": "A@A.com", "Cpf": "000.000.000.00", "pais": "EUA"}`
	cliente3 := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente), &cliente3)
	fmt.Println(cliente3)
}
