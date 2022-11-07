package repos

import (
	// "log"
	 "todo/models"
	"gorm.io/gorm"
)

var idQuery = "id = ?"

type TodoRepo struct {
	DB *gorm.DB
}


func NewTodoRepo(db *gorm.DB) TodoRepo {
	return TodoRepo {
		DB: db,
	}
}

type ITodoRepo interface {
	List( filters models.TodoFilterList) (models.Todos, error)
	Add(todo *models.Todo) (*models.Todo, error)
	Get(id int) (*models.Todo, error) 
	Put(form *models.PutTodoInput, id int) (*models.Todo, error) 
	Patch(form *models.PatchTodoInput, id int) (*models.Todo, error)
	 Delete(id int) (*models.Todo, error)
}


func (r TodoRepo) List(filters models.TodoFilterList) (models.Todos,  error) {
	todos := make([]*models.Todo, 0)

	db := models.DB.Find(&todos)

	if len(filters.Status) > 0 {
		db = db.Where("status = ?", filters.Status)
	}


	err := db.Find(&todos).Error

	if err != nil {
		return nil, err
	}

	if len(todos) == 0 {
		return nil, nil
	}

	return todos, nil
}

func (r TodoRepo) Add(todo *models.Todo) (*models.Todo, error) {

	if err := r.DB.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepo) Get(id int) (*models.Todo, error) {
	todo := new(models.Todo)
	err := models.DB.Where("id = ?",id).First(&todo).Error

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepo) Put(form *models.PutTodoInput, id int) (*models.Todo, error) {
	todo, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where(idQuery, id).Updates(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}


func (r TodoRepo) Patch(form *models.PatchTodoInput, id int) (*models.Todo, error) {
	todo, err := form.ToPatchModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where(idQuery, id).Updates(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r TodoRepo) Delete(id int) (*models.Todo, error) {
	todo := new(models.Todo)
	if err := r.DB.Where(idQuery, id).Delete(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

