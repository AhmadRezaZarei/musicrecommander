package ginutil

import (
	"net/http"

	"github.com/AhmadRezaZarei/musicrecommander/modules/util"
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, err *util.MainError) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": err,
	})
}

func SendErrorWithStatusCode(ctx *gin.Context, statusCode int, err *util.MainError) {
	ctx.JSON(statusCode, gin.H{
		"error": err,
	})
}

func SendWrappedInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": &util.MainError{
			Type:    util.ErrInternalServerError,
			Message: err.Error(),
		},
	})
}

func sendEmptySucessResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
