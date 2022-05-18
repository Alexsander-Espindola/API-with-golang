package model

import (
	"context"
	"fmt"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `bson:"name" valid:"notnull"`
	Email    string `bson:"email" valid:"notnull, email"`
	Password string `bson:"-" valid:"notnull"`
	Token    string `bson:"token" valid:"notnull, uuid"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) PrepareUser() error {
	if user.Password == "" {
		return fmt.Errorf("Senha em branco")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) validate() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(user User) (string, *mongo.InsertOneResult, error) {
	db := GetConnection()
	err := user.PrepareUser()
	if err != nil {
		return "", nil, err
	}

	collection := db.Database("user").Collection("userData")

	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", nil, err
	}

	return user.Token, insertResult, nil
}
