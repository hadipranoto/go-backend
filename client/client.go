package main

import (
	"context"
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)


func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hellow from the client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err!=nil {
		log.Fatalf("error when calling sayhello: %s",err)
	}
	log.Printf("Response from the server: %s",response.Body)
}