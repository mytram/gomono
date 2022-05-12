package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func readyz(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
		"time":   time.Now(),
	})
}
