# [Employee](https://www.restapiexample.com/golang-tutorial/creating-golang-api-echo-framework-postgresql/#Conclusion)



## API

| Route          | Method   | Type | Posted JSON                                                  | Description                                     |
| -------------- | -------- | ---- | ------------------------------------------------------------ | ----------------------------------------------- |
| /employee      | GET      | JSON | –                                                            | Get all employees data                          |
| /employee/{id} | GET      | JSON | –                                                            | Get a single employee data                      |
| /employee      | POST     | JSON | `{"Name": "Rachel", "Salary": "1200", "Age" : "23"}`         | Insert new employeerecord into database         |
| /employee      | PUT/{id} | JSON | `{"Name": "Rachel", "Salary": "1200", "Age" : "23", "Id" : "56"}` | Update customer record into database            |
| /employee      | DELETE   | JSON | `{"Id" : 59}`                                                | Delete particular employee record from database |

