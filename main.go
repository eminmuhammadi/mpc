package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	ID     string
	Val    float64
	Inbox  []float64
	Outbox []float64
	Z      float64
}

type Party struct {
	Nodes  []*Node
	Result float64
}

// Generate random float64 value
func (node *Node) Random() float64 {
	r := rand.Float64() * 1000

	node.Outbox = append(node.Outbox, r)

	return r
}

// Publish z values
func (node *Node) Publish() float64 {
	var z float64
	var inbox float64
	var outbox float64

	for _, val := range node.Inbox {
		inbox += val
	}

	for _, val := range node.Outbox {
		outbox += val
	}

	z = node.Val + (inbox - outbox)
	node.Z = z

	return z
}

// Broadcast random values to all nodes
func (party *Party) SendRandomValues() {
	for _, node := range party.Nodes {
		for _, neighbor := range party.Nodes {
			if node.ID != neighbor.ID {
				r := node.Random()
				neighbor.Inbox = append(neighbor.Inbox, r)
			}
		}
	}
}

// Calculate the average value of all nodes
func (party *Party) Average() float64 {
	var z float64

	for _, node := range party.Nodes {
		z += node.Z
	}

	party.Result = z / float64(len(party.Nodes))

	return party.Result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	Alice := &Node{ID: "Alice", Val: 4000.00}
	Bob := &Node{ID: "Bob", Val: 3000.00}
	Charlie := &Node{ID: "Charlie", Val: 2000.00}

	Party := &Party{
		Nodes: []*Node{Alice, Bob, Charlie},
	}

	Party.SendRandomValues()

	for _, node := range Party.Nodes {
		node.Publish()
	}

	fmt.Println(Party.Average())

	fmt.Println(Alice)
	fmt.Println(Bob)
	fmt.Println(Charlie)
}
