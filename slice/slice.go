package slice

// MARK: - Contain

// Contain 检查slice中是否包含元素target
//
// 这个方法使用泛型来支持任意可比较类型
//
// 例子:
//
//	Contain([]int{1,2,3}, 2) -> true
//
// 参数:
// - slice: 要检查的切片
// - target: 要查找的目标元素
//
// 返回值:
// - bool: 如果slice包含target,返回true;否则返回false
func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}

	return false
}

// ContainBy 检查 slice 中是否存在满足 predicate 的元素
//
// 这个方法支持任意类型的 slice
//
// 例子:
//
//	ContainBy([]int{1, 2, 3}, func(item int) bool { return item > 1 }) -> true
//
// 参数:
// - slice:要检查的切片
// - predicate:判断元素是否满足条件的函数
//
// 返回值:
// - bool:如果存在满足条件的元素,返回 true;否则返回 false
func ContainBy[T any](slice []T, predicate func(item T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}

	return false
}

// ContainSubSlice 检查切片 slice 是否包含子切片 subSlice
//
// 该方法要求切片元素类型是可比较的
//
// 例如:
//
//	ContainSubSlice([]int{1,2,3,4}, []int{2,3}) -> true
//
// 参数:
// - slice: 要检查的切片
// - subSlice: 子切片
//
// 返回值:
// - bool: 如果 slice 包含 subSlice,返回 true;否则返回 false
func ContainSubSlice[T comparable](slice, subSlice []T) bool {

	// 如果子切片长度大于父切片,直接返回 false
	if len(subSlice) > len(slice) {
		return false
	}

	for _, v := range subSlice {
		if !Contain(slice, v) {
			return false
		}
	}

	return true
}

// ContainsAll 检查切片是否包含所有指定的元素。
//
// 参数：
//   - slice: 要检查的切片。
//   - elements: 一个可变参数列表，包含要检查是否存在于切片中的元素。
//
// 返回值：
//   - 如果切片包含所有指定的元素，则返回 true；否则返回 false。
//
// 示例：
//   - ContainsAll([]int{1, 2, 3, 4}, 2, 4) 返回 true，因为切片包含所有指定的元素。
//   - ContainsAll([]string{"apple", "banana", "cherry"}, "banana", "grape") 返回 false，因为切片不包含所有指定的元素。
func ContainsAll[T comparable](slice []T, elements ...T) bool {

	// 如果子切片长度大于父切片,直接返回 false
	if len(elements) > len(slice) {
		return false
	}

	for _, v := range elements {
		if !Contain(slice, v) {
			return false
		}
	}

	return true
}

// ContainsAny 检查切片是否包含指定的任何一个元素。
//
// 参数：
//   - slice: 要检查的切片。
//   - elements: 一个可变参数列表，包含要检查是否存在于切片中的元素。
//
// 返回值：
//   - 如果切片包含指定的任何一个元素，则返回 true；否则返回 false。
//
// 示例：
//   - ContainsAny([]int{1, 2, 3, 4}, 2, 5) 返回 true，因为切片包含其中一个指定的元素（2）。
//   - ContainsAny([]string{"apple", "banana", "cherry"}, "grape", "orange") 返回 false，因为切片不包含任何一个指定的元素。
func ContainsAny[T comparable](slice []T, elements ...T) bool {
	for _, v := range elements {
		if Contain(slice, v) {
			return true
		}
	}
	return false
}

// MARK: - Difference

// Difference 返回两个切片的差异元素，即存在于切片 a 但不存在于切片 b，以及存在于切片 b 但不存在于切片 a 的元素。
//
// 参数：
//   - slice1: 第一个切片。
//   - slice2: 第二个切片。
//
// 返回值：
//   - 一个包含差异元素的切片，包括存在于切片 a 但不存在于切片 b 的元素，以及存在于切片 b 但不存在于切片 a 的元素。
//
// 示例：
//   - Difference([]int{1, 2, 3, 4, 5}, []int{3, 4, 6}) 返回 []int{1, 2, 5, 6}
func Difference[T comparable](slice, comparedSlice []T) []T {
	var diff []T

	for i := 0; i < 2; i++ {
		for _, s1 := range slice {
			found := false
			for _, s2 := range comparedSlice {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice, comparedSlice = comparedSlice, slice
		}
	}

	return diff
}

// DifferenceBy 返回一个新切片，其中包含在 slice 中但不在 comparedSlice 中的元素，使用 predicate 函数进行比较。
//
// 参数：
//   - slice: 要比较的切片。
//   - comparedSlice: 用于比较的切片。
//   - predicate: 用于比较切片元素的函数，返回值为 true 表示元素匹配。
//
// 返回值：
//   - 包含在 slice 中但不在 comparedSlice 中的元素的新切片。
//
// 示例：
//   - DifferenceBy([]int{1, 2, 3, 4}, []int{3, 4, 5}, func(s1, s2 int) bool { return s1 == s2 })
//     返回 []int{1, 2, 5}，使用 predicate 进行比较。
func DifferenceBy[T comparable](slice []T, comparedSlice []T, predicate func(s1 T, s2 T) bool) []T {
	var diff []T

	for i := 0; i < 2; i++ {
		for _, s1 := range slice {
			found := false
			for _, s2 := range comparedSlice {
				if predicate(s1, s2) {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice, comparedSlice = comparedSlice, slice
		}
	}

	return diff
}

// MARK: - Chunk

// Chunk 将切片按指定的大小分割成多个子切片。
//
// 参数：
//   - slice: 要分割的切片。
//   - size: 每个子切片的大小。
//
// 返回值：
//   - 一个包含多个子切片的切片，每个子切片包含原切片中的连续元素。
//
// 示例：
//   - Chunk([]int{1, 2, 3, 4, 5, 6, 7}, 3) 返回 [][]int{{1, 2, 3}, {4, 5, 6}, {7}}，将切片分为大小为 3 的子切片。
//   - Chunk([]string{"a", "b", "c", "d"}, 2) 返回 [][]string{{"a", "b"}, {"c", "d"}}，将切片分为大小为 2 的子切片。
//   - Chunk([]int{1, 2, 3, 4, 5}, 0) 返回 nil，因为指定的大小为 0。
//   - Chunk([]int{}, 3) 返回 nil，因为切片为空。
func Chunk[T any](slice []T, size int) [][]T {
	var result [][]T

	if len(slice) == 0 || size <= 0 {
		return result
	}

	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []T{})
			l++
		}

		result[l-1] = append(result[l-1], item)
	}

	return result
}

// MARK: - Insert

// Prepend 在切片的开头添加一个元素并返回新的切片。
//
// 参数：
//   - slice: 原始切片。
//   - e: 要添加到切片开头的元素。
//
// 返回值：
//   - 一个包含添加元素后的新切片。
//
// 示例：
//   - Prepend([]int{2, 3, 4}, 1) 返回 []int{1, 2, 3, 4}，在切片开头添加了元素 1。
//   - Prepend([]string{"b", "c"}, "a") 返回 []string{"a", "b", "c"}，在切片开头添加了元素 "a"。
func Prepend[T any](slice []T, e T) []T {
	return append([]T{e}, slice...)
}

// InsertAt 在指定索引处插入元素到切片中，并返回新的切片。
//
// 参数：
//   - slice: 原始切片。
//   - index: 要插入元素的索引位置。
//   - value: 要插入的元素值。可以是单个元素或切片。
//
// 返回值：
//   - 一个包含插入元素后的新切片。如果索引超出范围，将返回原始切片。
//
// 示例：
//   - InsertAt([]int{1, 2, 3}, 1, 4) 返回 []int{1, 4, 2, 3}，在索引 1 处插入元素 4。
//   - InsertAt([]string{"a", "b"}, 0, []string{"x", "y"}) 返回 []string{"x", "y", "a", "b"}，在索引 0 处插入切片 {"x", "y"}。
func InsertAt[T any](slice []T, index int, value any) []T {
	size := len(slice)

	if index < 0 || index > size {
		return slice
	}

	if v, ok := value.(T); ok {
		slice = append(slice[:index], append([]T{v}, slice[index:]...)...)
		return slice
	}

	if v, ok := value.([]T); ok {
		slice = append(slice[:index], append(v, slice[index:]...)...)
		return slice
	}

	return slice
}

// MARK: - Reverse

// Reverse 将切片中的元素顺序颠倒。
//
// 参数：
//   - slice: 要颠倒元素顺序的切片。
//
// 示例：
//   - Reverse([]int{1, 2, 3, 4, 5}) 将切片中的元素顺序颠倒。
//   - Reverse([]string{"a", "b", "c", "d"}) 将切片中的元素顺序颠倒。
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
