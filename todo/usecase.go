package todo

import "log"

type TodoUsecase interface {
	View(id int) (*Todo, error)
	Add() error
}

type todoUsecase struct {
	repository TodoRepository
}

func NewTodoUsecase(tr TodoRepository) TodoUsecase {
	return &todoUsecase{repository: tr}
}

func (tu *todoUsecase) View(id int) (*Todo, error) {
	todo, err := tu.repository.View(id)
	if err != nil {
		log.Fatal(err)
	}
	return todo, nil
}

func (tu *todoUsecase) Add() error {
	return nil
}
