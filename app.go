package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

var static embed.FS

func main() {
	r := gin.Default()
	r.Static("/", "./static")
	http.Handle("/static", http.FileServer(http.FS(static)))
	r.Run() // listen and serve on 0.0.0.0:8080
}
