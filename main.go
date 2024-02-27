package main

import (
	"github.com/edwinnambaje/controllers"
	initializers "github.com/edwinnambaje/initizializers"
	"github.com/edwinnambaje/middleare"
	"github.com/gin-gonic/gin"
)

func init (){
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup) 
	r.POST("/login", controllers.Login) 
	r.GET("/protect", middleware.ValidateToken, controllers.ProtectedRoute)
	r.Run()
}