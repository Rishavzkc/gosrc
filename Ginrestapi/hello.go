package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", display)
	r.Run(":7898")

}
func display(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}
