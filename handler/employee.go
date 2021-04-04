package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/safecornerscoffee/employee/model"
)

func (h *Handler) CreateEmployee(c echo.Context) error {
	u := new(model.Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	sqlStatement := `
	INSERT INTO employees (name, salary, age)
	VALUES ($1, $2, $3)
	RETURNING id`

	err := h.DB.QueryRow(sqlStatement, u.Name, u.Salary, u.Age).Scan(&u.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}

func (h *Handler) UpdateEmployee(c echo.Context) error {
	u := new(model.Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	sqlStatement := `
	UPDATE employees SET name=$1, salary=$2, age=$3
	WHERE id=$4`
	res, err := h.DB.Query(sqlStatement, u.Name, u.Salary, u.Age, u.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, u.Id)
}

func (h *Handler) DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	sqlStatement := `DELETE FROM employees
	WHERE id=$1`
	res, err := h.DB.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Deleted")
	}
	return c.String(http.StatusOK, id+"Deleted")
}

func (h *Handler) GetEmployee(c echo.Context) error {
	id := c.Param("id")
	u := new(model.Employee)
	sqlStatement := `SELECT id, name, salary, age FROM employees
	where id=$1`
	row := h.DB.QueryRow(sqlStatement, id)
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
}

func (h *Handler) GetEmployees(c echo.Context) error {
	sqlStatement := "SELECT id, name, salary, age FROM employees order by id"
	rows, err := h.DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := model.Employees{}

	for rows.Next() {
		employee := model.Employee{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			return err2
		}
		result.Employees = append(result.Employees, employee)
	}
	return c.JSON(http.StatusCreated, result)

	//return c.String(http.StatusOK, "ok")
}
