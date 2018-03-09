package queue

// Queue implements a FIFO queue
type Queue struct {
	list []interface{}
}

// Push pushes to the end of the queue
func (q *Queue) Push(item interface{}) {
	q.list = append(q.list, item)
}

// Pop pops from the begining of the queue
func (q *Queue) Pop() (item interface{}) {
	if len(q.list) == 0 {
		return nil
	}

	item, q.list = q.list[0], q.list[1:]

	return item
}

// Empty returns if the Queue is empty
func (q *Queue) Empty() bool {
	return len(q.list) == 0
}
