package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"funding/internal/application/user"
	"funding/internal/application/user/handler"
	infraEvent "funding/internal/infrastructure/event"
	"funding/internal/infrastructure/persistence/memory"
	httpHandler "funding/internal/transport/handler"
)

func main() {
	// 1. Initialize Infrastructure
	// Use In-Memory Repository for demo purposes (switch to Postgres in production)
	userRepo := memory.NewInMemoryUserRepository()

	// 2. Initialize Event Bus and Handlers
	// Mock Mailer for now
	mailer := &MockMailer{}
	welcomeEmailHandler := handler.NewSendWelcomeEmailHandler(mailer)

	// Register handlers to the bus
	// Note: In a real app, we might want a more dynamic registration
	eventBus := infraEvent.NewInMemoryBus(welcomeEmailHandler)

	// 3. Initialize Use Cases
	userUseCase := user.NewRegisterUserUseCase(userRepo, eventBus)

	// 4. Initialize HTTP Handlers
	userHandler := httpHandler.NewUserHandler(userUseCase)

	// 5. Setup Router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.Register)

	// 6. Start Server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

// MockMailer for demonstration
type MockMailer struct{}

func (m *MockMailer) Send(ctx context.Context, msg handler.WelcomeEmail) error {
	log.Printf("[MockMailer] Sending welcome email to %s (%s)\n", msg.Name, msg.Email)
	return nil
}
