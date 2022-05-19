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

type votes struct {
	NameGame      string
	TotalVotes    int32
	TotalSumVotes int64
}

func (service *Server) UserVote(ctx context.Context, in *pb.UserVoteRequest) (*pb.UserVoteResponse, error) {

	valorantVotes := votes{"Valorant", 0, 0}
	csVotes := votes{"CS", 0, 0}
	lolVotes := votes{"LOL", 0, 0}
	dotaVotes := votes{"DOTA2", 0, 0}
	eternalRetunrVotes := votes{"EternalReturn", 0, 0}

	gameName := in.GetNameGame()

	switch {
	case gameName == "Valorant":
		valorantVotes.TotalVotes += 1
		valorantVotes.TotalSumVotes += 4

	case gameName == "CS":
		csVotes.TotalVotes += 1
		csVotes.TotalSumVotes += 4

	case gameName == "LOL":
		lolVotes.TotalVotes += 1
		lolVotes.TotalSumVotes += 4

	case gameName == "DOTA2":
		dotaVotes.TotalVotes += 1
		dotaVotes.TotalSumVotes += 4

	case gameName == "EternalReturn":
		eternalRetunrVotes.TotalVotes += 1
		eternalRetunrVotes.TotalSumVotes += 4
	}

	voteRes := &pb.UserVoteResponse{
		ResponseVote: in.GetTotalVotes(),
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
