package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed passwords.txt
var passwords string

func main() {
	pwds := strings.Split(passwords, "\n")
	if len(os.Args) > 1 {
		for _, v := range pwds {
			if v == os.Args[1] {
				fmt.Println("true")
				os.Exit(0)
			}
		}
		fmt.Println("false")
	}
}
