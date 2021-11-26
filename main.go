package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"hello/proto"
	"net"
)

type Hello struct {
}

func(h *Hello) Echo(_ context.Context, request *proto.EchoRequest) (*proto.EchoResponse, error){
	log.Log().Msg(request.Message)
	if request.Message != "" {
		return &proto.EchoResponse{Message: "yo"}, nil
	} else {
		return &proto.EchoResponse{Message: "no"}, nil
	}
}

func main() {
	server := grpc.NewServer()
	proto.RegisterEchoServiceServer(server, &Hello{})

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Log().Msg("public server started")
	if err := server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("listen server")
	}
}
