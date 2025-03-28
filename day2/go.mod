module example

go 1.23.0

require gee v0.0.1

// 列出项目直接依赖的模块及其版本
// 此处是本地的gee

replace gee => ./gee
