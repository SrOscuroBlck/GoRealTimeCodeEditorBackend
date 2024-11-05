package main

import (
	"GoRealTimeCodeEditor/internal/server"
	"log"
)

func main() {
	srv := server.New()
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
