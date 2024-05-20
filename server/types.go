package main

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
