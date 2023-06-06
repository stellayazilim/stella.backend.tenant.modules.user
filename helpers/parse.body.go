package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// return true if body exist
func ParseBody[T any](ctx *gin.Context) T {

	var body T
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	return body
}
