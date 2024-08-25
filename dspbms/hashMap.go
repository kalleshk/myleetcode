package main

import "fmt"

func main() {
    // Create a map
    m := make(map[string]int)

    // Insert elements
    m["apple"] = 5
    m["banana"] = 3
    m["orange"] = 7

    fmt.Println("Initial map:", m)

    // Update an element
    m["apple"] = 10
    fmt.Println("After updating 'apple':", m)

    // Delete an element
    delete(m, "banana")
    fmt.Println("After deleting 'banana':", m)
}
