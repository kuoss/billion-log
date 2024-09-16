package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func getLoopCountsFromEnv() []int {
	defaultCounts := []int{1000, 1000}
	envValue := os.Getenv("LOOP_COUNTS")
	if envValue == "" {
		fmt.Println("LOOP_COUNTS not set, using default:", defaultCounts)
		return defaultCounts
	}

	countStrings := strings.Split(envValue, ",")
	loopCounts := make([]int, len(countStrings))
	for i, countStr := range countStrings {
		count, err := strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			fmt.Printf("Invalid LOOP_COUNTS value at index %d: %s, using default: %v\n", i, countStr, defaultCounts)
			return defaultCounts
		}
		loopCounts[i] = count
	}
	return loopCounts
}

func getDepthPrefix(depth int) string {
	if depth >= 0 && depth < 26 {
		return string('a' + depth)
	}
	return "?"
}

func getZeroPaddedFormat(loopSize int) string {
	digits := len(strconv.Itoa(loopSize - 1))
	return "%0" + strconv.Itoa(digits) + "d"
}

func generateLoops(loopCounts []int, depth int, currentIndices []string) {
	zeroPaddedFormat := getZeroPaddedFormat(loopCounts[depth])
	prefix := getDepthPrefix(depth)

	for i := 0; i < loopCounts[depth]; i++ {
		formattedIndex := fmt.Sprintf("%s=%s", prefix, fmt.Sprintf(zeroPaddedFormat, i))
		newIndices := append(currentIndices, formattedIndex)

		if depth == len(loopCounts)-1 {
			fmt.Printf("%s uuid=%s\n", strings.Join(newIndices, " "), uuid.New())
		} else {
			generateLoops(loopCounts, depth+1, newIndices)
		}
	}
}

func main() {
	loopCounts := getLoopCountsFromEnv()
	generateLoops(loopCounts, 0, []string{})
}
