package utils

import "lem-in/models"

// finding maximum turns the paths to be optimized can have
func GenerateTurns(option map[int][]int, paths []models.Path) int {
	maxturns := 0
	for i, path := range paths {
		rooms := len(path.Rooms) - 1
		ants := len(option[i])
		turns := rooms + ants - 1
		if turns > maxturns {
			maxturns = turns
		}
	}
	return maxturns
}
