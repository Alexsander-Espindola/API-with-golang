package model

import (
	"context"
	"fmt"

	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
)

// PostUser adiciona um usuario na Collection userData do banco de dados
func PostUser(user User) {
	db := GetConnection()
	collection := db.Database("user").Collection("userData")

	result, err := collection.InsertOne(context.Background(), user)
	utils.GetErro(err)
	fmt.Println(*result)
}

//
func FindOneUser() {

}
