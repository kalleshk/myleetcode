package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph (undirected graph)
func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v) // For undirected graph, add both v->w and w->v
}

// DFS Recursive Traversal
func (g *Graph) DFS_Recursive(v int, visited map[int]bool) {
	// Mark the current node as visited and print it
	visited[v] = true
	fmt.Printf("%d ", v)

	// Recur for all the vertices adjacent to this vertex
	for _, adjacent := range g.vertices[v] {
		if !visited[adjacent] {
			g.DFS_Recursive(adjacent, visited)
		}
	}
}

// DFS Iterative Traversal (using stack)
func (g *Graph) DFS_Iterative(v int) {
	// Stack for DFS
	stack := []int{v}
	visited := make(map[int]bool)

	// Push the root node into the stack
	visited[v] = true

	for len(stack) > 0 {
		// Pop a vertex from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", current)

		// Get all adjacent vertices, if not visited, push them to the stack
		for _, adjacent := range g.vertices[current] {
			if !visited[adjacent] {
				stack = append(stack, adjacent)
				visited[adjacent] = true
			}
		}
	}
}

func main() {
	// Create a graph
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// Perform DFS traversal (Recursive)
	fmt.Println("DFS Recursive starting from vertex 0:")
	visited := make(map[int]bool)
	graph.DFS_Recursive(0, visited) // Output: 0 1 2 3 4
	fmt.Println()

	// Perform DFS traversal (Iterative)
	fmt.Println("DFS Iterative starting from vertex 0:")
	graph.DFS_Iterative(0) // Output: 0 2 3 4 1
	fmt.Println()
}
