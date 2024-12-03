package utils

import (
	"lem-in/models"
)

// FindPaths finds all possible paths from start to end using BFS
func FindPaths(colony *models.AntColony) ([]models.Path, map[int][]int, int) {
	var allPaths []models.Path
	var queue [][]string
	queue = append(queue, []string{colony.Start})
	for i := range colony.Rooms {
		colony.Rooms[i].IsVisited = false
	}
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		currentRoom := currentPath[len(currentPath)-1]
		if currentRoom == colony.End {
			allPaths = append(allPaths, models.Path{Rooms: currentPath})
			continue
		}
		for _, nextRoom := range colony.Links[currentRoom] {
			if !containsRoom(currentPath, nextRoom) {
				// Create a new path with the next room added
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				newPath = append(newPath, nextRoom)
				queue = append(queue, newPath)
			}
		}
	}
	return ChooseOptimumPath(allPaths, colony)
}

// Helper function to check if a room is in a path
func containsRoom(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func ChooseOptimumPath(paths []models.Path, Antcolony *models.AntColony) ([]models.Path, map[int][]int, int) {
	shortest1 := OptimizedPaths1(paths)
	shortest2 := OptimizedPaths2(paths, Antcolony)
	firstop := PlaceAnts(Antcolony, shortest1)
	secondop := PlaceAnts(Antcolony, shortest2)
	turns1 := GenerateTurns(firstop, shortest1)
	turns2 := GenerateTurns(secondop, shortest2)
	finalpath := shortest1
	finalAntspalced := firstop
	turns := turns1
	if turns1 > turns2 {
		finalpath = shortest2
		finalAntspalced = secondop
		turns = turns2
	}
	return finalpath, finalAntspalced, turns
}

// returns a slice of paths that don't share rooms 
func OptimizedPaths1(paths []models.Path) []models.Path {
	optimized := make([]models.Path, 0)
	optimized = append(optimized, paths[0])
	for i := 1; i < len(paths); i++ {
		if Check(paths[i].Rooms, optimized) {
			optimized = append(optimized, paths[i])
		}
	}
	return optimized
}

// checks if a room in a path has already been utilized by other paths already stored in optimized paths
func Check(path []string, optimized []models.Path) bool {
	for _, optpath := range optimized {
		for k := 1; k < len(path)-1; k++ {
			for j := 1; j < len(optpath.Rooms)-1; j++ {
				if path[k] == optpath.Rooms[j] {
					return false
				}
			}
		}
	}
	return true
}

// finds an optimized path set with unique paths apart from start and end 
func OptimizedPaths2(paths []models.Path, colony *models.AntColony) []models.Path {
	half := colony.NumberOfAnts / 2
	optimized := make([]models.Path, 0)
	optimized = append(optimized, paths[0])
	for i := 1; i < len(paths); i++ {
		if len(paths[i].Rooms)-1 <= half {
			unique, index := Check2(paths[i].Rooms, optimized)
			if !unique {
				if len(optimized[index].Rooms) != len(paths[i].Rooms) {
					optimized = Remove(optimized, index)
					optimized = append(optimized, paths[i])
				}
			} else {
				optimized = append(optimized, paths[i])
			}
		}
	}
	return optimized
}

// checks if a path is unique and doesnt have common rooms with others
func Check2(path []string, optimized []models.Path) (bool, int) {
	for i, optpath := range optimized {
		for k := 1; k < len(path)-1; k++ {
			for j := 1; j < len(optpath.Rooms)-1; j++ {
				if path[k] == optpath.Rooms[j] {
					return false, i
				}
			}
		}
	}
	return true, 0
}

// removes a non unique path from the optimized paths
func Remove(optimized []models.Path, index int) []models.Path {
	new := make([]models.Path, 0)
	for i, path := range optimized {
		if i != index {
			new = append(new, path)
		}
	}
	return new
}
