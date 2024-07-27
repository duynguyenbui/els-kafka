package main

import (
	"context"
	"fmt"
	"internacs-els-kafka/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	// Make sure background process is running
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()
}
