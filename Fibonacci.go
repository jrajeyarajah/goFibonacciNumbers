package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(stop int) uint64 {
	var a, b uint64 = 0, 1
	for i := 0; i < stop; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	argsWithoutProg := os.Args[1:]

	/*/argsWithoutProg = /*/

	fmt.Println(argsWithoutProg[0])


	tmp, err := strconv.Atoi(argsWithoutProg[0])
	
	if err == nil {
		if tmp < 94 {
			fmt.Println(fib(tmp))
		} else {
			fmt.Println("Sorry can only go up to 93rd number.")
		}
	}
}
