package model

import (
	"context"
	"fmt"

	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
)

func PostUser(user User) {
	db := GetConnection()
	collection := db.Database("user").Collection("userData")

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		utils.GetErro(err)
	}

	fmt.Println(result)

	fmt.Println("Teste")
	fmt.Println(user)
}
