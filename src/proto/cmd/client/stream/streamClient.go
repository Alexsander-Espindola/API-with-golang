package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"google.golang.org/grpc"
)

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

	stream, err := client.TestStream(context.Background(), &pb.StreamRequest{})
	if err != nil {
		log.Fatalln("Couldn't request", err)
	}

	go func() {
		for {
			value, err := stream.Recv()
			fmt.Println(value)
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			ui.Render()
		}
	}()

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents: // Listen for Keyboard
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-stream.Context().Done():
			fmt.Println("All done")
			break
		}
	}

}
