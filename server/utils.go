package main

import (
	"errors"
	"math"
)

func calculateLevel(e Experience) (int, int, error) {

	exp, err := getExperience(e.UserEmail)
	if err != nil {
		return 0, 0, errors.New("error getExperience")
	}

	level, err := getLevel(e.UserEmail)
	if err != nil {
		return 0, 0, errors.New("error getLevel")
	}

	experienceNeeded := (4.0 * math.Pow(float64(level), 3.0)) / 5.0
	experienceNeeded = math.RoundToEven(experienceNeeded)

	for float64(exp) >= experienceNeeded {
		level++
		exp = exp - int(experienceNeeded)
	}

	return exp, level, nil
}

func getExperience(userEmail string) (int, error) {
	var exp int
	err := dbConnection.QueryRow("SELECT experience FROM Character WHERE user_email=?", userEmail).Scan(&exp)
	if err != nil {
		return 0, errors.New("db error, experience")
	}

	return exp, nil
}

func getLevel(userEmail string) (int, error) {
	var level int
	err := dbConnection.QueryRow("SELECT level FROM Character WHERE user_email=?", userEmail).Scan(&level)
	if err != nil {
		return 0, errors.New("db error, level")
	}

	return level, nil
}

func calculateLevelTest(gained int, exp int, level int) (int, int) {
	exp = exp + gained

	experienceNeeded := (4.0 * math.Pow(float64(level), 3.0)) / 5.0
	experienceNeeded = math.RoundToEven(experienceNeeded)

	for float64(exp) >= experienceNeeded {
		level++
		exp = exp - int(experienceNeeded)
	}

	return exp, level
}
