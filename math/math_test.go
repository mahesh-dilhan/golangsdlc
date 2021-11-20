package math

import (
	"fmt"
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

var table = []struct {
	x    int
	y    int
	want int
}{
	{2, 2, 4},
	{5, 3, 8},
	{8, 4, 12},
	{12, 5, 17},
}

func TestSumParalel(t *testing.T) {
	t.Parallel()
	for _, row := range table {
		testName := fmt.Sprintf("Test %d+%d", row.x, row.y)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got := Add(row.x, row.y)
			if got != row.want {
				t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
			}
		})
	}
}

//Table Test (Normal)
func TestSum(t *testing.T) {
	for _, row := range table {
		got := Add(row.x, row.y)
		if got != row.want {
			t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
		}
	}
}

//Table Test (With Subtest)
func TestSumSubtest(t *testing.T) {
	for _, row := range table {
		testName := fmt.Sprintf("Test %d+%d", row.x, row.y)
		t.Run(testName, func(t *testing.T) {
			got := Add(row.x, row.y)
			if got != row.want {
				t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
			}
		})
	}
}
