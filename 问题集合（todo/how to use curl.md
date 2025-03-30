curl 的基础用法

`curl https://example.com` 是发送HTTP默认的GET请求

`curl -o output.html https://example.com` 把GET得到的结果保存到 output.html 中

`curl -O https://example.com/image.jpg` 使用大O自动存为远程文件的名字，此处存为image.jpg


如果想查看 curl 发送的请求头和服务器返回的响应头（而不仅仅是响应体），需要什么参数？ 
使用 -v 或 --verbose 参数：
curl -v https://example.com

问题5

如何用 curl 发送一个 POST 请求，并附带表单数据（例如 username=alice&password=123）？

答案
使用 -X POST 指定方法，并用 -d 传递数据：
curl -X POST -d "username=alice&password=123" https://example.com/login


问题6

如果服务器要求请求头中包含 Content-Type: application/json，如何用 curl 设置这个头信息？

答案
使用 -H 参数指定自定义头：
curl -H "Content-Type: application/json" https://example.com/api

问题7

当服务器返回重定向（如 302 状态码）时，如何让 curl 自动跟随重定向获取最终内容？

答案
使用 -L 参数：
curl -L https://example.com/redirect-page

问题8

如何用 curl 上传本地文件（如 data.txt）到服务器？假设服务器接口是 https://example.com/upload，且使用 POST 方法。

答案
使用 -F 或 --form 参数上传文件：
curl -F "file=@data.txt" https://example.com/upload

问题9

如果想调试 curl 请求的实际耗时（如 DNS 解析、连接建立等阶段的时间），可以添加什么参数？

答案
使用 --trace-time 或结合 -w 自定义输出格式：
curl --trace-time -w "DNS解析耗时: %{time_namelookup}\n总耗时: %{time_total}\n" https://example.com

问题10

如果请求需要携带 Cookie（例如 sessionid=abc123），如何用 curl 实现？

答案
使用 -b 参数指定 Cookie：
curl -b "sessionid=abc123" https://example.com/dashboard

总结思考

这些问题覆盖了 curl 的以下核心功能：

基本 GET/POST 请求
文件下载与上传
头信息操作
调试与性能分析
重定向与 Cookie 处理