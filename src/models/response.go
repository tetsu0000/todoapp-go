package models

type Todos []*Todo

type GetTodosResponseBody struct {
	Message string `json:"message"`
	Todos   `json:"data"`
}

type PostTodoResponseBody struct {
	Message string `json:"message"`
	*Todo   `json:"data"`
}

type PutTodoResponseBody struct {
	Message string `json:"message"`
	*Todo   `json:"data"`
}

type DeleteTodoResponseBody struct {
	Message string `json:"message"`
}

func NewGetTodoResponseBody(message string, todos *Todos) *GetTodosResponseBody {
	r := new(GetTodosResponseBody)
	r.Message = message
	r.Todos = *todos
	return r
}

func NewPostTodoResponseBody(message string, todo *Todo) *PostTodoResponseBody {
	r := new(PostTodoResponseBody)
	r.Message = message
	r.Todo = todo
	return r
}

func NewPutTodoResponseBody(message string, todo *Todo) *PutTodoResponseBody {
	r := new(PutTodoResponseBody)
	r.Message = message
	r.Todo = todo
	return r
}

func NewDeleteTodoResponseBody(message string) *DeleteTodoResponseBody {
	r := new(DeleteTodoResponseBody)
	r.Message = message
	return r
}
