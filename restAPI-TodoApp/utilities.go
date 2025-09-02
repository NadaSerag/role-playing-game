package main

import "github.com/gin-gonic/gin"

type todo struct {
	id        int
	title     string
	completed bool
}

var Todos []todo

func InitializeTodoArray() []todo {
	//The in-memory store (slice) is initialized as empty.
	Todos = []todo{}
	return Todos
}

// - 200: Successful GET, POST, PUT, or DELETE.
// - 400: Invalid JSON or empty title.
// - 404: Todo not found for GET, PUT, or DELETE.

//When a request comes in, Gin creates a *gin.Context object
// and passes it into the function as the parameter
// the parameter is by convention named "c".
func GetTodos(c *gin.Context) {
	//200
	c.JSON(200, Todos)
}

