package utils

import "fmt"

type Queue[T any] struct {
	items[]T
}

// Enqueue method adds an item at the end of the queue .
func(queue *Queue[T]) Enqueue(item T) {
	queue.items = append(queue.items, item)
}

// Dequeue method removes an item from the queue .
func(queue *Queue[T]) Dequeue() (T, error) {
	if len(queue.items) == 0 {
		var zero T
		return zero, fmt.Errorf("the 'Dequeue' operation failed! The queue is empty")
	}
	item := queue.items[0]
	queue.items = queue.items[1:]
	return item, nil
}

// Peek method returns the item in front of the queue without removing it.
func(queue *Queue[T]) Peek() (T, error) {
	if len(queue.items) == 0 {
		var zero T
		return zero, fmt.Errorf("the 'Peek' opearion failed! the queue is not empty")
	}
	return queue.items[0], nil
}


// PrintLastQueue method prints and returns the item standing at the end of the queue .
func(queue *Queue[T]) PrintLastQueue() (T, error) {
	if len(queue.items) == 0 {
		var zero T
		return zero, fmt.Errorf("the 'PrintLastQueue' operation failed, the queue is empty" )
	}
	endItem := queue.items[len(queue.items)-1]
	return endItem, nil
}

// IsEmpty method checks if the queue is empty or not. returns 'true if it is and 'false' if it is not.
func(queue *Queue[T]) IsEmpty() bool {
	return len(queue.items) == 0
}
