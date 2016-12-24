package schedule

import (
	"testing"
)

type testTask struct {
	t, d   int
	result int
}

func TestCreateTask(t *testing.T) {
	s := New()
	tests := []testTask{
		{0, 1, 1},
		{0, 1, 2},
		{1, 0, 3},
		{1, -1, 3},
	}
	for _, test := range tests {
		length, _ := s.CreateTask(test.t, "", test.d)
		if length != test.result {
			t.Errorf("Expected task list lenght after creating task %d, got %d", test.result, length)
		}
	}
}

func TestCreateLink(t *testing.T) {
	s := New()
	// set incorrect tasks index in links
	tests := [][]int{
		[]int{0, 0},
		[]int{1, -1},
		[]int{-1, -1},
		[]int{100, 1},
		[]int{1, 100},
	}
	for _, test := range tests {
		if err := s.CreateLink(test[0], test[1]); err == nil {
			t.Errorf("Expected error by creating link between %d and %d, got nil", test[0], test[1])
		}
	}
}
