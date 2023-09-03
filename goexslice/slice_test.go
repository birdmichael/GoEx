package goexslice

import (
	"reflect"
	"sort"
	"testing"
)

func TestContain(t *testing.T) {
	testCases := []struct {
		slice  []int
		target int
		want   bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 0, false},
		{[]int{}, 2, false},
		{nil, 2, false},
	}

	for _, tc := range testCases {
		got := Contain(tc.slice, tc.target)
		if got != tc.want {
			t.Errorf("Contain(%v, %d) = %t, want %t", tc.slice, tc.target, got, tc.want)
		}
	}
}

func TestContainBy(t *testing.T) {

	greaterThan1 := func(item int) bool { return item > 1 }

	testCases := []struct {
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{[]int{1, 2, 3}, greaterThan1, true},
		{[]int{1, 0, -1}, greaterThan1, false},
		{[]int{}, greaterThan1, false},
	}

	for _, tc := range testCases {
		got := ContainBy(tc.slice, tc.predicate)
		if got != tc.want {
			t.Errorf("ContainBy(%v, %T) = %t, want %t",
				tc.slice, tc.predicate, got, tc.want)
		}
	}

}

func TestContainSubSlice(t *testing.T) {

	testCases := []struct {
		slice    []int
		subSlice []int
		want     bool
	}{
		{[]int{1, 2, 3, 4}, []int{2, 3}, true},
		{[]int{1, 2, 3, 4}, []int{3, 2}, true},
		{[]int{1, 2, 3, 4}, []int{2, 5}, false},
	}

	for _, tc := range testCases {
		got := ContainSubSlice(tc.slice, tc.subSlice)
		if got != tc.want {
			t.Errorf("ContainSubSlice(%v, %v) = %t, want %t",
				tc.slice, tc.subSlice, got, tc.want)
		}
	}

}

func TestContainsAll(t *testing.T) {
	t.Run("TestContainsAll_Positive", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		elements := []int{2, 4}
		result := ContainsAll(slice, elements...)
		if !result {
			t.Errorf("Expected ContainsAll to return true, got false")
		}
	})

	t.Run("TestContainsAll_Negative", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		elements := []string{"banana", "grape"}
		result := ContainsAll(slice, elements...)
		if result {
			t.Errorf("Expected ContainsAll to return false, got true")
		}
	})
}

func TestContainsAny(t *testing.T) {
	t.Run("TestContainsAny_Positive", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		elements := []int{2, 5}
		result := ContainsAny(slice, elements...)
		if !result {
			t.Errorf("Expected ContainsAny to return true, got false")
		}
	})

	t.Run("TestContainsAny_Negative", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		elements := []string{"grape", "orange"}
		result := ContainsAny(slice, elements...)
		if result {
			t.Errorf("Expected ContainsAny to return false, got true")
		}
	})
}

func TestDifference(t *testing.T) {
	t.Run("TestDifference_IntSlice", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{3, 4, 6}
		expected := []int{1, 2, 5, 6}
		result := Difference(a, b)
		sort.Ints(result)
		sort.Ints(expected)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestDifference_StringSlice", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"cherry", "grape"}
		expected := []string{"apple", "banana", "grape"}
		result := Difference(a, b)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestDifferenceBy(t *testing.T) {
	t.Run("TestDifferenceBy_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		comparedSlice := []int{3, 4, 5}
		expected := []int{1, 2, 5}
		result := DifferenceBy(slice, comparedSlice, func(s1, s2 int) bool { return s1 == s2 })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestDifferenceBy_StringSlice", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		comparedSlice := []string{"c", "d"}
		expected := []string{"a", "b", "d"}
		result := DifferenceBy(slice, comparedSlice, func(s1, s2 string) bool { return s1 == s2 })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestDifferenceBy_Predicate", func(t *testing.T) {
		slice := []string{"a", "aa", "aaa"}
		comparedSlice := []string{"aa", "value"}
		expected := []string{"a", "aaa", "value"}
		result := DifferenceBy(slice, comparedSlice, func(s1, s2 string) bool {
			return len(s1) == len(s2)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestGroup(t *testing.T) {
	t.Run("TestGroup_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7}
		size := 3
		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
		result := Group(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestGroup_StringSlice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d"}
		size := 2
		expected := [][]string{{"a", "b"}, {"c", "d"}}
		result := Group(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestGroup_ZeroSize", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		size := 0
		var expected [][]int

		result := Group(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestGroup_EmptySlice", func(t *testing.T) {
		var slice []int
		size := 3
		var expected [][]int
		result := Group(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestPrepend(t *testing.T) {
	t.Run("TestPrepend_IntSlice", func(t *testing.T) {
		slice := []int{2, 3, 4}
		element := 1
		expected := []int{1, 2, 3, 4}
		result := Prepend(slice, element)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestPrepend_StringSlice", func(t *testing.T) {
		slice := []string{"b", "c"}
		element := "a"
		expected := []string{"a", "b", "c"}
		result := Prepend(slice, element)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestPrepend_EmptySlice", func(t *testing.T) {
		slice := []int{}
		element := 1
		expected := []int{1}
		result := Prepend(slice, element)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("TestReverse_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		Reverse(slice)
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected %v, but got %v", expected, slice)
		}
	})

	t.Run("TestReverse_StringSlice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d"}
		expected := []string{"d", "c", "b", "a"}
		Reverse(slice)
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected %v, but got %v", expected, slice)
		}
	})

	t.Run("TestReverse_One", func(t *testing.T) {
		slice := []string{"a"}
		expected := []string{"a"}
		Reverse(slice)
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected %v, but got %v", expected, slice)
		}
	})

	t.Run("TestReverse_Nil", func(t *testing.T) {
		var slice []string
		var expected []string
		Reverse(slice)
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected %v, but got %v", expected, slice)
		}
	})
}

func TestInsertAt(t *testing.T) {
	t.Run("TestInsertAt_InsertSingleElement", func(t *testing.T) {
		slice := []int{1, 2, 3}
		index := 1
		value := 4
		expected := []int{1, 4, 2, 3}
		result := InsertAt(slice, index, value)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestInsertAt_InsertSlice", func(t *testing.T) {
		slice := []string{"a", "b"}
		index := 0
		value := []string{"x", "y"}
		expected := []string{"x", "y", "a", "b"}
		result := InsertAt(slice, index, value)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestInsertAt_InvalidIndex", func(t *testing.T) {
		slice := []int{1, 2, 3}
		index := 5
		value := 4
		expected := []int{1, 2, 3}
		result := InsertAt(slice, index, value)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestFindFirstBy(t *testing.T) {
	t.Run("TestFindFirstBy_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		expectedValue := 3
		resultValue, resultOK := FindFirstBy(slice, func(i int, v int) bool {
			return v > 2
		})
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFindFirstBy_StringSlice", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		expectedValue := "banana"
		resultValue, resultOK := FindFirstBy(slice, func(i int, v string) bool {
			return v == "banana"
		})
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFindFirstBy_ValueSlice", func(t *testing.T) {
		type Value struct {
			Name  string
			Color string
		}

		slice := []Value{
			{"apple", "red"},
			{"banana", "yellow"},
			{"cherry", "red"},
		}
		expectedValue := Value{"apple", "red"}
		resultValue, resultOK := FindFirstBy(slice, func(i int, v Value) bool {
			return v.Color == "red"
		})
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFindFirstBy_EmptySlice", func(t *testing.T) {
		var slice []int
		expectedValue := 0
		resultValue, resultOK := FindFirstBy(slice, func(i int, v int) bool {
			return v > 2
		})
		if resultValue != expectedValue || resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, false, resultValue, resultOK)
		}
	})
}

func TestFindLastBy(t *testing.T) {
	t.Run("TestFindLastBy_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		expectedValue := 4
		resultValue, resultOK := FindLastBy(slice, func(i int, v int) bool {
			return v > 2
		})
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFindLastBy_ValueSlice", func(t *testing.T) {
		type Value struct {
			Name  string
			Color string
		}

		slice := []Value{
			{"apple", "red"},
			{"banana", "yellow"},
			{"cherry", "red"},
		}
		expectedValue := Value{"cherry", "red"}
		resultValue, resultOK := FindLastBy(slice, func(i int, v Value) bool {
			return v.Color == "red"
		})
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFindLastBy_StringSlice", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		expectedValue := "cherry"
		resultValue, resultOK := FindLastBy(slice, func(i int, v string) bool {
			return len(v) > 5
		})
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFindLastBy_EmptySlice", func(t *testing.T) {
		var slice []int = nil
		expectedValue := 0
		resultValue, resultOK := FindLastBy(slice, func(i int, v int) bool {
			return v > 2
		})
		if resultValue != expectedValue || resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, false, resultValue, resultOK)
		}
	})
}

func TestFirst(t *testing.T) {
	t.Run("TestFirst_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		expectedValue := 1
		resultValue, resultOK := First(slice)
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFirst_StringSlice", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		expectedValue := "apple"
		resultValue, resultOK := First(slice)
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestFirst_EmptySlice", func(t *testing.T) {
		slice := []int{}
		expectedValue := 0
		resultValue, resultOK := First(slice)
		if resultValue != expectedValue || resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, false, resultValue, resultOK)
		}
	})

	t.Run("TestFirst_NilSlice", func(t *testing.T) {
		var slice []int
		expectedValue := 0
		resultValue, resultOK := First(slice)
		if resultValue != expectedValue || resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, false, resultValue, resultOK)
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("TestLast_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		expectedValue := 3
		resultValue, resultOK := Last(slice)
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestLast_StringSlice", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		expectedValue := "cherry"
		resultValue, resultOK := Last(slice)
		if resultValue != expectedValue || !resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, true, resultValue, resultOK)
		}
	})

	t.Run("TestLast_EmptySlice", func(t *testing.T) {
		slice := []int{}
		expectedValue := 0
		resultValue, resultOK := Last(slice)
		if resultValue != expectedValue || resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, false, resultValue, resultOK)
		}
	})

	t.Run("TestLast_NilSlice", func(t *testing.T) {
		var slice []int
		expectedValue := 0
		resultValue, resultOK := Last(slice)
		if resultValue != expectedValue || resultOK {
			t.Errorf("Expected value %v and ok %v, but got value %v and ok %v", expectedValue, false, resultValue, resultOK)
		}
	})
}

func TestRandomIn(t *testing.T) {
	t.Run("TestRandomIn_IntSlice", func(t *testing.T) {
		originalSlice := []int{1, 2, 3, 4, 5}
		slice := make([]int, len(originalSlice))
		copy(slice, originalSlice)
		RandomIn(slice)
		if len(slice) != len(originalSlice) {
			t.Errorf("Expected shuffled goexslice length %v, but got %v", len(originalSlice), len(slice))
		}
		if reflect.DeepEqual(slice, originalSlice) {
			t.Errorf("Expected shuffled goexslice to be different from original")
		}
	})
}

func TestRandomCopy(t *testing.T) {
	t.Run("TestRandomCopy_IntSlice", func(t *testing.T) {
		originalSlice := []int{1, 2, 3, 4, 5}
		slice := make([]int, len(originalSlice))
		copy(slice, originalSlice)

		resultSlice := RandomCopy(slice)
		if len(slice) != len(slice) {
			t.Errorf("Expected shuffled goexslice length %v, but got %v", len(slice), len(resultSlice))
		}
		if reflect.DeepEqual(resultSlice, originalSlice) {
			t.Errorf("Expected shuffled goexslice to be different from original")
		}
	})
}

func TestEnumerated(t *testing.T) {
	t.Run("Enumerated_IntSlice", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		result := Enumerated(slice)
		if len(result) != len(slice) {
			t.Errorf("Expected indices slice length %v, but got %v", len(slice), len(result))
		}
		for i, tuple := range result {
			if tuple.S1 != slice[i] || tuple.S2 != i {
				t.Errorf("Expected tuple %v to be {S1: %v, S2: %v}, but got {S1: %v, S2: %v}", i, slice[i], i, tuple.S1, tuple.S2)
			}
		}
	})
}

func TestSafeSwap(t *testing.T) {
	t.Run("TestSafeSwap_Success", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		expectedSlice := []int{30, 20, 10, 40}
		ok := SafeSwap(slice, 0, 2)
		if !ok {
			t.Errorf("Expected SafeSwap to return true for successful swap")
		}
		if !reflect.DeepEqual(slice, expectedSlice) {
			t.Errorf("Expected slice after swap to be %v, but got %v", expectedSlice, slice)
		}
	})

	t.Run("TestSafeSwap_SameIndices", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		originalSlice := make([]int, len(slice))
		copy(originalSlice, slice)
		ok := SafeSwap(slice, 1, 1)
		if ok {
			t.Errorf("Expected SafeSwap to return false for same indices")
		}
		if !reflect.DeepEqual(slice, originalSlice) {
			t.Errorf("Expected slice to remain unchanged for same indices")
		}
	})

	t.Run("TestSafeSwap_InvalidIndices", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		originalSlice := make([]int, len(slice))
		copy(originalSlice, slice)
		ok1 := SafeSwap(slice, -1, 2)
		ok2 := SafeSwap(slice, 0, 4)
		if ok1 || ok2 {
			t.Errorf("Expected SafeSwap to return false for invalid indices")
		}
		if !reflect.DeepEqual(slice, originalSlice) {
			t.Errorf("Expected slice to remain unchanged for invalid indices")
		}
	})
}
func TestSafeIndex(t *testing.T) {
	t.Run("TestSafeIndex_ValidIndex", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		v, ok := SafeIndex(slice, 2)
		if !ok {
			t.Errorf("Expected SafeIndex to return true for valid index")
		}
		if v != slice[2] {
			t.Errorf("Expected value at index 2 to be %v, but got %v", slice[2], v)
		}
	})

	t.Run("TestSafeIndex_InvalidIndex", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		v, ok := SafeIndex(slice, -1)
		if ok {
			t.Errorf("Expected SafeIndex to return false for invalid index")
		}
		if v != 0 {
			t.Errorf("Expected value for invalid index to be 0, but got %v", v)
		}
		v, ok = SafeIndex(slice, 4)
		if ok {
			t.Errorf("Expected SafeIndex to return false for invalid index")
		}
		if v != 0 {
			t.Errorf("Expected value for invalid index to be 0, but got %v", v)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("TestIsEmpty_EmptySlice", func(t *testing.T) {
		emptySlice := []int{}
		if !IsEmpty(emptySlice) {
			t.Errorf("Expected IsEmpty to return true for empty slice")
		}
	})

	t.Run("TestIsEmpty_NilSlice", func(t *testing.T) {
		var emptySlice []int
		if !IsEmpty(emptySlice) {
			t.Errorf("Expected IsEmpty to return false for non-empty slice")
		}
	})

	t.Run("TestIsEmpty_NonEmptySlice", func(t *testing.T) {
		nonEmptySlice := []int{10, 20, 30}
		if IsEmpty(nonEmptySlice) {
			t.Errorf("Expected IsEmpty to return false for non-empty slice")
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("TestFilter_KeepEvenNumbers", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6}
		evenPredicate := func(value int) bool {
			return value%2 == 0
		}
		expectedResult := []int{2, 4, 6}
		result := Filter(slice, evenPredicate)
		if len(result) != len(expectedResult) {
			t.Errorf("Expected filtered slice length %v, but got %v", len(expectedResult), len(result))
		}
		for i, val := range result {
			if val != expectedResult[i] {
				t.Errorf("Expected value at index %v to be %v, but got %v", i, expectedResult[i], val)
			}
		}
	})

	t.Run("TestFilter_EmptySlice", func(t *testing.T) {
		emptySlice := []int{}
		anyPredicate := func(value int) bool {
			return true
		}
		result := Filter(emptySlice, anyPredicate)
		if !IsEmpty(result) {
			t.Errorf("Expected filtered slice length 0 for empty input slice, but got %v", len(result))
		}
	})

	t.Run("TestFilter_NilPredicate", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6}
		result := Filter(slice, nil)
		if !ContainsAll(slice, result...) || len(result) != len(slice) {
			t.Errorf("Expected filtered slice length %v for nil predicate, but got %v", len(slice), len(result))
		}
	})
}
