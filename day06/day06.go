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

	lightsMap := initGrid()

	counter := 0

	for counter < len(inputData) {
		rawInstruction := inputData[counter]
		instruction := readInstruction(rawInstruction)

		lightsMap = followInstruction(instruction, lightsMap)

		counter ++

	}

	result := countBrightness(lightsMap)


	fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())

}

type Instruction struct {
    verb string
    start []string
	end []string
}

func countLights(lightsMap [1000][1000]int ) int {

	result := 0
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			if(lightsMap[a][b] != 0){
				result = result +1
			}	
		}
	}

	return result
	
}

func countBrightness(lightsMap [1000][1000]int ) int {

	result := 0
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
				result = result + lightsMap[a][b]
		}
	}

	return result
	
}



func initGrid() [1000][1000]int{

	const xMax = 1000
	const yMax = 1000
	
	lightsMap := [xMax][yMax]int{}

	for a := 0; a < xMax; a++ {
		for b := 0; b < yMax; b++ {
			lightsMap[a][b] = 0
		}
	}

	return lightsMap

}



/**
turn on 887,9 through 959,629
turn off 539,243 through 559,965
toggle 720,196 through 897,994
*/

func readInstruction(rawInstruction string) Instruction {

	temp := strings.Split(rawInstruction, " through ")

	//finding verb
	
	temp[0] = strings.Replace(temp[0], "turn ", "", 1)

	verb := strings.Split(temp[0], " ")[0]
	start := strings.Split(strings.Split(temp[0], " ")[1], ",")
	end := strings.Split(temp[1],",")

	result := Instruction{
		verb:verb,
		start:start,
		end:end,
	}  
	
	return result

}


func followInstruction(instruction Instruction, lightsMap [1000][1000]int) [1000][1000]int {

	startx, _:= strconv.Atoi(instruction.start[0])
	starty, _:= strconv.Atoi(instruction.start[1])

	endx, _:= strconv.Atoi(instruction.end[0])
	endy, _:= strconv.Atoi(instruction.end[1])

	lightsMap = executev2(instruction.verb, startx, starty, endx, endy, lightsMap)
	
	return lightsMap


}



func execute (verb string, startx int, starty int, endx int, endy int, lightsMap [1000][1000]int ) [1000][1000]int {

	if startx > endx{
		startx, endx = endx, startx
	}
	if starty > endy{
		starty, endy = endy, starty
	}

	for a := startx; a <= endx; a++ {
		for b := starty; b <= endy; b++ {
			status := lightsMap[a][b]
			if(verb == "on"){
				lightsMap[a][b] = 1
			}
			if(verb == "off"){
				lightsMap[a][b] = 0
			}
			if(verb == "toggle"){
				lightsMap[a][b] = 0
				if(status == 0){
					lightsMap[a][b] = 1
				}
			}
		}
	}

	return lightsMap

}

func executev2 (verb string, startx int, starty int, endx int, endy int, lightsMap [1000][1000]int ) [1000][1000]int {

	if startx > endx{
		startx, endx = endx, startx
	}
	if starty > endy{
		starty, endy = endy, starty
	}

	for a := startx; a <= endx; a++ {
		for b := starty; b <= endy; b++ {
			status := lightsMap[a][b]
			if(verb == "on"){
				lightsMap[a][b] = status + 1
			}
			if(verb == "off"){
				lightsMap[a][b] = status - 1
			}
			if(verb == "toggle"){
				lightsMap[a][b] = status + 2
			}

			if (lightsMap[a][b]) < 0 {
				lightsMap[a][b] = 0
			}
		}
	}

	return lightsMap

}