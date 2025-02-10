package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// func dd(data interface{}) {
// 	jsonData, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Erro ao serializar dados:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(string(jsonData))
// 	os.Exit(0)
// }

func main() {

	fmt.Println("Iniciando o servidor...")

	server := InitializeServer()

	// Usar contexto para gerenciar o ciclo de vida do servidor
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := server.Start(); err != nil {
			server.logger.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	<-ctx.Done()
	server.logger.Println("Desligando o servidor...")

	// Tempo de espera para o servidor desligar
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(shutdownCtx); err != nil {
		server.logger.Fatalf("Erro ao desligar o servidor: %v", err)
	}

	server.logger.Println("Servidor desligado com sucesso")
}

// go run cmd/server/main.go
