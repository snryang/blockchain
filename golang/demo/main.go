package main

import (
	"fmt"
)

func main() {
	ywb := f(10)
	fmt.Println(ywb())
	fmt.Println(ywb())

}

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}
