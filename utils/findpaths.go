package utils

import (
	"lem-in/resources"
)

// FindPaths finds all possible paths from start to end using BFS.
func FindPaths(colony *resources.AntColony) ([]resources.Path, map[int][]int, int) {
	paths := []resources.Path{}
	queue := [][]string{}
	queue = append(queue, []string{colony.Start})

	// BFS loop
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]

		// Current room is the last room in the path
		currentRoom := currentPath[len(currentPath)-1]

		// If we've reached the end, add the path to allPaths
		if currentRoom == colony.End {
			paths = append(paths, resources.Path{RoomsInThePath: currentPath})
			continue
		}

		// Explore adjacent rooms
		for _, nextRoom := range colony.Links[currentRoom] {
			if !containsRoom(currentPath, nextRoom) {
				newPath := append([]string(nil), currentPath...)
				newPath = append(newPath, nextRoom)
				queue = append(queue, newPath)
			}
		}
	}

	return ChooseOptimumPath(paths, colony)
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
func ChooseOptimumPath(paths []resources.Path, colony *resources.AntColony) ([]resources.Path, map[int][]int, int) {
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
func OptimizedPaths1(paths []resources.Path) []resources.Path {
	optimized := []resources.Path{paths[0]}
	for i := 1; i < len(paths); i++ {
		if Check(paths[i].RoomsInThePath, optimized) {
			optimized = append(optimized, paths[i])
		}
	}
	return optimized
}

// Check verifies if a path shares any rooms with already optimized paths (excluding start/end).
func Check(path []string, optimized []resources.Path) bool {
	// Use a map for faster lookups
	visitedRooms := make(map[string]struct{})
	for _, optpath := range optimized {
		for _, room := range optpath.RoomsInThePath[1 : len(optpath.RoomsInThePath)-1] { // Ignore start and end rooms
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
func OptimizedPaths2(paths []resources.Path, colony *resources.AntColony) []resources.Path {
	half := colony.NumberOfAnts / 2
	optimized := []resources.Path{paths[0]}

	for i := 1; i < len(paths); i++ {
		if len(paths[i].RoomsInThePath)-1 <= half { // Exclude start and end rooms
			unique, index := Check2(paths[i].RoomsInThePath, optimized)
			if !unique {
				// Replace path if lengths differ
				if len(optimized[index].RoomsInThePath) != len(paths[i].RoomsInThePath) {
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
func Check2(path []string, optimized []resources.Path) (bool, int) {
	for i, optpath := range optimized {
		for _, room := range path[1 : len(path)-1] { // Exclude start and end rooms
			for _, optRoom := range optpath.RoomsInThePath[1 : len(optpath.RoomsInThePath)-1] {
				if room == optRoom {
					return false, i
				}
			}
		}
	}
	return true, 0
}

// Remove removes a path from the optimized slice by index.
func Remove(optimized []resources.Path, index int) []resources.Path {
	return append(optimized[:index], optimized[index+1:]...)
}
