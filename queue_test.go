package searches

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[Node]int{
		Node{Tile{0, 0, "@1"}}: 3, Node{Tile{1, 1, "  "}}: 2, Node{Tile{1, 0, "  "}}: 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: float64(priority),
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    Node{Tile{0, 1, "@2"}},
		priority: float64(1),
	}
	heap.Push(&pq, item)

	expectedItems := []Node{
		Node{Tile{0, 1, "@2"}},
		Node{Tile{1, 1, "  "}},
		Node{Tile{0, 0, "@1"}},
		Node{Tile{1, 0, "  "}},
	}

	// Take the items out; they arrive in decreasing priority order.
	index := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item).value
		if item != expectedItems[index] {
			t.Error("Failed to get item in the right order.")
		}
		fmt.Println(item)
		index++
	}
}
