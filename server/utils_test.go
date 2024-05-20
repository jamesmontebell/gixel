package main

import (
	"testing"
)

func TestCalculateLevel(t *testing.T) {
	testCases := CalculateLevelTestCases{
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

func TestRealCalculateLevel(t *testing.T) {
	e := Experience{
		UserEmail: "test",
		Exp:       1,
	}

	expResult, levelResult, err := calculateLevel(e)
	if err != nil {
		t.Error("calculateLevel error", err)
	}

	if expResult != 0 {
		t.Errorf("calculateLevel returned %d, expected %d", expResult, 0)
	}
	if levelResult != 2 {
		t.Errorf("calculateLevel returned %d, expected %d", levelResult, 2)
	}
}
