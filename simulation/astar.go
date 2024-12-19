package simulation

import (
	"container/heap"
	"math"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
)

/*
	Code copyright-free modified from
	https://github.com/beefsack/go-astar/blob/master/
*/

type Node struct {
	Coordinates
	Cost     int
	Priority int
	Parent   *Node
	Index    int
}

// Priority queue for Nodes
type PriorityQueue []*Node

// Give length of PriorityQueue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Compare priorities between nodes
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap nodes
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push a new node
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}

// Remove the last node
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1
	*pq = old[0 : n-1]
	return node
}

// Manhattan distance between two points (Heuristic for our a*)
func Heuristic(a, b Coordinates) int {
	// Calculate horizontal distance with wrapping
	dx := int(math.Abs(float64(a.X() - b.X())))
	wrappedDx := constants.GridWidth - dx
	minDx := int(math.Min(float64(dx), float64(wrappedDx)))

	// Vertical distance remains the same
	dy := int(math.Abs(float64(a.Y() - b.Y())))

	return minDx + dy
}

// GetNeighbors retrieves neighbors for the current Coordinates
func GetNeighbors(p Coordinates, width, height int) []Coordinates {
	neighbors := []Coordinates{
		{p.X() - 1, p.Y()}, // Left
		{p.X() + 1, p.Y()}, // Right
		{p.X(), p.Y() - 1}, // Up
		{p.X(), p.Y() + 1}, // Down
	}

	// Make the world round horizontally
	if p.X() == 0 {
		neighbors = append(neighbors, Coordinates{width - 1, p.Y()})
	} else if p.X() == width-1 {
		neighbors = append(neighbors, Coordinates{0, p.Y()})
	}

	// Filter valid neighbors (does it go beyond the map or not)
	validNeighbors := []Coordinates{}
	for _, n := range neighbors {
		if n.X() >= 0 && n.X() < width && n.Y() >= 0 && n.Y() < height {
			validNeighbors = append(validNeighbors, n)
		}
	}
	return validNeighbors
}

// A* algorithm
func AStar(start, goal Coordinates) []Coordinates {
	var grid [54][96]int = constants.MapGrid

	width := len(grid[0])
	height := len(grid)

	// Priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Maps to track visited nodes and costs
	visited := make(map[Coordinates]bool)
	costs := make(map[Coordinates]int)

	// Start node
	startNode := &Node{
		Coordinates: start,
		Cost:        0,
		Priority:    Heuristic(start, goal),
		Parent:      nil,
	}
	heap.Push(pq, startNode)
	costs[start] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		// Check if we've reached the goal
		if current.Coordinates == goal {
			// Reconstruct the path
			path := []Coordinates{}
			for c := current; c != nil; c = c.Parent {
				if c.Coordinates != start {
					path = append([]Coordinates{c.Coordinates}, path...)
				}
			}
			return path
		}

		visited[current.Coordinates] = true

		// Process neighbors
		for _, neighbor := range GetNeighbors(current.Coordinates, width, height) {
			if visited[neighbor] || grid[neighbor.Y()][neighbor.X()] == 1 {
				continue
			}

			newCost := costs[current.Coordinates] + 1
			if cost, ok := costs[neighbor]; !ok || newCost < cost {
				costs[neighbor] = newCost
				priority := newCost + Heuristic(neighbor, goal)
				heap.Push(pq, &Node{
					Coordinates: neighbor,
					Cost:        newCost,
					Priority:    priority,
					Parent:      current,
				})
			}
		}
	}

	// Return an empty path if no path is found
	return []Coordinates{}
}
