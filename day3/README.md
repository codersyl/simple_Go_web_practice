# day 3
今天需要实现动态路由

# 两种模式

## :name
* 只能捕获一个值
* 参数匹配:。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
* 

## *filepath
* 能捕获 0 或多个值
* 通配*。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js
* 常用于静态服务器，能够递归地匹配子路径

# 使用Trie数据接口

## 结点设计分析
```go
type node struct {
	pattern  string // 待匹配路由，例如 /p/:lang
	part     string // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool // 是否精确匹配，part 含有 : 或 * 时为true
}
```

* 动态路由需要实现两个功能
    * 注册路由
        * 对应 Trie 的功能是插入节点
    * 匹配路由
        * 对应 Trie 的功能是查找


则结点的结构体需要两个函数（方法）：
1. 找到匹配成功的第一个结点，用于插入
2. 找到所有匹配成功的结点，用于查找

# day3实践感想

解耦（gee包分成几个文件）的好处初步显现，day3是增加动态路由功能，所以对context.go的部分只动了一点，并且gee.go完全不用动，因为调用的router的方法的参数个数和类型完全没变。