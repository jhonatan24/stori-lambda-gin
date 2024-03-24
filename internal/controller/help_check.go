package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelpCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "hola mundo")
}
