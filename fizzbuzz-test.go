package main

import (
	"fmt"
	"strconv"
)

// The fizzBuzz function print numbers from 'startNum' to 'endNum', but :
//   - It will print "Fizz" if the number is multiples of fizzNum,
//   - It will print "Buzz" if the number is multiples of buzzNum,
//   - It will print "FizzBuzz" if the number are multiples of both fizzNum and buzzNum.
//
// Parameters:
//  1. startNum (integer) : the starting number of the range (inclusive).
//  2. endNum (integer)   : the ending number of the range (inclusive),
//  3. fizzNum (integer)  : the number replaced with "Fizz", (fizzNum > 0).
//  4. buzzNum (integer)  : the number replaced with "Buzz", (buzzNum > 0)
func fizzBuzz(startNum, endNum, fizzNum, buzzNum int) {
	if startNum > endNum {
		fmt.Println("Error: 'startNum' should be less than or equal to 'endNum'")
		return
	}

	if fizzNum <= 0 || buzzNum <= 0 {
		fmt.Println("Error: 'fizzNum' and 'buzzNum' number should be greather than 0 (zero)")
		return
	}

	for i := startNum; i <= endNum; i++ {
		result := ""

		if i%fizzNum == 0 {
			result += "Fizz"
		}

		if i%buzzNum == 0 {
			result += "Buzz"
		}

		if result == "" {
			result = strconv.Itoa(i)
		}

		fmt.Println(result)
	}
}

func main() {
	startNum := 1
	endNum := 100
	fizzNum := 3
	buzzNum := 5

	fizzBuzz(startNum, endNum, fizzNum, buzzNum)
}
