# day 5


中间件(middlewares)，非业务的技术类组件。Web 框架本身不可能去理解所有的业务，不可能实现所有的功能。因此，框架需要有一个插口，允许用户自己定义功能，嵌入到框架中。


需要考虑两个点：
* 插入点在哪？
    * 使用框架的人并不关心底层逻辑的具体实现，如果插入点太底层，中间件逻辑就会非常复杂
    * 如果插入点离用户太近，那和用户直接定义一组函数，每次在 Handler 中手工调用没有多大的优势了。
* 中间件的输入是什么？
    * 需要有足够的输入

设计参考了 Gin

* 定义与handler()一致，处理的对象是 Context
* 可在handle前与handle后做一些操作
* 支持多个中间件，依次调用
* 可使用 c.Next()进行等待，等待其他操作结束后再回来
    * 比如说想统计处理本次请求所花费的总用时





```go
// 需要分两截执行的中间件 belike：
func A(c *Context) {
    part1
    c.Next()
    part2
}
// 由于长成这样了，所以多个「两截中间件」的后半部分执行顺序是压栈进行，后进先出
```

可以把路由映射的 handler 放到 Context 实例的handlers队列中的最后一个，这样更统一

## 需要修改的地方

Context 结构体需要多两个变量，一个方法
* handlers，登记所有的中间件
* index，context内部的中间件执行到了什么程度
* .Next()，用于先执行其他中间件

gee.go 中  
* 多一个 RouterGroup 的方法Use，功能是把中间件（列表）加到RouterGroup结点中去
* ServeHTTP也变了
    1. 遍历服务器中所有的RouterGroup结点
    2. 把符合前缀的中间件全都加入到该次服务的 Context 实例c中
    3. 最后handle c

router.go 中  
* handle的逻辑变成：
    * 动态路由匹配成功则干嘛则把handler加到handlers结尾
        * 匹配不成功则把「拒绝服务」写成handler加到handlers结尾
    * c.Next()依次执行所有handler