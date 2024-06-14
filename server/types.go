package main

import "io"

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

type CalculateLevelTestCases []struct {
	GainedExperience int
	CurrentLevel     int
	CurrentExp       int
	ExpectedLevel    int
	ExpectedExp      int
}

type RequestTest struct {
	Method string
	Path   string
	Body   io.Reader
	Code   int
}
