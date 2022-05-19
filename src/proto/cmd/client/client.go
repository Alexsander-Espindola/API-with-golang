package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

	if err := ui.Init(); err != nil {
		log.Fatalf("Couldn't init UI: %v", err)
	}
	defer ui.Close()

	lc := widgets.NewPlot()
	ui.Render(lc)

	channelGameVote := make(chan int)

	go ChannelForGame(channelGameVote, client)
	var stream pb.PostUser_TestStreamClient

	for i := 0; i < 100; i++ {
		stream, err = client.TestStream(context.Background(), &pb.StreamRequest{
			TotalSumVavaVotes:     TotalSumVavaVotes,
			TotalSumCsVote:        TotalSumCsVote,
			TotalSumLolzinhoVotes: TotalSumLolzinhoVotes,
			TotalSumDotinhaVotes:  TotalSumDotinhaVotes,
			TotalSumBestMobaVotes: TotalSumBestMobaVotes,
		})
		if err != nil {
			log.Fatalf("Erro na requisição da stream: %v", err)
		}
		gameRand1 := 1 + rand.Intn(5)

		value, err := stream.Recv()
		fmt.Println("Valorant:", value.TotalVavaVotes)
		fmt.Println("CS", value.TotalCsVote)
		fmt.Println("LOL", value.TotalLolzinhoVotes)
		fmt.Println("DOTA2", value.TotalDotinhaVotes)
		fmt.Println("EternalReturn", value.TotalBestMobaVotes)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		ui.Render()

		if err != nil {
			log.Fatalf("Erro durante a requisição: %v", err)
		}
		channelGameVote <- gameRand1
		time.Sleep(time.Second)
	}

	uiEvents := ui.PollEvents()

	for {
		fmt.Println("Context done")
		select {
		case e := <-uiEvents: // Listen for Keyboard
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-stream.Context().Done():
			fmt.Println("All done")
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
			break
		}
	}
}

func UserVoteRand(idGame int, client pb.PostUserClient) {
	switch {
	case idGame == 1:
		TotalSumVavaVotes += 1

	case idGame == 2:
		TotalSumCsVote += 1

	case idGame == 3:
		TotalSumLolzinhoVotes += 1

	case idGame == 4:
		TotalSumDotinhaVotes += 1

	case idGame == 5:
		TotalSumBestMobaVotes += 1
	}
}

func ChannelForGame(channelGame chan int, client pb.PostUserClient) {
	for idGame := range channelGame {
		UserVoteRand(idGame, client)
	}
}
