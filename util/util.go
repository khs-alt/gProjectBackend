package util

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CountFile(s string) (int, error) {
	fmt.Print(s)
	files, err := os.ReadDir(s)
	if err != nil {
		return 0, err
	}

	fileCount := 0

	for _, file := range files {
		if !file.IsDir() {
			fileCount++
		}
	}
	fmt.Println(fileCount)
	return fileCount, nil
}

func MakeNumber(s []string) []int {
	var newArray []int
	for _, video := range s {
		num := strings.TrimLeft(video, "originalVideo")
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		newArray = append(newArray, n)
	}
	return newArray
}
func MakeCSVtoIntList(s string) []int {
	// 입력 문자열을 쉼표로 분할하여 문자열 슬라이스로 변환
	stringList := strings.Split(s, ",")

	// 결과로 반환할 정수 목록 슬라이
	var intList []int

	// 문자열 슬라이스를 정수 슬라이스로 변환
	for _, str := range stringList {
		// 문자열을 정수로 변환
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		// 정수 슬라이스에 추가
		intList = append(intList, num)
	}

	return intList
}

func MakeIntListtoCSV(numArray []int) string {
	s := ""
	for _, num := range numArray {
		s = s + fmt.Sprint(num) + ","
	}
	s = strings.TrimRight(s, ",")
	return s
}

func MakeStringListtoCSV(str []string) string {
	sort.Strings(str)
	s := ""
	for _, word := range str {
		s = s + word + ","
	}
	s = strings.TrimRight(s, ",")
	return s
}

func MakeCSVToStringList(s string) []string {

	// 입력 문자열을 쉼표로 분할하여 문자열 슬라이스로 변환
	stringList := strings.Split(s, ",")

	//doing sort
	sort.Strings(stringList)
	// 결과로 반환할 정수 목록 슬라이
	return stringList
}

func FindMostFrequentElement(arr []string) string {
	// 맵을 사용하여 요소의 등장 횟수를 세기
	elementCount := make(map[string]int)
	maxCount := 0
	mostFrequentElement := ""

	for _, element := range arr {
		elementCount[element]++
		if elementCount[element] > maxCount {
			maxCount = elementCount[element]
			mostFrequentElement = element
		}
	}

	return mostFrequentElement
}

func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do nothing (duplicate element)
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
