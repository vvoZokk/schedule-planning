package schedule

import (
	"fmt"
)

const (
	Asp = iota
	Alp
)

// Taks type includes number, type (ASP or ALP), remark, duration,
// earliest and latest start time.
type Task struct {
	n, t    int
	r       string
	d, e, l float32
}

/*
func NewTaskByNumber(n int) *Task {
	return &Task{n, Asp, fmt.Sprint("Task_", n), 0.0, 0.0, 0.0}
}*/

// NewTask returns new task by number, type, remark, duration and
// with zero earliest and latest start time.
func NewTask(n, t int, r string, d float32) *Task {
	if t == Asp {
		t = Alp
	}
	return &Task{n, t, r, d, 0.0, 0.0}
}

// String returns short information about task.
func (t *Task) String() string {
	tT := "ASP"
	if t.t != Asp {
		tT = "ALP"
	}
	return fmt.Sprintf("Task #%d, %s: %f", t.n, tT, t.d)
}

func (t *Task) SetEarliest(e float32) {
	t.e = e
}

func (t *Task) SetLatest(l float32) {
	t.l = l
}

func (t *Task) Number() int {
	return t.n
}

func (t *Task) GetTypeAndDuration() (int, float32) {
	return t.t, t.d
}
