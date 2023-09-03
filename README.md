# GoEx

GoEx 是一个提供 Go 的实用工具库。

希望Go能像Swift一样，有非常多快捷的工具函数、方法，GoEx就是这样一个工具库。

## 安装

要在您的 Go 项目中使用 GoEx，首先确保您已经安装了 Go。然后，您可以使用以下命令将其添加为依赖：

```
go get github.com/birdmichael/GoEx
```

## 使用
<details>
<summary>切片包含操作</summary>

```go
import "github.com/birdmichael/GoEx/slice"

// 检查切片中是否包含特定元素
result := slice.Contain([]int{1, 2, 3}, 2) // 返回 true

// 检查切片中是否存在满足特定条件的元素
result := slice.ContainBy([]int{1, 2, 3}, func(item int) bool { return item > 1 }) // 返回 true

// 检查一个切片是否包含另一个子切片
result := slice.ContainSubSlice([]int{1, 2, 3, 4}, []int{2, 3}) // 返回 true
```

</details>

<details>
<summary>切片差异操作</summary>

```go
// 获取两个切片的差异元素
diff := slice.Difference([]int{1, 2, 3, 4, 5}, []int{3, 4, 6}) // 返回 []int{1, 2, 5, 6}
```

</details>

<details>
<summary>切片分块操作</summary>

```go
// 将切片按指定的大小分割成多个子切片
chunks := slice.Chunk([]int{1, 2, 3, 4, 5, 6, 7}, 3) // 返回 [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
```

</details>

<details>
<summary>切片插入操作</summary>

```go
// 在切片的开头添加一个元素
newSlice := slice.Prepend([]int{2, 3, 4}, 1) // 返回 []int{1, 2, 3, 4}

// 在指定索引处插入元素
newSlice := slice.InsertAt([]int{1, 2, 3}, 1, 4) // 返回 []int{1, 4, 2, 3}
```

</details>

<details>
<summary>切片反转操作</summary>

```go
// 将切片中的元素顺序颠倒
slice := []int{1, 2, 3, 4, 5}
slice.Reverse(slice) // slice 现在是 []int{5, 4, 3, 2, 1}
```

</details>





