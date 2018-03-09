package queue

// Stack implements a LIFO queue
type Stack struct {
	list []interface{}
}

// Push pushes an elemement to the end of the stack
func (s *Stack) Push(item interface{}) {
	s.list = append(s.list, item)
}

// Pop pops an element from the end of the stack
func (s *Stack) Pop() (item interface{}) {
	if s.Empty() {
		return nil
	}

	item, s.list = s.list[len(s.list)-1], s.list[:len(s.list)-1]

	return item
}

// Empty checks wether the stack is empty
func (s *Stack) Empty() bool {
	return len(s.list) == 0
}
