package main

import (
	"context"
	"log"
	"os"
	"rest-api/db"
	"rest-api/internal/user/handlers"
	userRepo "rest-api/internal/user/repo"
	"rest-api/internal/user/usecases"

	_ "rest-api/config"

	"github.com/gin-gonic/gin"
	"github.com/remiges-tech/logharbour/logharbour"
)

type Server struct {
	Engine *gin.Engine
}

func main() {

	//initializing server
	server := InitializeDi()
	server.Engine.Run(os.Getenv("PORT"))

}

func InitializeDi() *Server {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
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
	_, err = db.GetConn(ctx)
	if err != nil {
		log.Fatal("error acquiring connection for db")
	}
	//initializing querier
	querier, err := db.GetQuerier(ctx)
	if err != nil {
		log.Fatal("error acquiring querier for db operations")
	}

	//initializing user repo
	userRepoI := userRepo.NewUserRepo(querier)
	userUseCaseI := usecases.NewUserUseCases(userRepoI)
	userHandlerI := handlers.NewUserHandler(userUseCaseI,logger)

	//initialize repo

	//registering api routes
	handlers.RegisterRoutes(r, userHandlerI)

	//returning server
	return &Server{
		Engine: r,
	}

}
