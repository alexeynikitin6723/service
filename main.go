package main

import (
	"service/solutions"
)

func main() {
	err := solutions.MakeHandler()
	if err != nil {
		panic(err)
	}
}
