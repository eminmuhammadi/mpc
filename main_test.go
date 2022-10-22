package main

import "testing"

func TestAverage(t *testing.T) {
	A := &Node{ID: "A", Val: 400}
	B := &Node{ID: "B", Val: 200}

	party := &Party{Nodes: []*Node{A, B}}

	party.SendRandomValues()

	A.Publish()
	B.Publish()

	average := (A.Val + B.Val) / 2

	if party.Average() != average {
		t.Errorf("Average() = %v, want %v", party.Average(), average)
	}
}
