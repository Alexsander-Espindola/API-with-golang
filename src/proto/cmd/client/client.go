package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	fmt.Println("Abrindo um channel")
	votation := make(chan int32)
	fmt.Println("Chamando função ChannelForGame")
	go ChannelForGame("Valorant", votation, client)
	for i := 0; i < 10; i++ {
		votation <- int32(i)
		time.Sleep(time.Second)
	}

	// fmt.Println("chamando client.UserVote")
	// resVote, err := client.UserVote(context.Background(), reqVote)
	// if err != nil {
	// 	log.Fatalf("Erro durante a requisição: %v", err)
	// }
	// log.Println(resVote)
}

// func UserVoteRand() {

// }

func ChannelForGame(nameGame string, channelGame chan int32, client pb.PostUserClient) {
	var totalVote int32
	var totalSumVotes int32
	totalVote = 0
	totalSumVotes = 0
	for x := range channelGame {
		fmt.Println("Print do channelGame")
		totalVote += 1
		totalSumVotes += x
		reqVote := &pb.UserVoteRequest{
			NameGame:      nameGame,
			TotalVotes:    totalVote,
			TotalSumVotes: totalSumVotes,
		}
		resVote, err := client.UserVote(context.Background(), reqVote)
		if err != nil {
			log.Fatalf("Erro durante a requisição: %v", err)
		}
		log.Println(resVote)
	}
}
