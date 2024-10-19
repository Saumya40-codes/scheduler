package scheduler

import "fmt"

type Scheduler struct {
	NumberOfTasks int
}

func Create(tasks int) *Scheduler {
	return &Scheduler{NumberOfTasks: tasks}
}

func (s *Scheduler) Run() {
	fmt.Println("Running scheduler with", s.NumberOfTasks, "tasks")
}
