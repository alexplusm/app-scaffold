package main

import (
	"app/internal/db"
	handlerPkg "app/internal/handler"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	conn, err := db.NewConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewConnection failed: %v\n", err)
		os.Exit(1)
	}

	handler := handlerPkg.NewHandler(conn)
	defer handler.Close()

	e := echo.New()

	e.GET("/hello", handler.Hello)
	e.GET("/test", handler.Test)

	e.Logger.Fatal(e.Start(":8066"))
}
