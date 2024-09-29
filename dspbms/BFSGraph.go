package main

import (
	"fmt"
)

// Graph represents a graph using an adjacency list
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// AddEdge adds an edge between two vertices in an undirected graph
func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v) // For undirected graph
}

// BFS performs Breadth-First Search starting from a given source vertex
func (g *Graph) BFS(start int) {
	// Initialize a queue to process nodes in breadth-first order
	queue := []int{start}

	// Create a visited map to keep track of visited nodes
	visited := make(map[int]bool)
	visited[start] = true

	for len(queue) > 0 {
		// Dequeue the front element from the queue
		current := queue[0]
		queue = queue[1:]

		// Visit the current node
		fmt.Printf("%d ", current)

		// Get all adjacent vertices of the current node
		for _, adjacent := range g.vertices[current] {
			// If the adjacent node hasn't been visited yet, enqueue it
			if !visited[adjacent] {
				queue = append(queue, adjacent)
				visited[adjacent] = true
			}
		}
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// Perform BFS traversal starting from vertex 0
	fmt.Println("BFS traversal starting from vertex 0:")
	graph.BFS(0) // Output: 0 1 2 3 4
}
