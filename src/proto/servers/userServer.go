package servers

import (
	"context"

	"github.com/Alexsander-Espindola/API-with-golang/src/model"
	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
)

type PostUserClient struct {
	pb.UnimplementedPostUserServer
}

func (c *PostUserClient) PostUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	endereco := model.Endereco{
		Cidade: in.GetCidade(),
		Estado: in.GetEstado(),
	}

	user := model.User{
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Endereco: endereco,
	}

	model.PostUser(user)

	return &pb.UserResponse{
		Status: 201,
	}, nil
}
