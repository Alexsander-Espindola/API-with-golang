package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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
	fmt.Println("Chamando função ChannelForGame")
	for i := 0; i < 10; i++ {
		UserVoteRand(client)
		time.Sleep(time.Second)
	}

	// fmt.Println("chamando client.UserVote")
	// resVote, err := client.UserVote(context.Background(), reqVote)
	// if err != nil {
	// 	log.Fatalf("Erro durante a requisição: %v", err)
	// }
	// log.Println(resVote)
}

func UserVoteRand(client pb.PostUserClient) {
	gameRand := 1 + rand.Intn(5)
	value := 5 + rand.Intn(10)
	vavaVote := make(chan int)
	csVote := make(chan int)
	lolVote := make(chan int)
	dotaVote := make(chan int)
	etrnVote := make(chan int)
	switch {
	case gameRand == 1:
		fmt.Println("Abrindo um channel Vava")
		go ChannelForGame("Valorant", vavaVote, client)
		vavaVote <- value

	case gameRand == 2:
		fmt.Println("Abrindo um channel CS")
		go ChannelForGame("CS", csVote, client)
		csVote <- value

	case gameRand == 3:
		fmt.Println("Abrindo um channel LOL")
		go ChannelForGame("LOL", lolVote, client)
		lolVote <- value

	case gameRand == 4:
		fmt.Println("Abrindo um channel DOTA2")
		go ChannelForGame("DOTA2", dotaVote, client)
		dotaVote <- value

	case gameRand == 5:
		fmt.Println("Abrindo um channel Etrn")
		go ChannelForGame("EternalReturn", etrnVote, client)
		etrnVote <- value
	}
}

func ChannelForGame(nameGame string, channelGame chan int, client pb.PostUserClient) {
	var totalVote int32
	var totalSumVotes int32
	totalVote = 0
	totalSumVotes = 0
	for x := range channelGame {
		fmt.Println("Print do channelGame")
		totalVote += 1
		totalSumVotes += int32(x)
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
