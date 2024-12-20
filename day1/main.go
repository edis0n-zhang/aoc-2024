package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)
func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	rows := strings.Split(string(file_text), "\n")
	for _, row := range rows {
		if len(row) > 0 {
			fields := strings.Fields(row)
			num1, _ := strconv.Atoi(fields[0])
			num2, _ := strconv.Atoi(fields[1])
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	dist_sum := 0

	for i := 0; i < len(list1); i++ {
		dist_sum += int(math.Abs(float64(list2[i] - list1[i])))
	}

	fmt.Println("Total Distance:", dist_sum)

	count_map := make(map[int]int)

	for i:= 0; i < len(list2); i++ {
		if _, ok := count_map[list2[i]]; !ok {
			count_map[list2[i]] = 1
		} else {
			count_map[list2[i]] += 1
		}
	}

	similarity_score := 0

	for i := 0; i < len(list1); i++ {
		if val, ok := count_map[list1[i]]; !ok {
			continue
		} else {
			similarity_score += val * list1[i]
		}
	}

	fmt.Println("Similarity Score:", similarity_score)
}