# Go Modules

* go.mod 是 Go Modules 的核心
* Go Modules 用于之前 Go 语言的依赖管理混乱的问题

# go.mod
* go.mod 就像一个清单，列出了项目所使用的所有外部库及其版本

# 如何编写 go.mod

* 项目跟目录下运行 `go mod init <path>`
    * path 填写代码托管平台（如github）上该项目的地址
* 当导入一个新的外部库时，Go会自动更新 go.mod
* 也可以使用 go get <库路径> 添加


* `go get -u <库路径>` 可以把库更新到最新版本
* `go mod tidy` 可以清理go.mod中不再需要的依赖

## example
```go
module example // 模块名

go 1.23.0 // 项目所需最低版本 

require gee v0.0.1
// 列出项目直接依赖的模块及其版本
// 此处是本地的gee

replace gee => ./gee
```