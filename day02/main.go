package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidSequence(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	increasing := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if increasing {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || diff < -3 {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	reports := strings.Split(strings.TrimSpace(string(data)), "\n")
	count := 0

	for _, report := range reports {
		if report == "" {
			continue
		}

		strNums := strings.Fields(report)
		nums := make([]int, len(strNums))
		for i, v := range strNums {
			nums[i], _ = strconv.Atoi(v)
		}

		if isValidSequence(nums) {
			count++
			continue
		}

		for i := 0; i < len(nums); i++ {
			dampened := make([]int, 0, len(nums)-1)
			dampened = append(dampened, nums[:i]...)
			dampened = append(dampened, nums[i+1:]...)
			
			if isValidSequence(dampened) {
				count++
				break
			}
		}
	}

	fmt.Println("Count:", count)
}