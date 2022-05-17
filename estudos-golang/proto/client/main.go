package main

import (
	"context"
	"log"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	utils.GetErro(err)

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Oi",
	}

	res, err := client.RequestMessage(context.Background(), req)
	utils.GetErro(err)

	log.Println(res.GetStatus())
}
