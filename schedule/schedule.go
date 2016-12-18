package schedule

type Link struct {
	Before, After int
}

type Schedule struct {
	taskList []Task
	links    []Link
}

func New() *Schedule {
	return &Schedule{make([]Task, 0, 10), make([]Link, 0, 10)}
}
