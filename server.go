package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// src.StudentsQuery()
	fmt.Fprintf(w, "Hi there!")
}

type App struct {
	Router *mux.Router	
	HttpServer *http.Server
	GrpcServer *grpc.Server
}

func (a *App) Initialize () {		
	a.Router = mux.NewRouter();
	a.Router.HandleFunc("/", handler)		
}

func (a *App) RunApp(port string){
		
	a.HttpServer = &http.Server{
		Handler:      a.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%s",port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}	
	fmt.Println("[x] starting app")
	log.Fatal(a.HttpServer.ListenAndServe())	
	
}

func (a *App) InitGrpcServer(){
	fmt.Println("[x] Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")	
	
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	a.GrpcServer = grpc.NewServer()
	
	if err := a.GrpcServer.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)		
	}	
}

func main() {	
	myApp := App{}
	//myApp.Initialize()
	//myApp.RunApp("9001")
	myApp.InitGrpcServer()
	
	
}

//https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
//https://www.youtube.com/watch?v=BdzYdN_Zd9Q

//dahlah
//https://towardsdatascience.com/grpc-in-golang-bb40396eb8b1