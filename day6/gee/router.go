package gee

import (
	"strings"
	"net/http"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']


func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}


// 解析路由模式
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' { // 最多只能一个 *
				break
			}
		}
	}
	return parts
}

// 注册路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok { // 之前没有该请求方法的模式
		r.roots[method] = &node{}
	}
	
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}


// 查找路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	
	root, ok := r.roots[method]

	//不支持该种请求方法
	if !ok { 
		return nil, nil
	}

	n := root.search(searchParts, 0)

	// 没找到对应的 pattern结点
	if n == nil {
		return nil, nil
	}

	params := make(map[string]string) // 一些参数录入
	parts := parsePattern(n.pattern)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}
	return n, params
}


func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern // 此处 pattern 为节点中存储的

		// 把路由中注册的处理函数加到 context 实例中的执行列表里
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		// 找不到匹配的handler，拒绝服务
		// 保持一致性，所以把这个包成一个handler函数加到 handlers中去
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next() // 按序执行handlers
}