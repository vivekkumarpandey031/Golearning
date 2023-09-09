package slice1

import (
	"testing"
)

func TestSumSlice1(t *testing.T) {
	input := []int{2, 1, 3, 4, 5, 6}
	expectedOutput := 21
	actualOutput := getSum(input)

	if expectedOutput != actualOutput {
		t.Fail()
	}

}
