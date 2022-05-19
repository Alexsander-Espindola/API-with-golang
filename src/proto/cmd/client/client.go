package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	streamServer "github.com/Alexsander-Espindola/API-with-golang/src/proto/cmd/streamSever"
	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"google.golang.org/grpc"
)

var TotalSumVavaVotes int32
var TotalSumCsVote int32
var TotalSumLolzinhoVotes int32
var TotalSumDotinhaVotes int32
var TotalSumBestMobaVotes int32

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
	channelGameVote := make(chan int)
	fmt.Println("Chamando função ChannelForGame")
	go ChannelForGame(channelGameVote, client)

	for i := 0; i < 10; i++ {
		gameRand := 1 + rand.Intn(5)
		stream, err := client.TestStream(context.Background(), &pb.StreamRequest{
			// TotalSumVavaVotes:     TotalSumVavaVotes,
			// TotalSumCsVote:        TotalSumCsVote,
			// TotalSumLolzinhoVotes: TotalSumLolzinhoVotes,
			// TotalSumDotinhaVotes:  TotalSumDotinhaVotes,
			// TotalSumBestMobaVotes: TotalSumBestMobaVotes,
		})
		streamServer.StreamClient(stream, client)

		if err != nil {
			log.Fatalf("Erro durante a requisição: %v", err)
		}
		channelGameVote <- gameRand
		time.Sleep(time.Second)
	}
	reqVote := &pb.UserVoteRequest{
		TotalSumVavaVotes:     TotalSumVavaVotes,
		TotalSumCsVote:        TotalSumCsVote,
		TotalSumLolzinhoVotes: TotalSumLolzinhoVotes,
		TotalSumDotinhaVotes:  TotalSumDotinhaVotes,
		TotalSumBestMobaVotes: TotalSumBestMobaVotes,
	}

	fmt.Println("chamando client.UserVote")
	resVote, err := client.UserVote(context.Background(), reqVote)
	if err != nil {
		log.Fatalf("Erro durante a requisição: %v", err)
	}
	log.Println(resVote)
}

func UserVoteRand(idGame int, client pb.PostUserClient) {
	switch {
	case idGame == 1:
		fmt.Println("Channel Valorant")
		TotalSumVavaVotes += 1

	case idGame == 2:
		fmt.Println("Channel CS")
		TotalSumCsVote += 1

	case idGame == 3:
		fmt.Println("Channel LOL")
		TotalSumLolzinhoVotes += 1

	case idGame == 4:
		fmt.Println("Channel DOTA2")
		TotalSumDotinhaVotes += 1

	case idGame == 5:
		fmt.Println("Channel Etrn")
		TotalSumBestMobaVotes += 1
	}
}

func ChannelForGame(channelGame chan int, client pb.PostUserClient) {
	for idGame := range channelGame {
		fmt.Println("Print do channelGame")
		UserVoteRand(idGame, client)
	}
}
