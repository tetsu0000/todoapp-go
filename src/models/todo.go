package models

type Todo struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func NewTodo(id string, message string) *Todo {
	t := new(Todo)
	t.Id = id
	t.Message = message
	return t
}
