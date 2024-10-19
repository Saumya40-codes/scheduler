package scheduler

import "sync"

type Queue struct {
	size  int
	queue [](func(...int))
	mu    sync.Mutex
}

const defaultQueueSize = 10

// CreateQueue initializes a new Queue with a given size, or use 10 as the default size is size is less than or equal to 0
func CreateQueue(size int) *Queue {
	if size <= 0 {
		size = defaultQueueSize
	}

	return &Queue{
		size:  size,
		queue: make([](func(...int)), 0, size),
	}
}

// Enqueue adds a task to the queue
func (q *Queue) Enqueue(task func(params ...int)) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) >= q.size {
		return false // Queue is full
	}

	q.queue = append(q.queue, task)
	return true
}

// Dequeue removes and returns the task at the front of the queue
func (q *Queue) Dequeue() (func(...int), bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return nil, false // Queue is empty
	}

	task := q.queue[0]
	q.queue = q.queue[1:]
	return task, true
}

// GetTasks returns all tasks in the queue
func (q *Queue) GetTasks() [](func(...int)) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queue
}

// GetSize returns the current number of tasks in the queue
func (q *Queue) GetSize() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.queue)
}

// Peek returns the task at the front of the queue without removing it
func (q *Queue) Peek() (func(...int), bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return nil, false // Queue is empty
	}

	return q.queue[0], true
}

