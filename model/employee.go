package model

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
