package server

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
)

func LaunchServer(){
	mux := http.NewServeMux()
	mux.HandleFunc("/version", GetVersion)
	mux.HandleFunc("/decode", JsonDecoder)
	mux.HandleFunc("/hard-op", Sleeping)
	http_server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}


	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("server launched")
		if err := http_server.ListenAndServe(); err!=nil && err != http.ErrServerClosed{
			log.Fatal(err)
		}
	}()
	
	sig := <-sigChan
	log.Printf("Received signal: %v. Proccessing remaining quiries and shuting down server...\n", sig)
	
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := http_server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	   }
	  
	log.Println("Server has been gracefully shut down.")
	  
}
