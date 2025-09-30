package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/config"
	"rest-api/db"
	"rest-api/internal/user/handlers"
	userRepo "rest-api/internal/user/repo"
	userService "rest-api/internal/user/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remiges-tech/logharbour/logharbour"
)

type Server struct {
	cfg    *config.Config
	Engine *gin.Engine
}

// main sets up the root context and handles graceful shutdown.
func main() {
	// create a root context that is cancelled on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// initialize server with the shared context
	server := InitializeDi(ctx)

	// configure http.Server to allow graceful shutdown
	srv := &http.Server{
		Addr:    server.cfg.ServerPort,
		Handler: server.Engine,
	}

	go func() {
		log.Printf("starting server on %s", server.cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server exited with error: %v", err)
		}
	}()

	// waiting for interrupt signal
	<-ctx.Done()
	log.Println("shutdown signal received")

	// create a new context with timeout for shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// attempt graceful shutdown
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server gracefully stopped")
}

// InitializeDi wires dependencies using the provided context.
func InitializeDi(ctx context.Context) *Server {
	//initializing configs
	cfg := config.New()

	//initializing gin instance
	r := gin.Default()

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := logharbour.NewLogger(logharbour.NewLoggerContext(logharbour.Info), "REST-API", file)

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	//collecting pgx pool connection
	//if required for additional operations
	//which are not provided by sqlc
	_, err = db.GetConn(ctx, cfg)
	if err != nil {
		log.Fatal("error acquiring connection for db")
	}

	//initializing querier
	querier, err := db.GetQuerier(ctx, cfg)
	if err != nil {
		log.Fatal("error acquiring querier for db operations")
	}

	//initializing user repo
	userRepo := userRepo.NewUserRepo(querier)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, logger)

	//registering api routes
	handlers.RegisterRoutes(r, userHandler)

	//returning server
	return &Server{
		Engine: r,
		cfg:    cfg,
	}
}
