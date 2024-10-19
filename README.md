**Basic objectives this solves currently**
- Makes number of functions run concurrently
- If more then `N` functions are scheduled for running, it adds them into queue
- Enqueues the functions from a queue if `Number of running functions = (Total tasks allowed to run) - 2`
