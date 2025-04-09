package main  //liststart1

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed help
var helpInfo embed.FS

func main() {
	if len(os.Args) == 1 {
		printHelpFiles()
		os.Exit(0)
	}
	data, err := helpInfo.ReadFile("help/" + os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}  //listend1

func printHelpFiles() {  //liststart2
	fmt.Println("contents:")
	fs.WalkDir(helpInfo, "help",
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				_, fileName, _ := strings.Cut(path, "/")
				fmt.Println(fileName)
			}
			return nil
		})
}  //listend2
