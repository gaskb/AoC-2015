package main

import (
	"fmt"
	"time"
	"strings"

	"github.com/gaskb/fileutils/read"
)

func main() {

	start := time.Now()
	inputData := read.ReadFileAsStrings("input01.txt")

	if inputData == nil {
		fmt.Println("Huston, we have a problem with input file")
	}

	//inputData = []string{"ieodomkazucvgmuy"}

	result := checkInputStrings2(inputData)

	fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())

}



func checkInputStrings(inputData []string) int {

	counter := 0
	result := 0

	for counter < len(inputData) {
		
		actualString:= inputData[counter]


		if (howManyVowelsInString(actualString)> 2){
			if(stringContainsDouble(actualString)){
				if(!stringContainsForbiddenString(actualString)){
					result = result + 1
					// fmt.Println("result = ", result)
					fmt.Println("actualString = ", actualString)
					// fmt.Println("")
				}
			}
		}


		counter++
	}

	return result;
}


func checkInputStrings2(inputData []string) int {
	counter := 0
	result := 0

	for counter < len(inputData) {
		
		actualString:= inputData[counter]


		if (stringContainsDoublePairChar(actualString)){
			if(stringContainsSameCharTwoTimesWithAnotherCharBetween(actualString)){
				result = result + 1
			}
		}


		counter++
	}

	return result;
}


func stringContainsDoublePairChar(actualString string) bool {

	counter:=1
	for counter < len(actualString) {

		couple := string(actualString[counter -1]) + string(actualString[counter])

		counter2 := counter + 2
		for counter2 < len(actualString) {
			couple2 := string(actualString[counter2 -1]) + string(actualString[counter2])
			if couple == couple2 {
				return true
			}

			counter2++
		}
		counter ++
	}

	return false

}

func stringContainsSameCharTwoTimesWithAnotherCharBetween(actualString string) bool{

	counter:=2
	for counter < len(actualString) {

		actualChar := string(actualString[counter])
		tempChar := string(actualString[counter-2])

		if(actualChar == tempChar){
			return true
		}
		counter ++
	}

	return false
}



func stringContainsForbiddenString(actualString string) bool{

	counter := 0
	
	forbiddenStrings := []string{"ab","cd","pq","xy"}


	for counter < len(forbiddenStrings) {

		currentCheck := string(forbiddenStrings[counter])

		if (strings.Contains(actualString, currentCheck)){
			// fmt.Println("actualString ", actualString, " contains  ", currentCheck )

			// fmt.Println("")
			// fmt.Println("")
			return true
		}

		counter++
	}

	// fmt.Println("No forbidden string !")

	return false
	
}

func stringContainsDouble(actualString string) bool{

	counter := 0
	lastChar := ""
	for counter < len(actualString) {

		currentChar := string(actualString[counter])

		if currentChar == lastChar{
			fmt.Println("contains double ", lastChar )
			return true
		}

		lastChar = currentChar

		counter++
	}

	// fmt.Println("")
	// fmt.Println("")
	return false

}


func howManyVowelsInString(inputString string) int {

	counter := 0
	result := 0
	vowels := [] string {"a","e","i","o","u"}

	for counter < len(inputString) {

		currentChar := string(inputString[counter])

		//foundVowel := sort.SearchStrings(vowels, currentChar) 
		foundVowel := stringSliceContainsString(vowels, currentChar) 
		if (foundVowel){
			fmt.Println("foundVowel :", foundVowel, " -> ", currentChar )
			result = result + 1
		}
		counter++
	}

	//fmt.Println("result for ", inputString , " = ", result)
	return result

}




func stringSliceContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}





