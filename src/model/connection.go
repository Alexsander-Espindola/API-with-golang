package model

import (
	"context"

	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() *mongo.Client {
	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		utils.GetErro(err)
	}

	return db
}
