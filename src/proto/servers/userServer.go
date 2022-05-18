package servers

import (
	"context"

	"github.com/Alexsander-Espindola/API-with-golang/src/model"
	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
)

type Server struct {
	pb.UnimplementedPostUserServer
}

func (service *Server) PostUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	user := model.User{
		Name:     in.User.GetName(),
		Email:    in.User.GetEmail(),
		Password: in.User.GetPassword(),
	}

	result, _, err := model.InsertUser(user)
	utils.GetErro(err)

	return &pb.UserResponse{
		Token: result,
	}, nil
}
