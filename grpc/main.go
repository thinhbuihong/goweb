package main

import (
	"context"
	"io"
	"log"
	"net"

	HelloSV "github.com/thinhbuihong/goweb/grpc/protos"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *HelloSV.StringM) (*HelloSV.StringM, error) {
	reply := &HelloSV.StringM{Value: "hello from server " + args.Value}
	return reply, nil
}

func (p *HelloServiceImpl) BiChannel(stream HelloSV.HelloService_BiChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return nil
			}
			return err
		}

		reply := &HelloSV.StringM{Value: "hello from server " + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()

	//register service for grpc server
	HelloSV.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
