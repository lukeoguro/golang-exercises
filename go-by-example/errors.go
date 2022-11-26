package main

import (
	"errors"
	"fmt"
)

func f1(n int) (int, error) {
	if n == 42 {
		return -1, errors.New("42 sucks")
	}
	return n, nil
}

func main() {
	if n, e := f1(10); e != nil {
		fmt.Println("f1 failed:", e)
	} else {
		fmt.Println("f1 worked:", n)
	}

	if n, e := f1(42); e != nil {
		fmt.Println("f1 failed:", e)
	} else {
		fmt.Println("f1 worked:", n)
	}
}
