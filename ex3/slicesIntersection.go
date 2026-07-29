package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intersection(input1, input2 string) string {
	line1, line2 := strings.Fields(input1), strings.Fields(input2)
	mapLine := make(map[int]bool)

	for _, v := range line2 {
		valInt, err := strconv.Atoi(v)
		if err != nil {
			return "Invalid input"
		}
		mapLine[valInt] = true
	}

	result := make([]int, 0, len(mapLine)/2)
	resultMap := make(map[int]bool)
	for _, v := range line1 {
		valInt, err := strconv.Atoi(v)
		if err != nil {
			return "Invalid input"
		}
		if mapLine[valInt] && !resultMap[valInt] {
			result = append(result, valInt)
			resultMap[valInt] = true
		}
	}

	if len(result) == 0 {
		return "Empty intersection"
	}

	var res string
	for i := 0; i < len(result); i++ {
		res += fmt.Sprintf("%d", result[i])
		if i == len(result)-1 {
			res += fmt.Sprintf("")
		} else {
			res += fmt.Sprintf(" ")
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := scanner.Text()
	scanner.Scan()
	input2 := scanner.Text()

	fmt.Println(intersection(input1, input2))

}
