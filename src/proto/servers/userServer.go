package servers

import (
	"context"
	"math/rand"
	"time"

	"github.com/Alexsander-Espindola/API-with-golang/src/model"
	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		TotalVavaVotes:     in.GetTotalSumVavaVotes(),
		TotalCsVote:        in.GetTotalSumCsVote(),
		TotalLolzinhoVotes: in.GetTotalSumLolzinhoVotes(),
		TotalDotinhaVotes:  in.GetTotalSumDotinhaVotes(),
		TotalBestMobaVotes: in.GetTotalSumBestMobaVotes(),
	}

	return voteRes, nil
}

func (service *Server) TestStream(in *pb.StreamRequest, stream pb.PostUser_TestStreamServer) error {

	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Encerrando Stream")
		default:
			time.Sleep(time.Minute)
			value := 5 + rand.Intn(10)

			err := stream.SendMsg(&pb.StreamResponse{
				StrResponse: string(rune(value)),
			})
			if err != nil {
				return status.Error(codes.Canceled, "Encerrando Stream")
			}
		}
	}
}
