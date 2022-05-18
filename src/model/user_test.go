package model_test

import (
	"fmt"
	"testing"

	"github.com/Alexsander-Espindola/API-with-golang/src/model"
	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	fmt.Println("Testando um usuario valido")
	user := model.User{
		Name:     "Alexs",
		Email:    "Alexs@gmail.com",
		Password: "1234567",
	}
	err := user.PrepareUser()
	require.Nil(t, err)

	userEmailInvalid := model.User{
		Name:     "Alexsander",
		Email:    "Alexs@gmail",
		Password: "1234567",
	}
	errorInEmailValidate := userEmailInvalid.PrepareUser()
	require.Error(t, errorInEmailValidate)

	userPasswordInvalid := model.User{
		Name:     "Alexs3",
		Email:    "Alexsander@gmail.com",
		Password: "",
	}
	errorInPasswordValidate := userPasswordInvalid.PrepareUser()
	require.Error(t, errorInPasswordValidate)
}
