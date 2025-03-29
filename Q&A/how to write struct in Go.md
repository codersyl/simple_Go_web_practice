

## 定义
```go
type Person struct {
    Name    string
    Age     int
    Address string
}
```

## 实例创建
```go
person1 := Person{
    Name:    "Alice",
    Age:     30,
    Address: "123 Main St", // 最后一行也必须加逗号
}
```

## 方法编写
```go
// 方式1，使用值接收器，调用的时候创建一个副本，方法内修改内部字段的值，不会影响本体
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// 方式2，指针接收器，调用的时候修改字段内部的值会影响本体
// 使用场景，结构体很大，或者需要修改本体的值
func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}
```

* 使用 `Jack := &Person{}` 创建的人，仍然可以使用值接收器的方法，例如上述代码中的 `SayHello()` ，调用的话，Go会自动解引用指针