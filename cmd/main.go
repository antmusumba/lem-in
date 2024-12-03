package main

import (
	"fmt"
	"os"

	"lem-in/models"
	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go file.txt")
		return
	}
	filename := os.Args[1]
	// Parse the file
	colony, err := utils.ParseFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}

	// Find paths and determine moves
	paths, antsPerPath, turns := utils.FindPaths(colony)
	moves := utils.MoveAnts(paths, antsPerPath, turns)

	// Print the file contents
	fmt.Println(models.FileContents)

	// Print the moves
	for _, move := range moves {
		fmt.Println(move)
	}
}
