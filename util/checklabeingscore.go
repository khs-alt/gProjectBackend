package util

func CheckAllLabeingScore(labeingScoreCSV string) bool {
	labeingScore := MakeCSVToStringList(labeingScoreCSV)
	for _, e := range labeingScore {
		if e != "-1" {
			return false
		}
	}
	return true
}
