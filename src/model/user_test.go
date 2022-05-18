package model_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Alexsander-Espindola/API-with-golang/src/model"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestPrepareUser(t *testing.T) {
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

type mockDatabase struct {
}

func (m *mockDatabase) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {

	// Tentei pesar em algo mais útil para esse mock fazer mas parece que ele só vai verificar se o retorno dessa função é do tipo InsertOneResult
	ioOpts := options.MergeInsertOneOptions(opts...)
	imOpts := options.InsertMany()

	if ioOpts.BypassDocumentValidation != nil && *ioOpts.BypassDocumentValidation {
		imOpts.SetBypassDocumentValidation(*ioOpts.BypassDocumentValidation)
	}
	// res, err := coll.insert(ctx, []interface{}{document}, imOpts)

	// rr, err := processWriteError(err)
	// if rr&rrOne == 0 {
	// 	return nil, err
	// }
	return &mongo.InsertOneResult{}, nil
}

func TestInsertUser(t *testing.T) {
	user := model.User{
		Name:     "Alexs",
		Email:    "Alexs@gmail.com",
		Password: "1234567",
	}
	err := user.PrepareUser()
	require.Nil(t, err)

	mockDb := &mockDatabase{}

	insertResult, err := mockDb.InsertOne(context.Background(), user)
	require.Nil(t, err)
	require.IsType(t, &mongo.InsertOneResult{}, insertResult)

	userNil := model.User{}

	insertResult, err = mockDb.InsertOne(context.Background(), userNil)
	require.Nil(t, err)
	require.IsType(t, &mongo.InsertOneResult{}, insertResult)

}
