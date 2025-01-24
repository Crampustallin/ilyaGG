package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/Crampustallin/ilyaGG/internals/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	runServer()
}

func runServer() {
	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := server.NewServer()

	go func() {
		log.Println("http server listening at [::]", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("gin error: %s\n", err)
		}
	}()

	<-signalCtx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// Оповещаем о завершении через заданное время
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
