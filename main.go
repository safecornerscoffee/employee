package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/safecornerscoffee/employee/db"
	"github.com/safecornerscoffee/employee/handler"

	_ "github.com/lib/pq"
)

func main() {

	db := db.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	h := &handler.Handler{DB: db}

	e.POST("/employee", h.CreateEmployee)
	e.PUT("/employee", h.UpdateEmployee)
	e.DELETE("/empoyee/:id", h.DeleteEmployee)
	e.GET("/employee/:id", h.GetEmployee)
	e.GET("/employee", h.GetEmployees)

	e.Logger.Fatal(e.Start(":8080"))
}
