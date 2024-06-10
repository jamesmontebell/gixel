package main

import (
	"testing"
)

func TestCalculateLevel(t *testing.T) {
	testCases := CalculateLevelTestCases{
		{
			GainedExperience: 10,
			CurrentLevel:     1,
			CurrentExp:       0,
			ExpectedLevel:    3,
			ExpectedExp:      6,
		},
		{
			GainedExperience: 20,
			CurrentLevel:     2,
			CurrentExp:       2,
			ExpectedLevel:    4,
			ExpectedExp:      9,
		},
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
