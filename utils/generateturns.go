package utils

import "lem-in/models"

// GenerateTurns calculates the maximum number of turns any path can have in the given options.
func GenerateTurns(option map[int][]int, paths []models.Path) int {
	maxTurns := 0

	// Iterate through each path and calculate the number of turns
	for i, path := range paths {
		// The number of rooms excluding the start room
		rooms := len(path.Rooms) - 1
		// The number of ants for the current path
		ants := len(option[i])
		// Calculate turns: rooms + ants - 1
		turns := rooms + ants - 1

		// Update maxTurns if the current path's turns exceed the previous max
		if turns > maxTurns {
			maxTurns = turns
		}

		// Early exit if the maximum possible number of turns is found
		// This would depend on your domain constraints; here it's just a hypothetical example
		if maxTurns == 1000 { // Assume 1000 is a very high limit
			break
		}
	}

	return maxTurns
}
