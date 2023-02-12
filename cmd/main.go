package main

import (
	"context"
	"github.com/andyollylarkin/chitai_tech_test_task/books/infrastructure"
	"github.com/andyollylarkin/chitai_tech_test_task/books/models"
	"github.com/andyollylarkin/chitai_tech_test_task/books/persistance"
	"github.com/andyollylarkin/chitai_tech_test_task/books/services"
	"github.com/andyollylarkin/chitai_tech_test_task/tools"
	"github.com/andyollylarkin/go-app-lifecycle/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := tools.NewLogger()
	app := application.CreateApplication(logger, time.Second*3, time.Second*10, appMain, nil)
	err := app.Run()
	logger.Err(err)
	if err != nil {
		os.Exit(1)
	}
}

func appMain(ctx context.Context, wait func(), keeper *application.ServiceKeeper) error {
	router := gin.Default()
	logger := tools.NewLogger()

	initRouter(router)
	server := &http.Server{Addr: "0.0.0.0:8080", Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Err(err)
		}
	}()

	wait()
	println("After unlock")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func initRouter(router *gin.Engine) {
	// Этот ужасный код здесь по той причине, что не реализован отдельный хендлер на добавление книг.
	//По-хорошему FillBooks вообще не должно существовать, по этому в принципе не важно где его размещать
	book1, _ := models.NewBook(1, "Book1", "Author1", time.Now())
	book2, _ := models.NewBook(2, "Book2", "Author2", time.Now())
	book3, _ := models.NewBook(3, "Book3", "Author3", time.Now())
	books := []*models.Book{
		book1, book2, book3,
	}
	repo := persistance.NewInMemoryRepository()
	repo.FillBooks(books)
	logger := tools.NewLogger()
	service := services.NewBookService(repo, logger)
	infrastructure.RegisterHTTPEnpoints(router, service)
}
