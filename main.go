package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Saumya40-codes/scheduler/cmd/scheduler"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Creating a scheduler with a limit of 10 tasks")
	sched := scheduler.Create(10)

	for i := 0; i < 15; i++ {
		wg.Add(1)
		taskNum := i
		sched.Run(func(params ...int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d is executed\n", taskNum)
		}, taskNum)

		fmt.Println("=====================================")
		fmt.Printf("Number of Tasks currently running: %d\n", sched.GetRunningTasks())
		fmt.Println("=====================================")
	}

	wg.Wait()
}

