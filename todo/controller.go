package todo

import "log"

type TodoController interface {
	View(id int) (*Todo, error)
	Add(summary, comment, due string, priority Priority) error
}

type todoController struct {
	Conn string
}

func NewTodoController(conn string) TodoController {
	return &todoController{
		Conn: conn,
	}
}

func (tc *todoController) View(id int) (*Todo, error) {
	rep := NewTodoRepository(tc.Conn)
	tu := NewTodoUsecase(rep)
	todo, err := tu.View(id)
	if err != nil {
		log.Fatal(err)
	}
	return todo, nil
}

func (tc *todoController) Add(summary, comment, due string, priority Priority) error {
	return nil
}
