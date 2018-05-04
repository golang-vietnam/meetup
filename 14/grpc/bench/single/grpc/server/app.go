package main

import (
	"context"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {
	// try to listening to bind address
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		return
	}
	s := grpc.NewServer()
	RegisterHelloServer(s, &helloServer{})

	if e := s.Serve(listen); e != nil {
		panic(e)
	}
}

type helloServer struct{}

func (c *helloServer) SayHi(ctx context.Context, req *Payload) (resp *Payload, err error) {
	return req, nil
}
