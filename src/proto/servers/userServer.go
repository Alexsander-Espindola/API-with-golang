package servers

import (
	"context"
	"time"

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

func (service *Server) UserVote(ctx context.Context, in *pb.UserVoteRequest) (*pb.UserVoteResponse, error) {

	voteRes := &pb.UserVoteResponse{
		ResponseVote: in.GetRequestVote(),
	}

	return voteRes, nil
}

type StreamStruct struct {
}

func (c *StreamStruct) TestStream(ctx context.Context, in *pb.StreamRequest, stream pb.PostUser_TestStreamServer) error {
	message := in.GetStrRequest()

	for i := 0; i < 5; i++ {
		res := &pb.StreamResponse{
			StrResponse: message,
		}
		stream.Send(res)
		time.Sleep(time.Second * 2)
	}
	return nil
}
