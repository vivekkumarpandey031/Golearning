package slice1

import "testing"

func TestSumSlice1(t *testing.T)  {
	input:= {2,1,3,4,5,6}
	expectedOutput:=21
	actualOutput:=getSum(slice1)

	if expectedOutput !=actualOutput{
		t.fail()
	}
	
}