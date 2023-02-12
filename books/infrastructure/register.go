package infrastructure

import (
	"github.com/andyollylarkin/chitai_tech_test_task/books"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEnpoints(router *gin.Engine, service books.Service) {
	h := NewHandler(service)
	router.GET("/books", h.GetAll)
	router.DELETE("/books", h.Remove)
}
