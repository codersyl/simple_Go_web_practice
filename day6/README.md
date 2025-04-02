# day 6

今天添加的功能是模版与渲染。

采用服务端渲染

RouterGroup 增加两方法
* createStaticHandler 
* Static 把 网页的path 跟 服务器的path联系起来

html/template 库 为 HTML 提供了比较完整的支持

解析模版在实际开发中约等于读取模版文件，然后分析结构，提取变量、表达式等需要动态生成的部分，替换动态内容，生成一个完整文本（通常是HTML文件）。



## Glob 模式
glob 模式是一种用于匹配文件路径的通配符模式。它允许你使用特定的字符来指定一组文件或目录，而无需列出每个单独的文件或目录名称。

以下是一些常见的 glob 模式通配符及其含义：

* *（星号）：匹配零个或多个任意字符。
* ?（问号）：匹配单个任意字符。
* []（方括号）：匹配方括号内指定的字符集中的任何单个字符。例如：
    * file[0-9].txt 匹配 file0.txt、file1.txt 等。
    * file[abc].txt 匹配 filea.txt、fileb.txt、filec.txt。
* **（双星号）：匹配零个或多个目录。这个通配符在递归搜索子目录时非常有用。例如：`docs/**/*.txt` 匹配 docs 目录及其所有子目录中的所有 .txt 文件。


templates 文件夹下的三个模版文件不了解，用于充当模版

# 运行后可用的测试

`curl http://localhost:9999/date` 返回当前日期

`curl http://localhost:9999/time` 返回当前年月日时分秒

`curl http://localhost:9999/students` 返回写在main函数里预制的两个学生的信息

`curl http://localhost:9999/assets/css/jolyne.css` 返回静态文件jolyne.css的内容