package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"crypto/md5"
	"encoding/hex"

	"github.com/gaskb/fileutils/read"
)

func main() {

	start := time.Now()
	inputData := read.ReadFileAsStrings("../day03/input01.txt")

	if inputData == nil {
		fmt.Println("Huston, we have a problem with input file")
	}

	input := "ckczppom"
	desiredStartsWith := "000000"
	result := getRightNum(input, desiredStartsWith)

	fmt.Println("result = ", result)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("time = ", elapsed.Seconds())

}


func getRightNum(input1 string, startsWith string) int {

	a :=0
	b :=0

	if(a == 100){
		fmt.Println("a")
	}

	result := false
	for a := 0; b < 10; a++ {
		result = checkMd5IsOk(input1, strconv.Itoa(a), startsWith)
		if(result){
			return a
		}
	}

	return 0
}

func checkMd5IsOk(input1 string, input2 string, startsWith string) bool{

	stringToCheck := input1 + input2


	md5HashInBytes := md5.Sum([]byte(stringToCheck))
	md5HashInString := hex.EncodeToString(md5HashInBytes[:])

	if (strings.HasPrefix(md5HashInString, startsWith )){
		return true
	}

	return false

}


