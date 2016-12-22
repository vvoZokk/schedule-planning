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
	TT       int    // Type of task
	Remark   string // Comment
	Duration int    // Duration of task
	Start    int    // Time to start
	Earliest int    // Earliest time to start
	Latest   int    // Latest time to start
}

// NewTask returns new task by number, type, remark and duration
// with zero earliest and latest start time.
func NewTask(t int, r string, d int) *Task {
	if t != Asap {
		t = Alap
	}
	return &Task{t, r, d, 0, 0, 0}
}

// String returns short information about task.
func (t *Task) String() string {
	tT := "ASAP"
	if t.TT != Asap {
		tT = "ALAP"
	}
	return fmt.Sprintf("%s %s: %.2f", t.Remark, tT, t.Duration)
}
