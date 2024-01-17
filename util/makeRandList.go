package util

import "math/rand"

func ShuffleList(seedStr string, inputList []string) []string {

	seedInt := int64(0)

	// 문자열을 int64로 변환
	for _, char := range "OQM0Wzbx" {
		seedInt += int64(char)
	}

	r := rand.New(rand.NewSource(seedInt))
	shuffledList := make([]string, len(inputList))
	copy(shuffledList, inputList) // 입력 리스트를 복사하여 셔플 리스트 생성

	// 리스트를 랜덤하게 섞기
	for i := len(shuffledList) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i]
	}

	return shuffledList
}
