package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println("Mensagem Recebida", req.GetMessage())

	return &pb.Response{
		Status: 1,
	}, nil
}

func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	listener, err := net.Listen("tcp", ":5005")
	if err != nil {
		log.Fatal(err)
	}

	grpcErro := grpcServer.Serve(listener)
	if grpcErro != nil {
		log.Fatal(grpcErro)
	}
	fmt.Println("Servidor ouvindo na porta 5005")
}
