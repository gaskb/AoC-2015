package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gaskb/fileutils/read"
)

func main() {

	start := time.Now()
	inputData := read.ReadFileAsStrings("input01.txt")

	if inputData != nil {
		fmt.Println("result = ")

	}
	counter := 0
	result := 0

	for counter < len(inputData) {
		thisGiftSlice := strings.Split(inputData[counter], "x")
		orderedSlice := orderDim(thisGiftSlice)
		// neededPaper := getNeededPaper(orderedSlice)
		neededRibbon := getNeededRibbon(orderedSlice)

		result = result + neededRibbon
		counter++
	}

	// thisGiftSlice := strings.Split("4x3x2", "x")
	// result = getNeededPaper(thisGiftSlice)

	fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())

}

func orderDim(startingSlice []string) []int {
	counter := 0
	dimSlice := []int{}
	for counter < len(startingSlice) {
		dim, _ := strconv.Atoi(startingSlice[counter])
		dimSlice = append(dimSlice, dim)
		counter++
	}

	sort.Ints(dimSlice)

	return dimSlice
}

func getNeededPaper(dimSlice []int) int {

	result := dimSlice[0]*dimSlice[1]*3 + dimSlice[0]*dimSlice[2]*2 + dimSlice[1]*dimSlice[2]*2
	fmt.Println("getNeededPaper  = ", result)

	return result
}

func getNeededRibbon(dimSlice []int) int {
	result := (dimSlice[0]+dimSlice[1])*2 + dimSlice[0]*dimSlice[1]*dimSlice[2]
	fmt.Println("getNeededRibbon  = ", result)

	return result
}
