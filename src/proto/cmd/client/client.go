package main

import (
	"context"
	"log"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro na conexão com gRPC server: %v", err)
	}

	defer conn.Close()
	client := pb.NewPostUserClient(conn)

	// req := &pb.UserRequest{
	// 	User: &pb.User{
	// 		Name:     "Alexs",
	// 		Email:    "Alexs@gmail.com",
	// 		Password: "1234567",
	// 	},
	// }

	// res, err := client.PostUser(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("Erro durante a requisição: %v", err)
	// }
	// log.Println(res)

	reqVote := &pb.UserVoteRequest{
		NameGame:      "Valorant",
		TotalVotes:    20,
		TotalSumVotes: 126,
	}

	resVote, err := client.UserVote(context.Background(), reqVote)
	if err != nil {
		log.Fatalf("Erro durante a requisição: %v", err)
	}
	log.Println(resVote)

}

// streamReq := &pb.StreamRequest{
// 	StrRequest: "Olá Stream",
// }
// responseSteam, err := client.TestStream(context.Background(), streamReq)
// if err != nil {
// 	log.Fatal(err)
// }

// for {
// 	stream, err := responseSteam.Recv()
// 	fmt.Printf("Stream: %v", stream.GetStrResponse())
// 	if err != nil {
// 		break
// 	}
// }
