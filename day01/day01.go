package main

import (
	"fmt"
	"time"

	"github.com/gaskb/fileutils/read"
)

func main() {

	start := time.Now()
	inputData := read.ReadFileAsChars("input01.txt")

	//fmt.Println("inputData = ", inputData)

	counter := 0
	result := 0
	for counter < len(inputData) {
		thisChar := inputData[counter]
		//fmt.Println("parsing = ", thisChar)

		if thisChar == "(" {
			//fmt.Println("found ( ")

			result++
		}
		if thisChar == ")" {
			//fmt.Println("found ) ")

			result--
		}

		// if result == -1 {
		// 	fmt.Println("counter = ", counter)

		// 	break
		// }

		counter++
	}

	fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())
}
