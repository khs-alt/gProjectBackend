package models

type RequestData struct {
	Tags []string `json:"tags"`
}

// Define a struct to represent the data you want to fetch

type UserNewIdAndPassword struct {
	NewId       string `json:"newId"`
	NewPassword string `json:"newPassword"`
}

type UserIdAndPassword struct {
	Id       string `json:"adminId"`
	Password string `json:"adminPassword"`
}

type UserLoginData struct {
	ID       string `json:"userID"`
	Password string `json:"userPassword"`
	TestCode string `json:"testcode"`
}

type UserScoreData struct {
	Title       string
	CurrentUser string
	TestCode    string
	ImageId     int
	Score       int
}
type UserIdPassword struct {
	UserId       string
	UserPassword string
}
