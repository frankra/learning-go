package main

import "testing"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func BuildTree(numbers []int) *Node {
	var root *Node
	for _, num := range numbers {
		if root == nil {
			root = &Node{
				Value: num,
			}
		} else {
			if 
		}
	}
}



func TestTree(t *testing.T) {
	t.Run("should add the elements", func(t *testing.T) {
		values := []int{
			2, 1, 3, 0,
		}

		expectedTree := &Node{
			Value: 1,
			Left: &Node{
				Value: 2,
				Right: &Node{
					Value: 3,
				},
				Left: &Node{
					Value: 1,
					Left: &Node{
						Value: 0,
					},
				},
			},
		}



	})
}
