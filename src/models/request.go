package models

import "encoding/json"

type PostTodoRequest struct {
	Message string `json:"message"`
}

type PutTodoRequest struct {
	Message string `json:"message"`
}

func NewPostTodoRequestFromString(request string) *PostTodoRequest {
	r := new(PostTodoRequest)
	json.Unmarshal([]byte(request), &r)
	return r
}

func NewPutTodoRequestFromString(request string) *PutTodoRequest {
	r := new(PutTodoRequest)
	json.Unmarshal([]byte(request), &r)
	return r
}
