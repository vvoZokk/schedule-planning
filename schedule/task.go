package schedule

import (
	"fmt"
)

const (
	Asap = iota // as soon as possible
	Alap        // as late as possible
)

// Taks type includes type (ASAP or ALAP), remark, start time,
// duration, earliest and latest start time.
type Task struct {
	t    int
	r    string
	s, d float32
	e, l float32
}

// NewTask returns new task by number, type, remark and duration
// with zero earliest and latest start time.
func NewTask(t int, r string, d float32) *Task {
	if t != Asap {
		t = Alap
	}
	return &Task{t, r, 0.0, d, 0.0, 0.0}
}

// String returns short information about task.
func (t *Task) String() string {
	tT := "ASAP"
	if t.t != Asap {
		tT = "ALAP"
	}
	return fmt.Sprintf("Task %s: %.2f", tT, t.d)
}

func (t *Task) SetStartTime(s float32) {
	t.s = s
}

func (t *Task) SetEarliest(e float32) {
	t.e = e
}

func (t *Task) SetLatest(l float32) {
	t.l = l
}

func (t *Task) Type() int {
	return t.t
}

func (t *Task) Duration() float32 {
	return t.d
}

func (t *Task) StartTime() float32 {
	return t.s
}

func (t *Task) Earliest() float32 {
	return t.e
}

func (t *Task) Latest() float32 {
	return t.l
}

func (t *Task) Remark() string {
	return t.r
}
