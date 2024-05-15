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
