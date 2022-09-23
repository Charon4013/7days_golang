package gee

import (
	"fmt"
	"strings"
)

// Trie树
type node struct {
	pattern  string  // 等待匹配的路由 ex: /p/:lang
	part     string  // 路由中的一部分 ex: :lang
	children []*node // 子节点 ex: [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配, part有 : 或 * 时为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t", n.pattern, n.part, n.isWild)
}

// 第一个匹配成功的节点, 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点, 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// Trie Tree插入功能
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	// 递归查询
	child.insert(pattern, parts, height+1)
}

// Trie Tree查询功能
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
