package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_ "github.com/lib/pq"
)

func main() {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresDb := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDb)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < 360; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	type Employee struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Salary string `json:"salary"`
		Age    string `json:"age"`
	}

	type Employees struct {
		Employees []Employee `json:"employees"`
	}

	e.POST("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := `
		INSERT INTO employees (name, salary, age)
		VALUES ($1, $2, $3)
		RETURNING id`

		err := db.QueryRow(sqlStatement, u.Name, u.Salary, u.Age).Scan(&u.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := `
		UPDATE employees SET name=$1, salary=$2, age=$3
		WHERE id=$4`
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age, u.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, u.Id)
	})

	e.DELETE("/empoyee/:id", func(c echo.Context) error {
		id := c.Param("id")
		sqlStatement := `DELETE FROM employees
		WHERE id=$1`
		res, err := db.Query(sqlStatement, id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusOK, "Deleted")
		}
		return c.String(http.StatusOK, id+"Deleted")
	})

	e.GET("/employee/:id", func(c echo.Context) error {
		id := c.Param("id")
		u := new(Employee)
		sqlStatement := `SELECT id, name, salary, age FROM employees
		where id=$1`
		row := db.QueryRow(sqlStatement, id)
		err := row.Scan(&u.Id, &u.Name, &u.Salary, &u.Age)
		switch err {
		case sql.ErrNoRows:
			return c.String(http.StatusNotFound, "not found")
		case nil:
			return c.JSON(http.StatusOK, u)
		default:
			fmt.Println(err)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/employee", func(c echo.Context) error {
		sqlStatement := "SELECT id, name, salary, age FROM employees order by id"
		rows, err := db.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		}
		defer rows.Close()
		result := Employees{}

		for rows.Next() {
			employee := Employee{}
			err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
			// Exit if we get an error
			if err2 != nil {
				return err2
			}
			result.Employees = append(result.Employees, employee)
		}
		return c.JSON(http.StatusCreated, result)

		//return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
