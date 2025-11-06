package main

import (
	proto "Question2/grpc"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Mopper struct {
	proto.UnimplementedMopperServer
	queue []int64
}

func main() {
	server := &Mopper{queue: make([]int64, 0)}
	server.start_server()
}

func (m *Mopper) start_server() {
	grpcserver := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Server did not work")
	}

	proto.RegisterMopperServer(grpcserver, m)

	err = grpcserver.Serve(listener)
	if err != nil {
		log.Fatalf("this did not work")
	}
}

func (m *Mopper) RequestToken(ctx context.Context, request *proto.Request) (*proto.GrantAccess, error) {
	m.queue = append(m.queue, request.GetID())

	for {
		if request.GetID() == m.queue[0] {
			break
		}
	}

	return &proto.GrantAccess{}, nil
}

func (m *Mopper) ReleaseToken(ctx context.Context, release *proto.Release) (*proto.Empty, error) {
	if m.queue[0] == release.GetID() {
		m.queue = m.queue[1:]
	}

	return &proto.Empty{}, nil
}
