package main

import (
	handlerPkg "app/internal/handler"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println(echo.DELETE)

	handler := handlerPkg.NewHandler()

	e := echo.New()

	e.GET("/hello", handler.Hello)

	e.Logger.Fatal(e.Start(":8066"))
}
