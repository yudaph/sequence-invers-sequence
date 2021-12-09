package search_test

import (
	"fmt"
	"slice/search"
	"testing"
)

func TestSearch(t *testing.T) {
	result := []int{-1, 2}
	testVariables := []struct {
		slice  []int
		result *int
	}{
		{
			slice:  []int{-1, -2, -3, 8, 9, -3, -2, -1},
			result: &result[0],
		},
		{
			slice:  []int{1, 2, 3, 8, 9, 2, 1},
			result: nil,
		},
		{
			slice:  []int{7, 1, 2, 9, 7, 2, 1},
			result: &result[1],
		},
	}

	for i, testVariable := range testVariables {
		result := search.MaxValueSequenceAndSequenceInvers(testVariable.slice)
		if result == nil || testVariable.slice == nil {
			// if no match (result <nil>) the expected result should be <nil>
			if result != testVariable.result {
				t.Fatalf("Failed on test %x\n", i)
			}
		} else if *result != *testVariable.result {
			t.Fatalf("Failed on test %x\n", i)
		}
		fmt.Printf("test %x passed\n", i)
	}
}
