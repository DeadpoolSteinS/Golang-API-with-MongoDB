package main

import (
	"net/http"
	"practice_1/controllers"
	"practice_1/initializers"

	"github.com/julienschmidt/httprouter"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := httprouter.New()

	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)

	http.ListenAndServe("localhost:3000", r)
}
