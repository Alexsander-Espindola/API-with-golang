package servers

import (
	"context"

	"github.com/Alexsander-Espindola/API-with-golang/src/model"
	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
)

type Server struct {
	pb.UnimplementedPostUserServer
}

func (service *Server) PostUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	endereco := model.Endereco{
		Cidade: in.User.GetCidade(),
		Estado: in.User.GetEstado(),
	}

	user := model.User{
		Name:     in.User.GetName(),
		Email:    in.User.GetEmail(),
		Endereco: endereco,
	}

	model.PostUser(user)

	return &pb.UserResponse{
		Status: 201,
	}, nil
}
