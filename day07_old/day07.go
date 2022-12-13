package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/gaskb/fileutils/read"
)

func main() {

	start := time.Now()
	inputData := read.ReadFileAsStrings("input01.txt")

	if inputData == nil {
		fmt.Println("Huston, we have a problem with input file")
	}

	//inputData = []string{"turn off 0,0 through 999,999"}

	counter := 0

	for counter < len(inputData) {
		rawInstruction := inputData[counter]
		instruction := readInstruction(rawInstruction)

		fmt.Println("rawInstruction = ", rawInstruction)
		fmt.Println("instruction = ", instruction)

		fmt.Println("")

		counter ++

	}


	//fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())

}

type Instruction struct {
    verb string
    source1 string
	source2 string
	sourceInt1 int
	sourceInt2 int
	destination string
}


/**
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
*/



func executeInstruction(instruction Instruction, resultSlice []int) []int{

	if instruction.verb == "STORE" {
		resultSlice[instruction.destination] = instruction.sourceInt1
	}

	if instruction.verb == "NOT" {
		resultSlice[instruction.destination] = instruction.sourceInt1
	}



}

func readInstruction(rawInstruction string) Instruction {

	temp := strings.Split(rawInstruction, " -> ")

	//finding verb

	instruction := Instruction{}

	instruction.destination = temp[1]
	
	temp = strings.Split(temp[0], " ")

	//storage case
	if(len(temp) == 1){
		instruction.verb = "STORE"

		a, err := strconv.Atoi(temp[0])
		if(err != nil) {
			fmt.Println("ERROR")
			instruction.source1 = temp[0] 
		} else {
			instruction.sourceInt1 = a
		}
	}

	//not case
	if(len(temp) == 2){
		instruction.verb = "NOT"

		a, err := strconv.Atoi(temp[1])
		if(err != nil) {
			fmt.Println("ERROR")
			instruction.source1 = temp[1] 
		} else {
			instruction.sourceInt1 = a
		}	
	}

	if (len(temp) == 3){
		instruction.verb = temp[1]

		a, erra := strconv.Atoi(temp[1])
		if(erra != nil) {
			fmt.Println("ERROR")
			instruction.source1 = temp[0] 
		} else {
			instruction.sourceInt1 = a
		}

		b, errb := strconv.Atoi(temp[2])
		if(errb != nil) {
			fmt.Println("ERROR")
			instruction.source2 = temp[2] 
		} else {
			instruction.sourceInt2 = b
		}
	}
	
	return instruction

}


