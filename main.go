package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// src.StudentsQuery()

	fmt.Println("Hello world")
	

}

type App struct {
	Router *mux.Router	
}

func (a *App) Initialize () {		
	r := mux.NewRouter();
	r.HandleFunc("/", handler)	
	a.Router = r 
}
func (a *App) RunApp(port string){
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%s",port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func main() {	
	myApp := App{}
	myApp.Initialize()
	myApp.RunApp("9001")
	

}
