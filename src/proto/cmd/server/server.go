package main

import (
	"fmt"
	"net"

	"github.com/Alexsander-Espindola/API-with-golang/src/proto/pb"
	"github.com/Alexsander-Espindola/API-with-golang/src/proto/servers"
	"github.com/Alexsander-Espindola/API-with-golang/src/utils"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	fmt.Print("start server")

	pb.RegisterPostUserServer(grpcServer, &servers.Server{})

	listener, err := net.Listen("tcp", ":50051")
	utils.GetErro(err)

	grpcErro := grpcServer.Serve(listener)
	utils.GetErro(grpcErro)

	fmt.Println("Servidor ouvindo na porta 50051")
}
