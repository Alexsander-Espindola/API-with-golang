package model

import (
	"context"
	"log"

	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Token    string             `bson:"token"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) PrepareUser() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	utils.GetErro(err)

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()
	if err != nil {
		log.Fatalf("Erro na validação do usuario: %v", err)
	}

	return nil
}

func (user *User) validate() error {
	_, err := govalidator.ValidateStruct(user)
	return err
}

func InsertUser(user User) string {
	db := GetConnection()
	err := user.PrepareUser()
	utils.GetErro(err)

	collection := db.Database("user").Collection("userData")

	_, err = collection.InsertOne(context.Background(), user)
	utils.GetErro(err)

	return user.Token
}
