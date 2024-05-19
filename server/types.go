package server

type Experience struct {
	UserEmail string `json:"userEmail"`
	Exp       int    `json:"experience"`
}

type Character struct {
	User_id    int
	User_email string
	Name       string
	Level      int
	Experience int
}
