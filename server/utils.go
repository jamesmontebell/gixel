package server

import "errors"

func calculateLevel(e Experience) (int, int, error) {

	exp, err := getExperience(e.UserEmail)
	if err != nil {
		return 0, 0, errors.New("error getExperience")
	}

	level, err := getLevel(e.UserEmail)
	if err != nil {
		return 0, 0, errors.New("error getLevel")
	}

	exp = exp + e.Exp
	experienceNeeded := (4*level ^ 3) / 5

	for exp >= experienceNeeded {
		level++
		exp = exp - experienceNeeded
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
