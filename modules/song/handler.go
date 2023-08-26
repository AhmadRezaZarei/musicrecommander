package song

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {

	r.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")

		if err == nil {
			c.SaveUploadedFile(file, filepath.Join("songs", file.Filename))
		}

		c.Request.FormValue("song_id")

		c.JSON(http.StatusOK, gin.H{
			"error": nil,
		})
	})

}
