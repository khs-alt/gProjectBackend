package util

func MakeNewTestCodeTagList(testCodeList []string, tagList []string) ([]string, []string) {
	newTestCodeList := make([]string, 0)
	newTagList := make([]string, 0)
	if len(testCodeList) == 0 {
		return newTestCodeList, newTagList
	}
	newTestCodeList = append(newTestCodeList, testCodeList[0])
	newTagList = append(newTagList, tagList[0])
	for i := 1; i < len(testCodeList); i++ {
		if testCodeList[i] == testCodeList[i-1] {
			newTagList[len(newTagList)-1] += "," + tagList[i]
			continue
		}
		newTestCodeList = append(newTestCodeList, testCodeList[i])
		newTagList = append(newTagList, tagList[i])
	}
	return newTestCodeList, newTagList
}
