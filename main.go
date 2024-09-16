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
		return defaultCounts
	}

	countStrings := strings.Split(envValue, ",")
	loopCounts := make([]int, len(countStrings))
	for i, countStr := range countStrings {
		count, err := strconv.Atoi(countStr)
		if err != nil || count <= 0 {
			return defaultCounts
		}
		loopCounts[i] = count
	}
	return loopCounts
}

func getDepthPrefix(depth int) string {
	if depth >= 0 && depth < 26 {
		return string(rune('a' + depth))
	}
	return "?"
}

func generateLoops(loopCounts []int, depth int, currentIndices []string) {
	prefix := getDepthPrefix(depth)

	for i := 1; i <= loopCounts[depth]; i++ {
		formattedIndex := fmt.Sprintf("%s=%d", prefix, i)
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
