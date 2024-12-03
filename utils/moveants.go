package utils

import (
	"fmt"
	"lem-in/models"
)

// returns a slice of moves that indicates paths takes by each ant
func MoveAnts(paths []models.Path, antsperroom map[int][]int, turns int) []string {
	moves := make([]string, turns)
	for i, path := range paths {
		ants := antsperroom[i]
		for j, ant := range ants {
			for k, room := range path.Rooms[1:] {
				moves[j+k] += fmt.Sprintf("L%v-%v ", ant, room)
			}
		}
	}
	return moves
}
