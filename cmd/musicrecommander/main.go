package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	limits "github.com/gin-contrib/size"
	"github.com/joho/godotenv"

	"github.com/AhmadRezaZarei/musicrecommander/database"
	"github.com/AhmadRezaZarei/musicrecommander/modules/recognition"
	"github.com/AhmadRezaZarei/musicrecommander/modules/songlog"
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

	database.Migrate()

	port := os.Getenv("PORT")

	r := gin.Default()
	r.Use(limits.RequestSizeLimiter(8 << 20))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// start: Add routes ----------
	songlog.AddRoutes(r)
	// end: ------------------

	// start: Add jobs
	err = recognition.StartJob()
	if err != nil {
		panic(err)
	}

	// end: -------------

	r.Run(fmt.Sprintf("0.0.0.0:%s", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
