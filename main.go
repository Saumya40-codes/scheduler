package main

import (
	"fmt"
	"sync"

	"github.com/Saumya40-codes/scheduler/cmd/scheduler"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Creating a scheduler with a limit of 10 tasks")
	sched := scheduler.Create(10)

	for i := 0; i < 15; i++ {
		fmt.Println("=====================================")
		fmt.Printf("Number of Tasks currently running: %d \n", sched.GetRunningTasks())
		fmt.Println("=====================================")
		fmt.Println()
		wg.Add(1)
		taskNum := i
		sched.Run(func(params ...int) {
			defer wg.Done()
			fmt.Println("*************************************")
			fmt.Printf("Task %d is executed\n", taskNum)
			fmt.Println("*************************************")
			fmt.Println()
		}, taskNum)
	}

	wg.Wait()
}

