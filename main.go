package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	books = make(map[int64]*Book)
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/books", listBooks)
	e.GET("/book/:id", getBook)
	e.POST("/book", createBook)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Book data structure
type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var books map[int64]*Book
var seq = 1

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func listBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func getBook(c echo.Context) (err error) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, books[id])
}

func createBook(c echo.Context) (err error) {
	b := new(Book)
	if err = c.Bind(b); err != nil {
		return
	}
	b.ID = int64(seq)
	seq++
	books[b.ID] = b
	return c.JSON(http.StatusOK, b)
}
