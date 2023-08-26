package main

import (
	"net/http"

	limits "github.com/gin-contrib/size"
	"github.com/joho/godotenv"

	"github.com/AhmadRezaZarei/musicrecommander/modules/song"
	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(limits.RequestSizeLimiter(8 << 20))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	song.AddRoutes(r)
	r.Run("0.0.0.0:3500") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
