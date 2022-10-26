package models

type Todo struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Text  string `json:"text"`
  Status string `json:"status"`
}

type CreateTodoInput struct {
	Text  string `json:"text" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type UpdateTodoInput struct  {
	Text  string `json:"text"`
	Status string `json:"status"`  
}

