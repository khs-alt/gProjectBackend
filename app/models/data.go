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
	//CurrentMode string `json:"current_mode"`
	ID       string `json:"user_id"`
	Password string `json:"user_password"`
	TestCode string `json:"test_code"`
}

type UserVideoInitInfo struct {
	CurrentPage           string   `json:"currentPage"`
	VideoList             []string `json:"videoList"`
	OriginalVideoNameList []string `json:"originalVideoNameList"`
	ArtifactVideoNameList []string `json:"artifactVideoNameList"`
	OriginalVideoFPSList  []string `json:"originalVideoFPSList"`
	ArtifactVideoFPSList  []string `json:"artifactVideoFPSList"`
	UserScore             int      `json:"userScore"`
}

type UserInfoData struct {
	CurrentUser string
	ImageId     int
	TestCode    string
}

type UserImageScoreData struct {
	CurrentUser string `json:"current_user"`
	ImageId     int    `json:"image_id"`
	TestCode    string `json:"test_code"`
	Score       []int  `json:"score"`
}

type UserScoreData struct {
	Title       string
	CurrentUser string
	TestCode    string
	ImageId     int
	Score       int
}

type UserImageInfo struct {
	CurrentUser string `json:"current_user"`
	ImageID     int    `json:"image_id"`
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

type VideoFrameTimeData struct {
	VideoIndex       int    `json:"videoIndex"`
	VideoCurrentTime string `json:"selectedVideoTime"`
}

type UserScoringListData struct {
	CurrentUser string `json:"userID"`
	TestCode    string `json:"Testcode"`
}

type UserLabelingListData struct {
	UserID   string `json:"user_id"`
	TestCode string `json:"testcode"`
}
