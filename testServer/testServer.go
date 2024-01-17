package testServer

import (
	"fmt"
	"net/http"
	"slices"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	concurrentRequests := 145 // 동시 요청 수
	var result []string
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			startTime := time.Now()
			response, err := http.Get("https://pilab.smu.ac.kr/label/")
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", i, err)
				return
			}
			defer response.Body.Close()
			duration := time.Since(startTime)
			result = append(result, fmt.Sprintf("Request %d: %v\n", i, duration))
		}(i)
	}

	wg.Wait()
	slices.Sort(result)
}
