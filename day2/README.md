# Context

* 生命周期：随着请求的出现而产生，随着请求的结束而销毁
* 与请求强相关的信息都由Context承载
* 路由的处理函数、将要实现的中间件、参数都统一使用 Context 实例

工作流程：
1. 请求出现，把请求中的信息解析出来，装入 context 实例 c 中
2. 调用 router，对实例 c 进行处理
    * 根据 c 中的信息，查找路由映射表中是否含有对应处理方法
        * 有就调用
        * 没有就输出404
3. 过程 2 中 c 的剩下一半信息（响应部分）也写好了

# Router

* 把与路由相关的内容封装到 Router 中，比如路由映射表、路由注册等
* 组成成分：
    * 路由映射表
    * 路由注册函数
    * 对 context 实例进行处理的函数handle
