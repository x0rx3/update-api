package api

import "github.com/gin-gonic/gin"

type ServerHandler interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Patch(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Post(ctx *gin.Context)
}

type ResultHandler interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
