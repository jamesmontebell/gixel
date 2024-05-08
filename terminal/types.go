package main

type Outputs struct {
	Changes   int
	Inserts   int
	Deletions int
}

type TestCases []struct {
	Input    string
	Expected int
}
