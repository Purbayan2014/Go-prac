package main

import (
	"Go-prac/Testing"
	"testing"
)

// benchmarking
func BenchmarkMySum(t *testing.B) {
	type test struct {
		data []int
		ans  int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{2, 2}, 4},
		test{[]int{4, 2}, 6},
		test{[]int{1, 3}, 4},
	}
	for _, v := range tests {
		x := Testing.MySum(v.data...)
		if x != v.ans {
			t.Error("Expected ", v.ans, " got ", x)
		}
	}

}

func TestMySum(t *testing.T) {

	// table test
	type test struct {
		data []int
		ans  int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{2, 2}, 4},
		test{[]int{4, 2}, 6},
		test{[]int{1, 3}, 4},
	}

	// range over the tests
	for _, v := range tests {
		x := Testing.MySum(v.data...)
		if x != v.ans {
			t.Error("Expected ", v.ans, " got ", x)
		}
	}

	// testing with one value
	x := Testing.MySum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, " Got : ", x)
	}
}
