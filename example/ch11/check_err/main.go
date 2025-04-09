package main

import (
	"errors"
	"fmt"
)

func main() {
	err := returnErr(false)
	if err != nil {
		fmt.Println(err)
	}
	err = returnErr(true)
	fmt.Println("end of program")
}

func returnErr(b bool) error {
	if b {
		return errors.New("err")
	}
	return nil
}
