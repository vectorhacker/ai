package search

import "github.com/vectorhacker/ai/queue"

// Successor is triple of state, action, and cost
type Successor struct {
	State  interface{}
	Action interface{}
	Cost   int
}

// Problem represents a search problem
type Problem interface {
	GoalTest(state interface{}) bool
	Sucessors(state interface{}) []Successor
	InitialState() interface{}
}

// Node represents a leaf node
type Node struct {
	Path  []interface{}
	Cost  int
	State interface{}
}

// Frontier is anything that can be pushed, poped,
// or be empty
type Frontier interface {
	Push(interface{})
	Pop() interface{}
	Empty() bool
}

func genericSearch(problem Problem, frontier Frontier) *Node {
	explored := map[interface{}]struct{}{}

	for !frontier.Empty() {
		node := frontier.Pop().(*Node)
		explored[node.State] = struct{}{}

		if problem.GoalTest(node.State) {
			return node
		}

		for _, successor := range problem.Sucessors(node.State) {
			if _, ok := explored[successor.State]; !ok {
				cost := node.Cost + successor.Cost
				child := &Node{
					State: successor.State,
					Path:  append(node.Path, successor.Action),
					Cost:  cost,
				}

				frontier.Push(child)
			}
		}
	}

	return nil
}

type Search func(problem Problem) *Node

type HeuristicFunc func(problem Problem, state interface{}) int

func DepthFirst(problem Problem) *Node {
	return genericSearch(problem, &queue.Stack{})
}

func BreadthFirst(problem Problem) *Node {
	return genericSearch(problem, &queue.Queue{})
}

func UniformCost(problem Problem) *Node {
	return genericSearch(problem, queue.NewPriorityQueue(func(item interface{}) int {
		node := item.(*Node)
		return node.Cost
	}))
}

func Greedy(problem Problem, heuristic HeuristicFunc) *Node {

	return genericSearch(problem, queue.NewPriorityQueue(func(item interface{}) int {
		node := item.(*Node)
		return heuristic(problem, node.State)
	}))
}

// AStar creates an AStar search with a heuristic
func AStar(heuristic HeuristicFunc) Search {
	return func(problem Problem) *Node {
		return genericSearch(problem, queue.NewPriorityQueue(func(item interface{}) int {
			node := item.(*Node)
			return node.Cost + heuristic(problem, node.State)
		}))
	}
}
