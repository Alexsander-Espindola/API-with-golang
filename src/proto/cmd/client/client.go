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

	client := pb.NewPostUserClient(conn)
	defer conn.Close()

	req := &pb.UserRequest{
		User: &pb.User{
			Name:     "Alexs",
			Email:    "Alexs@gmail.com",
			Password: "1234567",
		},
	}

	res, err := client.PostUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro durante a requisição: %v", err)
	}

	log.Println(res)
}
