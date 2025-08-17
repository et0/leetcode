package lh0297 // https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	var data strings.Builder
	null := 0

	nodes := make([]*TreeNode, 0, 10000)
	nodes = append(nodes, root)

	for size := len(nodes); size > 0; size = len(nodes) {
		for _, n := range nodes {
			if n == nil {
				null++
				continue
			}

			if null > 0 {
				data.WriteString(strings.Repeat(":n", null))
				null = 0
			}
			if data.Len() > 0 {
				data.WriteString(":")
			}
			data.WriteString(strconv.Itoa(n.Val))

			if n.Left != nil {
				nodes = append(nodes, n.Left)
			} else {
				nodes = append(nodes, nil)
			}

			if n.Right != nil {
				nodes = append(nodes, n.Right)
			} else {
				nodes = append(nodes, nil)
			}
		}
		nodes = nodes[size:]
	}

	return data.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	vals := make([]*TreeNode, 0, 10000)
	nodes := make([]*TreeNode, 0, 10000)

	split := strings.Split(data, ":")

	val, _ := strconv.Atoi(split[0])
	head := &TreeNode{Val: val}
	nodes = append(nodes, head)
	sizeNodes := len(nodes)
	sizeSplit := len(split)
	for i := 1; i < sizeSplit; i++ {
		if split[i] == "n" {
			vals = append(vals, nil)
		} else {
			val, _ := strconv.Atoi(split[i])
			vals = append(vals, &TreeNode{Val: val})
		}

		sizeVals := len(vals)
		if sizeVals < sizeNodes*2 && i+1 < sizeSplit {
			continue
		}

		for k, n := range nodes {
			if k*2 >= sizeVals {
				break
			}
			if vals[k*2] != nil {
				n.Left = vals[k*2]
				nodes = append(nodes, n.Left)
			}

			if k*2+1 >= sizeVals {
				break
			}
			if vals[k*2+1] != nil {
				n.Right = vals[k*2+1]
				nodes = append(nodes, n.Right)
			}
		}

		vals = vals[:0]
		nodes = nodes[sizeNodes:]
		sizeNodes = len(nodes)

	}

	return head
}

func Wrapper(root *TreeNode) {
	c := Constructor()
	// data := c.serialize(root)
	c.deserialize("1:2:3:n:n:4:17:5:11:18:n:6:n:12:14:n:n:7:9:13:n:15:n:n:8:10:n:n:n:16")
}
