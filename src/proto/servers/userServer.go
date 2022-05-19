package servers

import (
	"context"
	"fmt"
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

type votes struct {
	NameGame      string
	TotalVotes    int32
	TotalSumVotes int32
}

func (service *Server) UserVote(ctx context.Context, in *pb.UserVoteRequest) (*pb.UserVoteResponse, error) {

	valorantVotes := votes{"Valorant", 0, 0}
	csVotes := votes{"CS", 1, 1}
	lolVotes := votes{"LOL", 1, 1}
	dotaVotes := votes{"DOTA2", 1, 1}
	eternalRetunrVotes := votes{"EternalReturn", 1, 1}

	gameName := in.GetNameGame()
	totalSumVotes := in.GetTotalSumVotes()
	totalVotes := in.GetTotalVotes()
	fmt.Println(gameName, totalSumVotes, totalVotes)

	switch {
	case gameName == "Valorant":
		valorantVotes.TotalSumVotes = totalSumVotes
		valorantVotes.TotalVotes = totalVotes

	case gameName == "CS":
		csVotes.TotalSumVotes = totalSumVotes
		csVotes.TotalVotes = totalVotes

	case gameName == "LOL":
		lolVotes.TotalSumVotes = totalSumVotes
		lolVotes.TotalVotes = totalVotes

	case gameName == "DOTA2":
		dotaVotes.TotalSumVotes = totalSumVotes
		dotaVotes.TotalVotes = totalVotes

	case gameName == "EternalReturn":
		eternalRetunrVotes.TotalSumVotes = totalSumVotes
		eternalRetunrVotes.TotalVotes = totalVotes
	}

	voteRes := &pb.UserVoteResponse{
		TotalVavaVotes:     valorantVotes.TotalSumVotes / valorantVotes.TotalVotes,
		TotalCsVote:        csVotes.TotalSumVotes / csVotes.TotalVotes,
		TotalLolzinhoVotes: lolVotes.TotalSumVotes / lolVotes.TotalVotes,
		TotalDotinhaVotes:  dotaVotes.TotalSumVotes / dotaVotes.TotalVotes,
		TotalBestMobaVotes: eternalRetunrVotes.TotalSumVotes / eternalRetunrVotes.TotalVotes,
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
