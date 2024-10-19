**Basic objectives this solves currently**
- Makes number of functions run concurrently
- If more then `N` functions are scheduled for running, it adds them into queue
- Enqueues the functions from a queue if `Number of running functions = (Total tasks allowed to run) - 2`

Basic methods:

1. Create a scheduler
  ```
  scheduler := scheduler.Create(10)  // where 10 -> number of tasks allowed to run concurrently
  ```

2. Put a function into scheduler for running
   ```
   scheduler.Run(func(params ...int) {
			fmt.Println("Function I want to run")
		}, 1,2,3) // here 1,2,3 are the parameters which will be passed to the function you have passed
   ```
   
3. Get number of currently active tasks(ones which are running in scheduler)

   ```
   scheduler.GetRunningTasks()
   ```
