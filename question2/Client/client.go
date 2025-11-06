package main

import (
	proto "Question2/grpc"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ID int64

func main() {
	ID, _ := strconv.ParseInt(os.Args[1], 10, 32)

	fmt.Println("I am client: ", ID)

	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect")
	}

	client := proto.NewMopperClient(conn)

	client.RequestToken(context.Background(), &proto.Request{ID: ID})

	log.Printf("Yay!!! jeg har fået adgang! nu er jeg inde i Critical Section")
	time.Sleep(5000 * time.Millisecond)

	client.ReleaseToken(context.Background(), &proto.Release{ID: ID})

	log.Printf("Jeg har nu relesed, så andre kan få adgang")

}
