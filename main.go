package main

import (
	"net/http"
	"runtime"
	"github.com/PostSample/controllers/check"

	"github.com/PostSample/middleware/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	r := gin.Default()
	//Set Only allow Patch GET POST OPTION Method
	r.Use(cors.AllowAllOrigins())

	r.POST("/", check.GetPath)

	http.ListenAndServe(":8080", context.ClearHandler(r))

}
