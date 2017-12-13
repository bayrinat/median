package metric

import (
	"container/heap"
	"fmt"
	"github.com/pkg/errors"
	"median/structs"
)

type Delay interface {
	// Adds delay to the collection, returns error if something goes wrong
	AddDelay(delay int) error
	// Returns the wantMedian value of the collection
	GetMedian() float64
}

// Base implementation of Delay interface
type DelayBase struct {
	// heap for getting wantMedian
	medianHeap *structs.MedianHeap
	// stores elements inside sliding window
	queue    *structs.Queue
	capacity int
}

// Returns DelayBase struct with given capacity
func NewDelay(capacity int) (*DelayBase, error) {
	if capacity <= 0 {
		return nil, errors.Wrap(errNotPositiveCapacity, fmt.Sprintf(", got: %v", capacity))

	}
	delay := DelayBase{
		medianHeap: structs.NewMedianHeap(),
		queue:      &structs.Queue{},
		capacity:   capacity,
	}

	heap.Init(delay.queue)

	return &delay, nil
}

// Adds the new value to the collection, returns err if something goes wrong
func (d *DelayBase) AddDelay(delay int) error {
	// remove element before inserting if collection is full
	if d.queue.Len() >= d.capacity {
		pop := heap.Pop(d.queue)
		d.medianHeap.Remove(pop.(int))
	}

	heap.Push(d.queue, delay)
	d.medianHeap.Add(delay)

	return nil
}

// Returns the wantMedian of the collection
//
// If there are no elements available in the sliding window the answer is 0.
// If only one element available in the sliding window the answer is -1.
// If n is odd then Median (M) = value of ((n + 1)/2)th item from a sorted array of length n.
// If n is even then Median (M) = value of [((n)/2)th item term + ((n)/2 + 1)th item term ] /2
func (d *DelayBase) GetMedian() float64 {
	if d.queue.Len() == 0 {
		return 0
	}

	if d.queue.Len() == 1 {
		return -1
	}

	return d.medianHeap.Median()
}
