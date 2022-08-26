package main

import (
	"net/http"
	"os"
	"practice_1/controllers"
	"practice_1/initializers"

	"github.com/julienschmidt/httprouter"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")
	r := httprouter.New()

	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)

	http.ListenAndServe(":"+port, r)
}
