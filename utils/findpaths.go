package utils

import (
	"lem-in/models"
)

// FindPaths finds all possible paths from start to end using BFS.
func FindPaths(colony *models.AntColony) ([]models.Path, map[int][]int, int) {
	var allPaths []models.Path
	var queue [][]string
	queue = append(queue, []string{colony.Start})

	// Reset IsVisited flag for all rooms
	for i := range colony.Rooms {
		colony.Rooms[i].IsVisited = false
	}

	// BFS loop
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]

		// Current room is the last room in the path
		currentRoom := currentPath[len(currentPath)-1]

		// If we've reached the end, add the path to allPaths
		if currentRoom == colony.End {
			allPaths = append(allPaths, models.Path{Rooms: currentPath})
			continue
		}

		// Explore adjacent rooms
		for _, nextRoom := range colony.Links[currentRoom] {
			if !containsRoom(currentPath, nextRoom) {
				// Create a new path by copying the current path and appending the next room
				newPath := append([]string(nil), currentPath...) // Efficient copy
				newPath = append(newPath, nextRoom)
				queue = append(queue, newPath)
			}
		}
	}

	return ChooseOptimumPath(allPaths, colony)
}

// Helper function to check if a room is in the path
func containsRoom(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

// ChooseOptimumPath selects the optimum paths based on the number of turns.
func ChooseOptimumPath(paths []models.Path, colony *models.AntColony) ([]models.Path, map[int][]int, int) {
	shortest1 := OptimizedPaths1(paths)
	shortest2 := OptimizedPaths2(paths, colony)
	firstop := PlaceAnts(colony, shortest1)
	secondop := PlaceAnts(colony, shortest2)
	turns1 := GenerateTurns(firstop, shortest1)
	turns2 := GenerateTurns(secondop, shortest2)

	// Choose the path with fewer turns
	if turns1 <= turns2 {
		return shortest1, firstop, turns1
	}
	return shortest2, secondop, turns2
}

// OptimizedPaths1 filters paths that don't share rooms.
func OptimizedPaths1(paths []models.Path) []models.Path {
	optimized := []models.Path{paths[0]}
	for i := 1; i < len(paths); i++ {
		if Check(paths[i].Rooms, optimized) {
			optimized = append(optimized, paths[i])
		}
	}
	return optimized
}

// Check verifies if a path shares any rooms with already optimized paths (excluding start/end).
func Check(path []string, optimized []models.Path) bool {
	// Use a map for faster lookups
	visitedRooms := make(map[string]struct{})
	for _, optpath := range optimized {
		for _, room := range optpath.Rooms[1 : len(optpath.Rooms)-1] { // Ignore start and end rooms
			visitedRooms[room] = struct{}{}
		}
	}

	for _, room := range path[1 : len(path)-1] { // Ignore start and end rooms
		if _, found := visitedRooms[room]; found {
			return false
		}
	}
	return true
}

// OptimizedPaths2 filters paths based on colony's ant count and unique room usage.
func OptimizedPaths2(paths []models.Path, colony *models.AntColony) []models.Path {
	half := colony.NumberOfAnts / 2
	optimized := []models.Path{paths[0]}

	for i := 1; i < len(paths); i++ {
		if len(paths[i].Rooms)-1 <= half { // Exclude start and end rooms
			unique, index := Check2(paths[i].Rooms, optimized)
			if !unique {
				// Replace path if lengths differ
				if len(optimized[index].Rooms) != len(paths[i].Rooms) {
					optimized[index] = paths[i] // Efficient path replacement
				}
			} else {
				optimized = append(optimized, paths[i])
			}
		}
	}
	return optimized
}

// Check2 checks if a path is unique and does not share rooms with existing paths.
func Check2(path []string, optimized []models.Path) (bool, int) {
	for i, optpath := range optimized {
		for _, room := range path[1 : len(path)-1] { // Exclude start and end rooms
			for _, optRoom := range optpath.Rooms[1 : len(optpath.Rooms)-1] {
				if room == optRoom {
					return false, i
				}
			}
		}
	}
	return true, 0
}

// Remove removes a path from the optimized slice by index.
func Remove(optimized []models.Path, index int) []models.Path {
	return append(optimized[:index], optimized[index+1:]...)
}
