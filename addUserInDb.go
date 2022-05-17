package main

import (
	"github.com/Alexsander-Espindola/API-with-golang/src/model"
)

func main() {
	user := model.User{
		Name:  "Alexs",
		Email: "Alexs@gmail.com",
		Endereco: model.Endereco{
			Estado: "Goiás",
			Cidade: "Aparecida de Goiânia",
		},
	}

	model.PostUser(user)
}
