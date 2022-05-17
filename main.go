package main

import (
	"github.com/Alexsander-Espindola/API-with-golang/src/model"
)

func main() {
	user := model.User{
		Name:  "Alexs",
		Email: "Alexs@gmail.com",
	}

	model.PostUser(user)
}
