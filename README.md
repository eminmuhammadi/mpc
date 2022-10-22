# Multiparty Computation

An example of a multiparty computation protocol when nodes can broadcast random values to each other.

## Example

```go
// Seed the random number generator
rand.Seed(time.Now().UnixNano())

// Party members
Alice := &Node{ID: "Alice", Val: 4000.00}
Bob := &Node{ID: "Bob", Val: 3000.00}
Charlie := &Node{ID: "Charlie", Val: 2000.00}

// Party
Party := &Party{
	Nodes: []*Node{Alice, Bob, Charlie},
}

// Broadcast random values to each other
Party.SendRandomValues()

// Print the values
for _, node := range Party.Nodes {
	node.Publish()
}

// Get the average value
fmt.Println(Party.Average())
```