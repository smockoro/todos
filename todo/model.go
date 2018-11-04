package todo

type Priority int

const (
	High Priority = iota + 1
	Middle
	Low
)

type Todo struct {
	Summary  string
	Comment  string
	Due      string
	Priority Priority
}
