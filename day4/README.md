# Day 4 分组控制 Route Group Control

如果没有分组控制，需要为每个路由单独进行控制，但实际场景中，经常有一组路由有相似or相同的处理。

一般的路由分组是以相同的前缀来区分。

本次实现：
* /post 开头的可匿名访问
* /admin 开头的需要鉴权
* /api 开头的是 RESTful 接口，可以对接第三方，需要第三方平台鉴权

接下来是Group对象的设计：
* 需要前缀
* 支持分组嵌套，所以有 parent 结点
* 还有支持的中间件
* 资源由 engine（服务器实例）统一协调
    * 所以属性有它的 *Engine，且 Enigne 实例中有绑在它身上的所有 Group 的列表

```go
// 分组的结点设计
RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}
```

```go
// Engine 的设计变化
Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}
```

这样之后，与路由的交互自然而然地从 Engine 上转到 RouterGroup上，功能进一步抽象与封装。


# 完事感想
今天实现了「分组控制」，但其实着重在分组的实现上，出了把之前的简单访问功能分了组，并没进行特别的控制。

但也实现了通过 /v1 与 /v2前缀来导向不同的响应。

明天是中间件，所以我猜明天应该会把中间件的功能填到RouterGroup结点预留的属性上。