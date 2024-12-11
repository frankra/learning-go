package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Node struct {
	Depth    int
	IsDir    bool
	Name     string
	Children []*Node
}

func (n *Node) Print() {
	prefix := ""
	for i := 0; i < n.Depth; i++ {
		prefix += "_"

	}
	fmt.Printf("%s%s\n", prefix, n.Name)

	for _, child := range n.Children {
		child.Print()
	}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func main() {
	// var ref *Node

	root := &Node{
		Depth:    0,
		IsDir:    true,
		Name:     "/",
		Children: []*Node{},
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		parts := strings.Split(path, "/")

		addToTree(root, parts, 1, info)

		return nil
	})

	if err != nil {
		fmt.Print(err)
	}

	root.Print()
}

func addToTree(node *Node, parts []string, depth int, info os.FileInfo) {
	if len(parts) > 0 {
		part := parts[0]
		for _, child := range node.Children {
			if child.Name == part {
				addToTree(child, parts[1:], depth+1, info)
				return
			}
		}
		newChildren := &Node{
			Depth:    depth,
			IsDir:    info.IsDir(),
			Name:     part,
			Children: []*Node{},
		}
		node.AddChild(newChildren)
	}
}
