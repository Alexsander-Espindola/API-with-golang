package model

import (
	"context"
	"fmt"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `bson:"name" valid:"notnull"`
	Email    string `bson:"email" valid:"notnull, email"`
	Password string `bson:"password" valid:"notnull"`
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

func InsertUser(user User) (string, error) {
	db := GetConnection()
	err := user.PrepareUser()
	if err != nil {
		return "", err
	}

	collection := db.Database("user").Collection("userData")

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	return user.Token, nil
}
