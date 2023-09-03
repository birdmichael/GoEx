package slice

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

func TestChunk(t *testing.T) {
	t.Run("TestChunk_IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7}
		size := 3
		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
		result := Chunk(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestChunk_StringSlice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d"}
		size := 2
		expected := [][]string{{"a", "b"}, {"c", "d"}}
		result := Chunk(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestChunk_ZeroSize", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		size := 0
		var expected [][]int

		result := Chunk(slice, size)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("TestChunk_EmptySlice", func(t *testing.T) {
		var slice []int
		size := 3
		var expected [][]int
		result := Chunk(slice, size)
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
