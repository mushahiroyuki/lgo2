package main

import (
	"embed"
	"fmt"
)

//go:embed parent_dir
var noHidden embed.FS

//go:embed parent_dir/*
var parentHiddenOnly embed.FS

//go:embed all:parent_dir
var allHidden embed.FS

func main() {
	checkForHidden("noHidden", noHidden)
	checkForHidden("parentHiddenOnly", parentHiddenOnly)
	checkForHidden("allHidden", allHidden)
}

func checkForHidden(name string, dir embed.FS) {
	fmt.Println(name)
	allFileNames := []string{"parent_dir/.hidden", "parent_dir/child_dir/.hidden"}
	for _, v := range allFileNames {
		_, err := dir.Open(v)
		if err == nil {
			fmt.Println(v, "found")
		}
	}
	fmt.Println()
}
