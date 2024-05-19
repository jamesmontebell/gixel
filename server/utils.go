package server

func calculateLevel(email string) int {
	level, err := dbConnection.Exec("SELECT level FROM Character WHERE user_email=?", email)
	if err != nil {
		return 0
	}

	exp, err := dbConnection.Exec("SELECT experience FROM Character WHERE user_email=?", email)
	if err != nil {
		return 0
	}

	experienceNeeded := (4*level ^ 3) / 5

	if exp >= experienceNeeded {

	}

}
