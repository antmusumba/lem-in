package utils

import (
	"fmt"
	"strings"
	"lem-in/models"
)

// MoveAnts generates a slice of moves indicating the paths taken by each ant.
// Each move is represented as a string "L<ant>-<room>".
func MoveAnts(paths []models.Path, antsPerRoom map[int][]int, totalTurns int) []string {
	moves := make([]string, totalTurns)

	for pathIndex, path := range paths {
		ants := antsPerRoom[pathIndex] // Ants assigned to this path
		for antIndex, ant := range ants {
			for turnOffset, room := range path.Rooms[1:] {
				moveIndex := antIndex + turnOffset
				if moveIndex >= totalTurns {
					break // Avoid out-of-bounds issues
				}
				var builder strings.Builder
				if moves[moveIndex] != "" {
					builder.WriteString(moves[moveIndex])
				}
				builder.WriteString(fmt.Sprintf("L%d-%s ", ant, room))
				moves[moveIndex] = builder.String()
			}
		}
	}

	// Trim trailing spaces from each move string
	for i, move := range moves {
		moves[i] = strings.TrimSpace(move)
	}

	return moves
}
