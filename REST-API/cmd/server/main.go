package main

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
}

func main() {

}

func RegisterRoutes(r *gin.Engine) {

	user := r.Group("api/v1/student/")
	{
		user.POST("/create")
		user.GET("/list")
		user.GET("/get")
		user.PUT("/update")
		user.PUT("/delete")
	}

}
func InitializeDi() *Server {
	//initialzing gin instance
	r := gin.Default()

	//registering api routes
	RegisterRoutes(r)

	//returning server
	return &Server{
		Engine: r,
	}

}
