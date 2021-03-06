// Package schedule implements the critical path method (CPM) for scheduling.
package schedule

import (
	"errors"
	"fmt"
)

// Link between tasks.
type Link struct {
	previous, next int
}

// Schedule with slice of tasks and slice of links.
type Schedule struct {
	tasks []Task
	links []Link
}

// New returns new schedule with empty slices.
func New() *Schedule {
	return &Schedule{make([]Task, 0), make([]Link, 0)}
}

func previous(l Link) int {
	return l.previous
}

func next(l Link) int {
	return l.next
}

// CreateTask creates new task and returns length of task list
// and nil or error.
func (s *Schedule) CreateTask(t int, r string, d int) (int, error) {
	if d < 0 {
		err := errors.New(fmt.Sprintf("Incorrect duration: %f", d))
		return len(s.tasks), err
	} else {
		task := *NewTask(t, r, d)
		s.tasks = append(s.tasks, task)
	}
	return len(s.tasks), nil
}

// CreateLink creates new link between existing tasks and
// returns nil or error.
func (s *Schedule) CreateLink(p, n int) error {
	if p < 0 || n < 0 || p == n {
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

func (s *Schedule) findRoot(elem func(Link) int) []int {
	result := make([]int, 0)
	for i, _ := range s.tasks {
		flag := true
		for _, link := range s.links {
			if i == elem(link) {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}

func (s *Schedule) findRelative(e, r func(Link) int) map[int][]int {
	result := make(map[int][]int, 0)
	for i, _ := range s.tasks {
		result[i] = make([]int, 0)
	}
	for _, link := range s.links {
		if slice, ok := result[e(link)]; ok {
			slice = append(slice, r(link))
			result[e(link)] = slice
		} else {
			result[e(link)] = []int{r(link)}
		}
	}
	return result
}

func (s *Schedule) findEarliest(root []int, m map[int][]int) {
	for i := 0; i < len(root); i++ {
		elem := root[i]
		root = append(root, m[elem]...)
		for _, child := range m[elem] {
			time := s.tasks[elem].Earliest + s.tasks[elem].Duration
			if time > s.tasks[child].Earliest {
				s.tasks[child].Earliest = time
			}
		}
	}
}

func (s *Schedule) findLatest(root []int, m map[int][]int) {
	// find max time
	var max int = 0
	for _, elem := range root {
		if time := s.tasks[elem].Earliest + s.tasks[elem].Duration; max < time {
			max = time
		}
	}
	for i, _ := range s.tasks {
		s.tasks[i].Latest = max - s.tasks[i].Duration
	}
	// find latest time for each task
	for i := 0; i < len(root); i++ {
		elem := root[i]
		root = append(root, m[elem]...)
		for _, parent := range m[elem] {
			time := s.tasks[elem].Latest - s.tasks[parent].Duration
			if time < s.tasks[parent].Latest {
				s.tasks[parent].Latest = time
			}
		}
	}
}

// CalculateCP finds critical path for schedule and returns
// slice of numbers of critical path tasks.
func (s *Schedule) CalculateCP() []int {
	// map of cheldren
	children := s.findRelative(previous, next)
	// map of parents
	parents := s.findRelative(next, previous)

	s.findEarliest(s.findRoot(next), children)
	s.findLatest(s.findRoot(previous), parents)

	// find time to start for each task
	root := s.findRoot(next)
	for _, elem := range root {
		if s.tasks[elem].TT == Alap {
			s.tasks[elem].Start = s.tasks[elem].Latest
		}
	}
	root = s.findRoot(previous)
	for i := 0; i < len(root); i++ {
		elem := root[i]
		root = append(root, parents[elem]...)
		for _, parent := range parents[elem] {
			if s.tasks[elem].TT == Asap {
				time := s.tasks[parent].Latest + s.tasks[parent].Duration
				if s.tasks[parent].TT == Alap && time > s.tasks[elem].Earliest {
					s.tasks[elem].Start = time
				} else {
					if s.tasks[elem].Start == 0 {
						s.tasks[elem].Start = s.tasks[elem].Earliest
					}
				}
			} else {
				s.tasks[elem].Start = s.tasks[elem].Latest
			}
		}
	}
	// find critical path
	cp := make([]int, 0)
	for i, t := range s.tasks {
		if t.Earliest == t.Latest {
			cp = append(cp, i)
		}
	}
	return cp
}

// Tasks return slice of tasks.
func (s *Schedule) Tasks() []Task {
	return s.tasks
}
