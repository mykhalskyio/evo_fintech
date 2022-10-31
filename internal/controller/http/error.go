package http

import "github.com/gin-gonic/gin"

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, response{Error: msg})
}
