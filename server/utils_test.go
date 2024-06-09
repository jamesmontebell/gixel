package main

import (
	"testing"
)

func TestCalculateLevel(t *testing.T) {
	testCases := CalculateLevelTestCases{
		{100, 1, 0, 101, 0},
		{100, 5, 0, 6, 0},
		{1, 100, 0, 100, 1},
		{1, 100, 799999, 101, 0},
		{0, 100, 799999, 100, 799999},
		{1, 100, 799998, 100, 799999},
	}

	for _, tc := range testCases {
		expResult, levelResult := calculateLevelTest(tc.GainedExperience, tc.CurrentExp, tc.CurrentLevel)

		if expResult != tc.ExpectedExp {
			t.Errorf("calculateLevelTest returned %d, expected %d", expResult, tc.ExpectedExp)
		}
		if levelResult != tc.ExpectedLevel {
			t.Errorf("calculateLevelTest returned %d, expected %d", levelResult, tc.ExpectedLevel)
		}
	}
}
