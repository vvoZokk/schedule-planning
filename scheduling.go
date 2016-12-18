// The critical path method (CPM) for scheduling.
package main

import (
	"fmt"
	"schedule-planning/schedule"
)

func main() {

	taskNumber := 1
	task := schedule.NewTask(1, taskNumber, fmt.Sprintf("Task_%d", taskNumber), 9.5)
	fmt.Printf("Info: %s", task)
}
