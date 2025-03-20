package controller

import (
	"github.com/junaidmdv/clean_architure/BookStore/intrface"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	BookService intrface.BookService
}

func NewBookController(echoCtx *echo.Echo, BookServiceObject intrface.BookService) {

	BookControllerObject := &BookController{
		BookService: BookServiceObject,
	}

	echoCtx.GET("/print-author", BookControllerObject.printAuthor)

}

func (B *BookController) printAuthor(ctx echo.Context) error {
	
	return nil

}
