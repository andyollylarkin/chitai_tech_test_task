package infrastructure

import (
	"github.com/andyollylarkin/chitai_tech_test_task/books"
	"github.com/andyollylarkin/chitai_tech_test_task/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service books.Service
}

func NewHandler(service books.Service) *Handler {
	return &Handler{service: service}
}

type response struct {
	Id            int64  `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Author        string `json:"author,omitempty"`
	PublisherYear int    `json:"publisher_year,omitempty"`
}

func (h *Handler) GetAll(c *gin.Context) {
	booksCollection := h.service.RetrieveBooks()
	resp := make([]response, 0)
	for _, v := range booksCollection {
		r := response{Id: v.GetId(), Title: v.GetTitle(), Author: v.GetAuthor(), PublisherYear: v.GetPublisherYear()}
		resp = append(resp, r)
	}
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, resp)
}

type userInput struct {
	Id int64 `json:"book_id"`
}

func (h *Handler) Remove(c *gin.Context) {
	inp := new(userInput)
	logger := tools.NewLogger()
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	logger.Info().Msgf("Start deleting book with id %d", inp.Id)
	err := h.service.RemoveBook(inp.Id)
	if err != nil {
		logger.Error().Err(err)
		c.AbortWithError(http.StatusNotFound, err)
	}
	c.Status(http.StatusOK)
}
