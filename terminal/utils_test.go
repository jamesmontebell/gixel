package main

import (
	"fmt"
	"testing"
)

func TestFindFilesChanged(t *testing.T) {

	testCases := TestCases{
		{Input: "[main93b2f56]parser123fileschanged,22insertions(+),1deletion(-)", Expected: 123},
		{Input: "[main93b2f56]parser12fileschanged,22insertions(+),1deletion(-)", Expected: 12},
		{Input: "[main93b2f56]parser1fileschanged,22insertions(+),1deletion(-)", Expected: 1},
		{Input: "[main93b2f56]parser1filechanged,22insertions(+),1deletion(-)", Expected: 1},
	}

	for _, tc := range testCases {
		result, err := findFilesChanged(tc.Input)
		if err != nil {
			fmt.Println("error: ", err)
			t.Errorf("test error")
		}
		if result != tc.Expected {
			t.Errorf("findFilesChanged(%s) returned %d, expected %d", tc.Input, result, tc.Expected)
		}
	}
}

func TestFindInsertions(t *testing.T) {
	testCases := TestCases{
		{Input: "[main93b2f56]parser123filechanged,1234insertions(+),1deletion(-)", Expected: 1234},
		{Input: "[main93b2f56]parser123filechanged,123insertions(+),1deletion(-)", Expected: 123},
		{Input: "[main93b2f56]parser123filechanged,12insertions(+),1deletion(-)", Expected: 12},
		{Input: "[main93b2f56]parser123filechanged,1insertions(+),1deletion(-)", Expected: 1},
		{Input: "[main93b2f56]parser123filechanged,1insertion(+),1deletion(-)", Expected: 1},
	}

	for _, tc := range testCases {
		result, err := findInsertions(tc.Input)
		if err != nil {
			fmt.Println("error: ", err)
			t.Errorf("test error")
		}
		if result != tc.Expected {
			t.Errorf("findInsertions(%s) returned %d, expected %d", tc.Input, result, tc.Expected)
		}
	}
}

func TestFindDeletions(t *testing.T) {
	testCases := TestCases{
		{Input: "[main93b2f56]parser123filechanged,22insertions(+),1234deletions(-)", Expected: 1234},
		{Input: "[main93b2f56]parser123filechanged,22insertions(+),123deletions(-)", Expected: 123},
		{Input: "[main93b2f56]parser123filechanged,22insertions(+),12deletions(-)", Expected: 12},
		{Input: "[main93b2f56]parser123filechanged,22insertions(+),1deletions(-)", Expected: 1},
		{Input: "[main93b2f56]parser123filechanged,22insertions(+),1deletion(-)", Expected: 1},
	}

	for _, tc := range testCases {
		result, err := findDeletions(tc.Input)
		if err != nil {
			fmt.Println("error: ", err)
			t.Errorf("test error")
		}
		if result != tc.Expected {
			t.Errorf("findDeletions(%s) returned %d, expected %d", tc.Input, result, tc.Expected)
		}
	}
}
