package queue

import "container/heap"

type priorityItem struct {
	Item     interface{}
	Priority int
}

type priorityHeap []*priorityItem

func (h priorityHeap) Len() int { return len(h) }

func (h priorityHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h priorityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *priorityHeap) Push(i interface{}) {
	item := i.(*priorityItem)
	*h = append(*h, item)
}

func (h *priorityHeap) Pop() (item interface{}) {
	item, *h = (*h)[h.Len()], (*h)[:h.Len()-1]

	return item
}

type PriorityFunc func(item interface{}) int

// PriorityQueue implements a priority queue
type PriorityQueue struct {
	heap         priorityHeap
	priorityFunc PriorityFunc
}

func NewPriorityQueue(priorityFunc PriorityFunc) *PriorityQueue {
	pq := &PriorityQueue{}

	heap.Init(&pq.heap)

	return pq
}

func (pq *PriorityQueue) Push(item interface{}) {

	heap.Push(&pq.heap, &priorityItem{
		Item:     item,
		Priority: pq.priorityFunc(item),
	})
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(&pq.heap).(*priorityItem).Item
}

// Empty checks wether the queue is empty
func (pq *PriorityQueue) Empty() bool {
	return len(pq.heap) == 0
}
