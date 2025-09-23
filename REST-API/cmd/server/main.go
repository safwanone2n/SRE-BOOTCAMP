package main

import (
	"context"
	"log"
	"os"
	"rest-api/db"
	"rest-api/internal/user/handlers"
	userRepo "rest-api/internal/user/repo"
	"rest-api/internal/user/usecases"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func main() {

	//initializing server
	server := InitializeDi()
	server.Engine.Run(os.Getenv("PORT"))

}

func RegisterRoutes(r *gin.Engine, userHandlerI handlers.UserHandlerI) {

	user := r.Group("api/v1/student/")
	{
		user.POST("/create", userHandlerI.CreateUserHandler)
		user.GET("/list", userHandlerI.GetAllUsersHandler)
		user.GET("/get", userHandlerI.GetUserHandler)
		user.PUT("/update", userHandlerI.UpdateUserHandler)
		user.PUT("/delete", userHandlerI.DeleteUserHandler)
	}

}
func InitializeDi() *Server {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//initialzing gin instance
	r := gin.Default()

	//collecting pgx pool connection
	//if required for additional operations
	//which are not provided by sqlc
	_, err := db.GetConn(ctx)
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
	userHandlerI := handlers.NewUserHandler(userUseCaseI)

	//initialize repo

	//registering api routes
	RegisterRoutes(r,userHandlerI)

	//returning server
	return &Server{
		Engine: r,
	}

}
