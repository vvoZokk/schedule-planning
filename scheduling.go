package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"schedule-planning/schedule"
	"time"
)

func PrintTaskList(s *schedule.Schedule) string {
	var bufer bytes.Buffer
	tasks := s.Tasks()
	bufer.WriteString("# task\tstart\tfinish\tearly\tlate\n")
	for _, t := range tasks {
		bufer.WriteString(fmt.Sprintf("%s\t%5d\t%5d\t%5d\t%5d\n", t.Remark, t.Start, t.Start+t.Duration, t.Earliest, t.Latest))
	}
	return bufer.String()
}

func PrintCP(s *schedule.Schedule, cp []int) string {
	width := 80
	var bufer bytes.Buffer
	for i := 0; i < width; i++ {
		bufer.WriteString("_")
	}
	bufer.WriteString("\n")

	for elem, t := range s.Tasks() {
		flag := false
		for _, j := range cp {
			if j == elem {
				flag = true
				break
			}
		}
		for i := 0; i < t.Earliest; i++ {
			bufer.WriteString(" ")
		}

		for i := t.Earliest; i < t.Start; i++ {
			if i == t.Earliest {
				bufer.WriteString("|")
			} else {
				bufer.WriteString("-")
			}
		}
		bufer.WriteString(fmt.Sprintf("[%s", t.Remark))
		for i := 0; i < t.Duration-7; i++ {
			if flag {
				bufer.WriteString("#")
			} else {
				bufer.WriteString(":")
			}
		}
		bufer.WriteString("]")
		for i := t.Start; i < t.Latest; i++ {
			if i == t.Latest-1 {
				bufer.WriteString("|")
			} else {
				bufer.WriteString("-")
			}
		}
		bufer.WriteString("\n")
	}
	for i := 0; i < width; i++ {
		bufer.WriteString("_")
	}
	bufer.WriteString("\n")

	return bufer.String()
}

func main() {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	outFile := os.Stdout
	defer outFile.Close()

	outputFlag := flag.String("o", "", "write task list to file")
	flag.Parse()

	scheduling := schedule.New()

	for i := 0; i < 10; i++ {
		tp := 0
		if rand.Float32() > 0.6 {
			tp = 1
		}
		scheduling.CreateTask(tp, fmt.Sprintf("Task_%d", i), rand.Int()%8+10)
	}
	scheduling.CreateLink(0, 1)
	scheduling.CreateLink(0, 4)
	scheduling.CreateLink(1, 3)
	scheduling.CreateLink(1, 7)
	scheduling.CreateLink(2, 4)
	scheduling.CreateLink(2, 7)
	scheduling.CreateLink(3, 5)
	scheduling.CreateLink(4, 9)
	scheduling.CreateLink(5, 8)
	scheduling.CreateLink(6, 7)
	scheduling.CreateLink(7, 8)
	// find and print critical path
	cp := scheduling.CalculateCP()
	fmt.Println(PrintCP(scheduling, cp))

	if *outputFlag != "" {
		if file, err := os.Create(*outputFlag); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			outFile = file
			writer := bufio.NewWriter(outFile)
			writer.WriteString(PrintTaskList(scheduling))

			if err := writer.Flush(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
