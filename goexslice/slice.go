package goexslice

import (
	"github.com/birdmichael/GoEx/tupleext"
	"math/rand"
	"time"
)

type Predicate[E any] func(value E) bool

type Comparator[E any] func(value1 E, value2 E) bool

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
func Contain[S ~[]E, E comparable](slice S, target E) bool {
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
func ContainBy[S ~[]E, E any](slice S, predicate func(item E) bool) bool {
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
func ContainSubSlice[S ~[]E, E comparable](slice, subSlice S) bool {

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
func ContainsAll[S ~[]E, E comparable](slice S, elements ...E) bool {

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
func ContainsAny[S ~[]E, E comparable](slice S, elements ...E) bool {
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
func Difference[S ~[]E, E comparable](slice, comparedSlice S) S {
	var diff []E

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
func DifferenceBy[S ~[]E, E any](slice, comparedSlice S, predicate func(s1 E, s2 E) bool) S {
	var diff []E

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

// MARK: - Group

// Group 将切片按指定的大小分割成多个子切片。
//
// 参数：
//   - slice: 要分割的切片。
//   - size: 每个子切片的大小。
//
// 返回值：
//   - 一个包含多个子切片的切片，每个子切片包含原切片中的连续元素。
//
// 示例：
//   - Group([]int{1, 2, 3, 4, 5, 6, 7}, 3) 返回 [][]int{{1, 2, 3}, {4, 5, 6}, {7}}，将切片分为大小为 3 的子切片。
//   - Group([]string{"a", "b", "c", "d"}, 2) 返回 [][]string{{"a", "b"}, {"c", "d"}}，将切片分为大小为 2 的子切片。
//   - Group([]int{1, 2, 3, 4, 5}, 0) 返回 nil，因为指定的大小为 0。
//   - Group([]int{}, 3) 返回 nil，因为切片为空。
func Group[S ~[]E, E any](slice S, size int) []S {
	var result []S

	if len(slice) == 0 || size <= 0 {
		return result
	}

	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []E{})
			l++
		}

		result[l-1] = append(result[l-1], item)
	}

	return result
}

// MARK: - Pairs

// Filter 根据条件函数过滤切片中的元素。
//
// 参数：
//   - slice: 要过滤的切片。
//   - predicate: 用于判断是否保留元素的条件函数。
//
// 返回值：
//   - 经过过滤的新切片，其中仅包含满足条件的元素。
func Filter[S ~[]E, E any](slice S, predicate Predicate[E]) []E {
	if predicate == nil {
		return slice
	}
	result := make(S, 0)
	for _, value := range slice {
		b := predicate(value)
		if b {
			result = append(result, value)
		}
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
func Prepend[S ~[]E, E any](slice S, e E) S {
	return append([]E{e}, slice...)
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

// MARK: - Find

// FindFirstBy 根据指定的条件查找切片中第一个满足条件的元素。
//
// 参数：
//   - slice: 要查找元素的切片。
//   - predicate: 用于判断元素是否满足条件的函数，接受元素索引和元素值作为参数，返回布尔值。
//
// 返回值：
//   - v: 第一个满足条件的元素。
//   - ok: 表示是否找到满足条件的元素。
//
// 示例：
//   - FindFirstBy([]int{1, 2, 3, 4}, func(i int, v int) bool { return v > 2 }) 返回 3 和 true，找到第一个大于 2 的元素。
//   - FindFirstBy([]string{"apple", "banana", "cherry"}, func(i int, v string) bool { return len(v) > 5 }) 返回 "banana" 和 true，找到第一个长度大于 5 的元素。
//   - FindFirstBy([]int{}, func(i int, v int) bool { return v > 2 }) 返回 0 和 false，切片为空，没有找到满足条件的元素。
func FindFirstBy[S ~[]E, E any](slice S, predicate func(index int, item E) bool) (v E, ok bool) {
	index := -1

	for i, item := range slice {
		if predicate(i, item) {
			index = i
			break
		}
	}

	if index == -1 {
		return v, false
	}

	return slice[index], true
}

// FindLastBy 根据指定的条件查找切片中最后一个满足条件的元素。
//
// 参数：
//   - slice: 要查找元素的切片。
//   - predicate: 用于判断元素是否满足条件的函数，接受元素索引和元素值作为参数，返回布尔值。
//
// 返回值：
//   - v: 最后一个满足条件的元素。
//   - ok: 表示是否找到满足条件的元素。
//
// 示例：
//   - FindLastBy([]int{1, 2, 3, 4}, func(i int, v int) bool { return v > 2 }) 返回 4 和 true，找到最后一个大于 2 的元素。
//   - FindLastBy([]string{"apple", "banana", "cherry"}, func(i int, v string) bool { return len(v) > 5 }) 返回 "cherry" 和 true，找到最后一个长度大于 5 的元素。
//   - FindLastBy([]int{}, func(i int, v int) bool { return v > 2 }) 返回 0 和 false，切片为空，没有找到满足条件的元素。
func FindLastBy[S ~[]E, E any](slice S, predicate func(index int, item E) bool) (v E, ok bool) {
	index := -1

	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(i, slice[i]) {
			index = i
			break
		}
	}

	if index == -1 {
		return v, false
	}

	return slice[index], true
}

// MARK: - Safe

// SafeSwap 安全地交换切片中两个索引位置的元素。
//
// 参数：
//   - slice: 要进行元素交换的切片。
//   - from: 要交换的第一个元素的索引。
//   - to: 要交换的第二个元素的索引。
//
// 返回值：
//   - 如果交换成功，返回 true；否则返回 false。
func SafeSwap[S ~[]E, E any](slice S, from, to int) (ok bool) {
	if from == to {
		return false
	}
	if from < 0 || from >= len(slice) {
		return false
	}
	if to < 0 || to >= len(slice) {
		return false
	}

	slice[from], slice[to] = slice[to], slice[from]
	return true
}

// SafeIndex 安全地获取切片中指定索引位置的元素。
//
// 参数：
//   - slice: 要获取元素的切片。
//   - index: 要获取元素的索引。
//
// 返回值：
//   - 如果索引有效，返回该索引位置的元素和 true；
//     如果索引无效，返回零值和 false。
func SafeIndex[S ~[]E, E any](slice S, index int) (v E, ok bool) {
	if index < 0 || index >= len(slice) {
		return v, false
	}

	return slice[index], true
}

// MARK: - Property

// Reverse 将切片中的元素顺序颠倒。
//
// 参数：
//   - slice: 要颠倒元素顺序的切片。
//
// 示例：
//   - Reverse([]int{1, 2, 3, 4, 5}) 将切片中的元素顺序颠倒。
//   - Reverse([]string{"a", "b", "c", "d"}) 将切片中的元素顺序颠倒。
func Reverse[S ~[]E, E any](slice S) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// First 返回切片中的第一个元素。
//
// 参数：
//   - slice: 要获取第一个元素的切片。
//
// 返回值：
//   - v: 切片中的第一个元素。
//   - ok: 表示是否找到第一个元素。
//
// 示例：
//   - First([]int{1, 2, 3}) 返回 1 和 true，获取切片中的第一个元素。
//   - First([]string{"apple", "banana", "cherry"}) 返回 "apple" 和 true，获取切片中的第一个元素。
//   - First([]int{}) 返回 0 和 false，切片为空，没有找到元素。
//   - First(nil) 返回 0 和 false，切片为 nil，没有找到元素。
func First[S ~[]E, E any](slice S) (v E, ok bool) {
	if len(slice) == 0 {
		return v, false
	}

	return slice[0], true
}

// Last 返回切片中的最后一个元素。
//
// 参数：
//   - slice: 要获取最后一个元素的切片。
//
// 返回值：
//   - v: 切片中的最后一个元素。
//   - ok: 表示是否找到最后一个元素。
//
// 示例：
//   - Last([]int{1, 2, 3}) 返回 3 和 true，获取切片中的最后一个元素。
//   - Last([]string{"apple", "banana", "cherry"}) 返回 "cherry" 和 true，获取切片中的最后一个元素。
//   - Last([]int{}) 返回 0 和 false，切片为空，没有找到元素。
func Last[S ~[]E, E any](slice S) (v E, ok bool) {
	if len(slice) == 0 {
		return v, false
	}

	return slice[len(slice)-1], true
}

// RandomIn 随机打乱切片中的元素顺序，直接修改原始切片的值。
//
// 参数：
//   - slice: 要打乱顺序的切片。
//
// 示例：
//   - slice := []int{1, 2, 3, 4, 5}
//   - RandomIn(slice) 会修改原始切片 slice 的元素顺序。
func RandomIn[S ~[]E, E any](slice S) {
	rand.NewSource(time.Now().UnixNano())

	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// RandomCopy 返回一个随机打乱切片中的元素顺序的新切片。
//
// 参数：
//   - slice: 要打乱顺序的切片。
//
// 返回值：
//   - 一个包含随机打乱顺序后的新切片。
//
// 示例：
//   - originalSlice := []int{1, 2, 3, 4, 5}
//   - resultSlice := RandomCopy(originalSlice) 返回一个随机打乱顺序后的新切片。
func RandomCopy[S ~[]E, E any](slice S) S {
	newSlice := make(S, len(slice))
	copy(newSlice, slice)

	RandomIn(newSlice)

	return newSlice
}

// Enumerated 返回一个包含切片中每个元素及其对应索引的元组切片。
//
// 参数：
//   - slice: 要获取索引的切片。
//
// 返回值：
//   - 一个包含元组的切片，每个元组包含切片中的元素和对应的索引。
//
// 示例：
//   - slice := []int{10, 20, 30, 40}
//   - Enumerated(slice) 返回一个包含元组的切片，例如 []Tuple[int, int]{{S1: 10, S2: 0}, {S1: 20, S2: 1}, {S1: 30, S2: 2}, {S1: 40, S2: 3}}
func Enumerated[S ~[]E, E any](slice S) []tupleext.Tuple[E, int] {
	var result []tupleext.Tuple[E, int]

	for i, item := range slice {
		result = append(result, tupleext.Tuple[E, int]{S1: item, S2: i})
	}

	return result
}

// IsEmpty 判断切片是否为空。
//
// 参数：
//   - slice: 要判断是否为空的切片。
//
// 返回值：
//   - 如果切片为空，返回 true；否则返回 false。
func IsEmpty[S ~[]E, E any](slice S) bool {
	return len(slice) == 0
}
