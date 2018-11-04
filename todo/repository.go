package todo

import (
	"fmt"
)

type TodoRepository interface {
	View(id int) (*Todo, error)
	Add(summary, comment, due string, priority Priority) error
}

type todoRepository struct {
	Conn string
}

func NewTodoRepository(conn string) TodoRepository {
	return &todoRepository{Conn: conn}
}

func (tr *todoRepository) View(id int) (*Todo, error) {
	todo := &Todo{
		Summary:  "summary",
		Comment:  "test comment",
		Due:      "2019/01/01",
		Priority: High,
	}
	return todo, nil
}

func (tr *todoRepository) Add(summary, comment, due string, priority Priority) error {
	fmt.Printf("aaa")
	return nil
}
