package bultin

import (
	"fmt"
	"sync"
	"testing"
)

// one of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes.

func Test_map(t *testing.T) {
	m1 := make(map[string]int)
	m2 := map[string]int{}
	t.Log(m1)
	t.Log(m2)

	m1["key"] = 1
	t.Log(m1["key"])
	t.Log(len(m1))
	delete(m1, "key")
	i, ok := m1["key"]
	t.Log(i, ok)
	t.Log(len(m1))
}

func Test_traverse(t *testing.T) {
	type Node struct {
		Next  *Node
		Value interface{}
	}
	first := &Node{
		Next: &Node{
			Next:  &Node{},
			Value: "second",
		},
		Value: "first",
	}

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

}

func Test_populated(t *testing.T) {
	type Person struct {
		Name  string
		Likes []string
	}
	// people := []*Person{
	// 	{
	// 		Name:  "DK",
	// 		Likes: []string{"cheese"},
	// 	},
	// 	{
	// 		Name:  "DK",
	// 		Likes: []string{"bacon"},
	// 	},
	// }

	people := make([]*Person, 0, 10)

	people = append(people,
		&Person{
			Name:  "DK",
			Likes: []string{"cheese"},
		},
		&Person{
			Name:  "DK",
			Likes: []string{"bacon"},
		},
	)

	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}

	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese.")
	}

	fmt.Println(len(likes["bacon"]), "people like bacon.")
}

func Test_innerMap(t *testing.T) {
	type Key struct {
		Path, Country string
	}

	hits := make(map[Key]int)

	k1 := Key{
		Path:    "/",
		Country: "vn",
	}
	k2 := Key{
		Path:    "/",
		Country: "vn",
	}

	hits[k1]++
	hits[k2]++
	t.Log(len(hits), hits[k1])

}

func Test_concurrentMmap(t *testing.T) {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)
}

func Test_mapOrder(t *testing.T) {

}
