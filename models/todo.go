package models

import (
	// "gorm.io/gorm"
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type Todos []*Todo

type CreateTodoInput struct {
	Text   string `json:"text" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type UpdateTodoInput struct {
	Text   string `json:"text" example:"Todo 1"`
	Status string `json:"status" example:"Completed"`
}

type PutTodoInput struct {
	Text string `json:"text" example:"Todo 1"`
}

type PatchTodoInput struct {
	Status string `json:"status" example:"COMPLETED"`
}

type TodoPatch struct {
	ID     uint   `json:"id" example:"1"`
	Status string `json:"status" example:"COMPLETED"`
}

func (f *UpdateTodoInput) ToModel() (*Todo, error) {
	return &Todo{
		Text:   f.Text,
		Status: f.Status,
	}, nil
}

func (f *CreateTodoInput) ToModel() (*Todo, error) {
	return &Todo{
		Text:   f.Text,
		Status: f.Status,
	}, nil
}

func (f *PutTodoInput) ToModel() (*Todo, error) {
	return &Todo{
		Text: f.Text,
	}, nil
}

func (f *PatchTodoInput) ToPatchModel() (*Todo, error) {
	return &Todo{
		Status: f.Status,
	}, nil
}

func (b Todo) ToDto() *Todo {
	return &Todo{
		ID:        b.ID,
		Status:    b.Status,
		Text:      b.Text,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

type TodoFilterList struct {
	Status string `json:"status" form:"status" example:"COMPLETED"`
}
