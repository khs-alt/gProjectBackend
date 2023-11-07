package models

type TestCodeData struct {
	TestCode string `json:"testcode"`
}

type TagData struct {
	Tag string `json:"tag"`
}

type RequestData struct {
	Tags []string `json:"tags"`
}

// Define a struct to represent the data you want to fetch

type UserCurrentScore struct {
	Score int `json:"score"`
}

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

type UserVideoInitInfo struct {
	CurrentPage           string   `json:"currentPage"`
	VideoList             []string `json:"videoList"`
	OriginalVideoNameList []string `json:"originalVideoNameList"`
	ArtifactVideoNameList []string `json:"artifactVideoNameList"`
	OriginalVideoFPSList  []string `json:"originalVideoFPSList"`
	ArtifactVideoFPSList  []string `json:"artifactVideoFPSList"`
}

type UserInfoData struct {
	CurrentUser string
	ImageId     int
	TestCode    string
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

var MimeTypes = map[string]string{
	".mov":   "video/quicktime",
	".wmv":   "video/x-ms-wmv",
	".avi":   "video/x-msvideo",
	".avchd": "video/x-mts",
	".flv":   "video/x-flv",
	".f4v":   "video/x-f4v",
	".swf":   "application/x-shockwave-flash",
	".mkv":   "video/x-matroska",
	".webm":  "video/webm",
	".mp4":   "video/mp4",
}
