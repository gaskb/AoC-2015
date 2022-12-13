package main

import (
	"fmt"
	"time"

	"github.com/gaskb/fileutils/read"
)

func main() {

	start := time.Now()
	inputData := read.ReadFileAsStrings("input01.txt")

	if inputData != nil {
		fmt.Println("Huston, we have a problem with input file")

	}

	//test := "^v^v^v^v^v"

	result := goSanta(inputData[0])

	fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())

}




func goSanta(directions string) int {

	result := 0
	counter := 0

	const xMax = 1000
	const yMax = 1000
	santaMap := [xMax][yMax]int{}

	for a := 0; a < xMax; a++ {
		for b := 0; b < yMax; b++ {
			santaMap[a][b] = 0
		}
	}
	
	x := 500
	y := 500

	w := 500
	z := 500

	santaMap[x][y]=1

	
	for counter < len(directions) {

		visits :=0 
		direction := directions[counter]

		if ( counter % 2 == 0 ) {
			if(string(direction) == ">"){
				x=x+1
			}
			if(string(direction) == "<"){
				x=x-1
			}
			if(string(direction) == "v"){
				y=y-1
			}
			if(string(direction) == "^"){
				y=y+1
			}

			fmt.Println("x = ", x)
			fmt.Println("y = ", y)

			visits = santaMap[x][y]

			santaMap[x][y] = visits +1

		}else{
			if(string(direction) == ">"){
				w=w+1
			}
			if(string(direction) == "<"){
				w=w-1
			}
			if(string(direction) == "v"){
				z=z-1
			}
			if(string(direction) == "^"){
				z=z+1
			}

			fmt.Println("w = ", w)
			fmt.Println("z = ", z)

			visits = santaMap[w][z]

			santaMap[w][z] = visits +1
		}

		fmt.Println("direction = ", direction)		

		fmt.Println("")
		fmt.Println("")

		counter++
	}


	for a := 0; a < xMax; a++ {
		for b := 0; b < yMax; b++ {
			if(santaMap[a][b] != 0){
				result = result +1
			}	
		}
	}


	return result;


}
