package main

type ParsedOutputs struct {
	Changes   int
	Inserts   int
	Deletions int
}

type TestCases []struct {
	Input    string
	Expected int
}

type Experience struct {
	UserEmail string `json:"userEmail"`
	Exp       int    `json:"experience"`
}

type Character struct {
	User_id    int    `json:"userId"`
	User_email string `json:"userEmail"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Experience int    `json:"experience"`
}
