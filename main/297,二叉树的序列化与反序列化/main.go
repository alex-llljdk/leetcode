package main

import (
	"strconv"
	"strings"
)

// dfs 写法
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

const (
	NULL = "null"
)

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	strings.Trim()
	return build()
}

// 层序遍历写法
func (this *Codec) serialize(root *TreeNode) string {
	var result []string
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	hasNext := true
	for hasNext && len(queue) > 0 {
		size := len(queue)
		hasNext = false
		for i := 0; i < size; i++ {
			item := queue[i]
			if item == nil {
				result = append(result, NULL)
				continue
			}
			result = append(result, strconv.Itoa(item.Val))
			//  item != nil
			queue = append(queue, item.Left)
			queue = append(queue, item.Right)
			if item.Left != nil || item.Right != nil {
				hasNext = true
			}
		}
		queue = queue[size:]
	}
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = strings.Trim(strings.TrimSpace(data), ",")
	items := strings.Split(data, ",")
	if len(items) == 0 {
		return nil
	}
	if items[0] == NULL {
		return nil
	}
	val, _ := strconv.Atoi(items[0])
	root := &TreeNode{Val: val}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	itemIdx := 1
	for len(queue) > 0 && itemIdx < len(items) {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			node.Left = buildNode(items[itemIdx])
			itemIdx++
			node.Right = buildNode(items[itemIdx])
			itemIdx++
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return root
}

func buildNode(s string) *TreeNode {
	if s == NULL {
		return nil
	}
	val, _ := strconv.Atoi(s)
	return &TreeNode{Val: val}
}

func main() {

}
