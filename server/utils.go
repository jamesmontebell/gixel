package main

import (
	"errors"
	"math"
)

// Retrieves a users current level and experience then calculates new level based on
// the current amount of experience and gained experiences
func calculateLevel(e Experience) (int, int, error) {
	exp, err := getExperience(e.UserEmail)
	if err != nil {
		return 0, 0, errors.New("error getExperience")
	}

	level, err := getLevel(e.UserEmail)
	if err != nil {
		return 0, 0, errors.New("error getLevel")
	}

	exp += e.Exp

	calculateExperienceNeeded := func(level int) int {
		return int(math.RoundToEven((4.0 * math.Pow(float64(level), 3.0)) / 50.0))
	}

	experienceNeeded := calculateExperienceNeeded(level)

	for float64(exp) >= float64(experienceNeeded) {
		exp -= experienceNeeded
		level++
		experienceNeeded = calculateExperienceNeeded(level)
	}

	return exp, level, nil
}

// Retrieves and returns a character's current experience
func getExperience(userEmail string) (int, error) {
	var exp int
	err := dbConnection.QueryRow("SELECT experience FROM Character WHERE user_email=?", userEmail).Scan(&exp)
	if err != nil {
		return 0, errors.New("db error, experience")
	}

	return exp, nil
}

// Retrieves and returns a character's current level
func getLevel(userEmail string) (int, error) {
	var level int
	err := dbConnection.QueryRow("SELECT level FROM Character WHERE user_email=?", userEmail).Scan(&level)
	if err != nil {
		return 0, errors.New("db error, level")
	}

	return level, nil
}

// Test function to test experience gained/new level algorithm
func calculateLevelTest(gained int, exp int, level int) (int, int) {
	exp = exp + gained

	experienceNeeded := (4.0 * math.Pow(float64(level), 3.0)) / 50.0
	experienceNeeded = math.RoundToEven(experienceNeeded)

	for float64(exp) >= experienceNeeded {
		level++
		exp = exp - int(experienceNeeded)
	}

	return exp, level
}
