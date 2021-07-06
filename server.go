package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hadipranoto/go-backend.git/chat"

	"google.golang.org/grpc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// src.StudentsQuery()
	fmt.Fprintf(w, "Hi there!")
}

type App struct {	
	HttpServer *http.Server
	GrpcServer *grpc.Server
}


func RunApp(app App){
	router := mux.NewRouter();
	router.HandleFunc("/", handler)			
	app.HttpServer = &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s","9001"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}	
	fmt.Println("[x] starting app")
	log.Fatal(app.HttpServer.ListenAndServe())	
	
}

func InitGrpcServer(app App){
	fmt.Println("[x] Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")	
	
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	app.GrpcServer = grpc.NewServer()
	s := chat.Server{}
	chat.RegisterChatServiceServer(app.GrpcServer, &s)
	
	if err := app.GrpcServer.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)		
	}	
}

func main() {	
	
	myApp := App{}
	
	//lol its tricky but it works hahaha
	go func ()  {
		RunApp(myApp)	
	}()	
	
	InitGrpcServer(myApp)
	
	
}

//https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
//https://www.youtube.com/watch?v=BdzYdN_Zd9Q

//dahlah
//https://towardsdatascience.com/grpc-in-golang-bb40396eb8b1