package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}


// 结点的两个基础方法
func (n *node) matchKid(part string) *node { // 入参是url的一段
	for _, kid := range n.children {
		if kid.part == part || kid.isWild {
			return kid
		}
	}
	return nil // 没找到
}

func (n *node) matchKids(part string) []*node {
	goodnodes := []*node{}
	for _, kid := range n.children {
		if kid.part == part || kid.isWild {
			goodnodes = append(goodnodes, kid)
		}
	}
	return goodnodes
}





// 用于注册路由
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return 
	}

	part := parts[height]
	kid := n.matchKid(part)

	// 没有符合的 kid，新建结点
	if kid == nil {
		kid = &node {
			part: part,
			isWild : part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, kid)
	}
	kid.insert(pattern, parts, height+1)
}


func (n *node) search(parts []string, height int) *node {
	// 找到底了 or *filepath 通配
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { // 无对应 pattern
			return nil
		}
		return n
	}

	part := parts[height]
	kids := n.matchKids(part)

	for _, kid := range kids {
		cur := kid.search(parts, height+1)
		if cur != nil {
			return cur
		}
	}

	return nil
}