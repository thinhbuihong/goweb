package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	HelloSV "github.com/thinhbuihong/goweb/grpc/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := HelloSV.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &HelloSV.StringM{Value: "thinh dep zai"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("====bi========")
	stream, err := client.BiChannel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() { //send
		for {
			if err := stream.Send(&HelloSV.StringM{Value: "hi thinh dep zai"}); err != nil {
				time.Sleep(time.Second)
			}
		}
	}()

	// go func() { //receive
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
	// }()

	fmt.Println(reply.GetValue())
}
