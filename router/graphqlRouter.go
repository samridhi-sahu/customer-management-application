package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func GraphQLRoutes(srv *handler.Server, router *gin.Engine) {
	router.POST("/graphql", gin.WrapH(srv))
}
