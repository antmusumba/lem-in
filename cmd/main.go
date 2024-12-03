package main

import (
	"fmt"

	"lem-in/models"
	"lem-in/utils"
)

func main() {
	filename, errmsg := utils.ParseArgs()
	if errmsg != "" {
		fmt.Println(errmsg)
		return
	}
	Antcolony, err := utils.ParseFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}
	paths, antsperpath, turns := utils.FindPaths(Antcolony)
	moves := utils.MoveAnts(paths, antsperpath, turns)

	// Print the file contents
	fmt.Println(models.FileContents)

	for _, move := range moves {
		fmt.Println(move)
	}
}
