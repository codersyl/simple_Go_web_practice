# 如何在 Go 中做单元测试

* Go 语言推荐测试文件和源代码文件放在一块，测试文件以 _test.go 结尾。
    * 比如，当前 package 有 calc.go 一个文件，我们想测试 calc.go 中的 Add 和 Mul 函数，那么应该新建 calc_test.go 作为测试文件。
* 测试用例名称一般命名为 Test 加上待测试的方法名。
* 测试用的参数有且只有一个，在这里是 t *testing.T。
* 基准测试(benchmark)的参数是 *testing.B，TestMain 的参数是 *testing.M 类型。


```go
// calc.go
package main

func Add(a int, b int) int {
    return a + b
}

func Mul(a int, b int) int {
    return a * b
}
```

```go
// calc_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}
```

