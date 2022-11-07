package services

import (
	"fmt"
	"errors"
	 "todo/models"
	 "todo/repos"
)


type ITodoService interface {
	List(filters models.TodoFilterList) (models.Todos,error)
	Add(form *models.CreateTodoInput) (*models.Todo, error)
	Get(id int)(*models.Todo, error)
	Put(form *models.PutTodoInput, id int) (*models.Todo, error)
	Patch(form *models.PatchTodoInput, id int) (*models.Todo, error)
	Delete(id int) (*models.Todo, error)
	// Update(form *models.UpdateTodoInput, id int) (*models.Todo, error)
}

type Service struct {
	TodoRepo   repos.ITodoRepo
}

type ServiceConfig struct {
	TodoRepo   repos.ITodoRepo
}

func NewService(c ServiceConfig) ITodoService {
	return &Service{
		TodoRepo:   c.TodoRepo,
	}
}

func (s *Service) Add(form *models.CreateTodoInput) (*models.Todo, error) {
	fmt.Println(&form)
	if(form.Text == "") {
		return nil, errors.New("empty todo")
	}

	if(form.Status == ""){
		form.Status="INCOMPLETE"
	}

	todoModel, _ := form.ToModel()

	todo, err := s.TodoRepo.Add(todoModel)

	if err != nil {
		return nil, err
	}

	todoDto := todo.ToDto()

	return todoDto, nil
	
}


func (s *Service) Get(id int) (*models.Todo, error) {
	todo, err := s.TodoRepo.Get(id)

	if err != nil {
		return nil,  err
	}

	 todoDto := todo.ToDto()

	return todoDto, err
}

func (s *Service) Put(form *models.PutTodoInput, id int) (*models.Todo, error) {

	_, err :=s.TodoRepo.Get(id)

	if(err!= nil ){
		return nil, err
	}
	fmt.Println(form)
	if(form.Text == "") {
		return nil, errors.New("nothing to update")
	}
	todo, err := s.TodoRepo.Put(form,id)

	if err != nil {
		return nil, err
	}

	todoDto := todo.ToDto()

	return todoDto, nil
	
}

func (s *Service) Patch(form *models.PatchTodoInput, id int) (*models.Todo, error) {

	_, err :=s.TodoRepo.Get(id)

	if(err!= nil ){
		return nil, err
	}
	if(form.Status == "") {
		return nil, errors.New("nothing to update")
	}
	todo, err := s.TodoRepo.Patch(form,id)

	if err != nil {
		return nil, err
	}

	todoDto := todo.ToDto()

	return todoDto, nil
	
}

func (s *Service) Delete(id int) (*models.Todo, error) {
	todo, err := s.TodoRepo.Delete(id)
	fmt.Println(*todo)
	if err != nil {
		return nil, err
	}

	todoDto := todo.ToDto()

	return todoDto, nil
}

func (s *Service) List(filters models.TodoFilterList) (models.Todos, error) {
	todos, err := s.TodoRepo.List(filters)

	if err != nil {
		return nil,  err
	}


	return todos, err
}