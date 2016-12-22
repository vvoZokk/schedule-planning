package schedule

import (
	"errors"
	"fmt"
)

// Link between tasks.
type Link struct {
	previous, next int
}

// Schedule with task list, list of links and critical path.
type Schedule struct {
	tasks []Task
	links []Link
	cp    []int
}

// New returns new schedule with empty lists.
func New() *Schedule {
	return &Schedule{make([]Task, 0), make([]Link, 0), []int{}}
}

func previous(l Link) int {
	return l.previous
}

func next(l Link) int {
	return l.next
}

// CreateTask creates new task and returns its number in list
// and nil or error.
func (s *Schedule) CreateTask(t int, r string, d float32) (int, error) {
	if d < 0 {
		err := errors.New(fmt.Sprintf("Incorrect duration %f", d))
		return len(s.tasks) - 1, err
	} else {
		task := *NewTask(t, r, d)
		s.tasks = append(s.tasks, task)
	}
	return len(s.tasks) - 1, nil
}

// CreateLink creates new link between existing tasks and
// returns nil or error.
func (s *Schedule) CreateLink(p, n int) error {
	if p < 0 && n < 0 {
		err := errors.New(fmt.Sprintf("Incorrect link: previous %d, next %d", p, n))
		return err
	} else if p >= len(s.tasks) {
		err := errors.New(fmt.Sprintf("Incorrect link: previos task #%d is not exist", p))
		return err
	} else if n >= len(s.tasks) {
		err := errors.New(fmt.Sprintf("Incorrect link: next task #%d is not exist", n))
		return err
	} else {
		s.links = append(s.links, Link{p, n})
	}
	return nil
}
