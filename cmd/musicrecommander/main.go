package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	limits "github.com/gin-contrib/size"
	"github.com/joho/godotenv"

	"github.com/AhmadRezaZarei/musicrecommander/modules/song"
	"github.com/gin-gonic/gin"
)

func main() {

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = godotenv.Load(curDir + "/.env")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	r := gin.Default()
	r.Use(limits.RequestSizeLimiter(8 << 20))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// start: Add routes ----------
	song.AddRoutes(r)
	// end: ------------------

	r.Run(fmt.Sprintf("0.0.0.0:%s", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
