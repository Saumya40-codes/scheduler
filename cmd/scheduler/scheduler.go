package scheduler

import "fmt"

type Scheduler struct {
	NumberOfTasks int
	ActiveTasks   int
	queue         *Queue
	status        bool
}

func Create(tasks int) *Scheduler {
	return &Scheduler{
		NumberOfTasks: tasks,
		queue:         CreateQueue(100),
		status:        true,
	}
}

func (s *Scheduler) GetRunningTasks() int {
	return s.ActiveTasks
}

// Run executes a task if the number of active tasks is less than the limit, else queues the task
func (s *Scheduler) Run(task func(params ...int), params ...int) int {
	// Check if adding another task exceeds the limit
	if s.ActiveTasks >= s.NumberOfTasks {
		if ok := s.queue.Enqueue(task); ok {
			fmt.Println("Task queued")
			return 0
		}

		fmt.Println("Task discarded as the queue is full")
		s.status = false
		return -1
	}

	s.ActiveTasks++

	go func() {
		defer func() {
			s.ActiveTasks--
			s.executeQueuedTasks() // if any
		}()

		task(params...)
	}()

	return 1
}

func (s *Scheduler) executeQueuedTasks() {
	// if running tasks are equal to totaltasks-2, then prepare minimum of 2 tasks to run
	if s.ActiveTasks != s.NumberOfTasks-2 {
		return
	}

	for i := 0; i < 2; i++ {
		if task, ok := s.queue.Dequeue(); ok {
			s.Run(task)
		}
	}

	s.status = true
}
