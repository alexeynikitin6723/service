package main

import (
	"fmt"
	"service/solutions"
)

func main() {
	fmt.Println(solutions.GetResult(0))
	err := solutions.MakeHandler()
	if err != nil {
		panic(err)
	}
}
