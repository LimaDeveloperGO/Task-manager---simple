package main

import (
	"log"
	"task-manager/internal/server"
)

func main() {
	srv, err := server.NewServer("3000")

	if err != nil {
		log.Fatal("❌ Erro ao criar servidor:", err)
	}
	defer srv.Close()
	if err := srv.Start(); err != nil {
		log.Fatal("❌ Erro ao iniciar servidor:", err)
	}
}
