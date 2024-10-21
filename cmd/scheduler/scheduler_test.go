package scheduler

import (
	"testing"
	"time"
)

func timeConsumingTask(params ...int) {
	time.Sleep(100 * time.Millisecond)
}

func TestCreateScheduler(t *testing.T) {
	scheduler := Create(5)

	if scheduler.NumberOfTasks != 5 {
		t.Errorf("Expected 5, got %d", scheduler.NumberOfTasks)
	}

	if scheduler.ActiveTasks != 0 {
		t.Errorf("Expected 0, got %d", scheduler.ActiveTasks)
	}

	if scheduler.queue == nil {
		t.Error("Expected a queue to be initialized, but got nil")
	}
}

func TestSchedulerRun(t *testing.T) {
	scheduler := Create(2)

	result := scheduler.Run(timeConsumingTask)
	if result != 1 {
		t.Errorf("Expected Run to return 1, got %d", result)
	}

	if scheduler.GetRunningTasks() != 1 {
		t.Errorf("Expected GetRunningTasks to return 1, got %d", scheduler.GetRunningTasks())
	}

	scheduler.Run(timeConsumingTask)
	if scheduler.GetRunningTasks() != 2 {
		t.Errorf("Expected GetRunningTasks to return 2, got %d", scheduler.GetRunningTasks())
	}

	result = scheduler.Run(timeConsumingTask)
	if result != 0 {
		t.Errorf("Expected task to be queued, got %d", result)
	}

	if scheduler.queue.GetSize() == 0 {
		t.Errorf("Expected task to be queued, but the queue is empty")
	}
}

func TestSchedulerExecuteQueuedTasks(t *testing.T) {
	scheduler := Create(2)

	scheduler.Run(timeConsumingTask)
	scheduler.Run(timeConsumingTask)

	scheduler.Run(timeConsumingTask)

	time.Sleep(300 * time.Millisecond)

	if scheduler.GetRunningTasks() != 0 {
		t.Errorf("Expected all tasks to complete, but running tasks is %d", scheduler.GetRunningTasks())
	}

	if !(scheduler.queue.GetSize() == 0) {
		t.Errorf("Expected the queue to be empty, but it is not")
	}
}
