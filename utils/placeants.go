package utils

import "lem-in/models"

// PlaceAnts assigns ants to paths in the colony and returns a map of path indices to the ants assigned to them.
func PlaceAnts(colony *models.AntColony, paths []models.Path) map[int][]int {
	totalAnts := colony.NumberOfAnts
	pathAssignments := make(map[int][]int)

	for ant := 1; ant <= totalAnts; ant++ {
		placeAnt(ant, paths, pathAssignments)
	}

	return pathAssignments
}

// placeAnt attempts to place an ant on a path recursively, ensuring an optimal distribution of ants.
func placeAnt(ant int, paths []models.Path, pathAssignments map[int][]int) bool {
	return placeAntHelper(ant, paths, pathAssignments, 0)
}

// placeAntHelper is a recursive helper function for placing ants optimally across paths.
func placeAntHelper(ant int, paths []models.Path, pathAssignments map[int][]int, currentPath int) bool {
	// Base case: Assign to the last path if it's the only choice
	if currentPath == len(paths)-1 {
		pathAssignments[currentPath] = append(pathAssignments[currentPath], ant)
		return true
	}

	// Calculate distribution cost (rooms + ants already assigned) for the current and next paths
	currentPathLoad := len(paths[currentPath].Rooms) - 2 + len(pathAssignments[currentPath])
	nextPathLoad := len(paths[currentPath+1].Rooms) - 2 + len(pathAssignments[currentPath+1])

	// Assign to the less costly path
	if currentPathLoad > nextPathLoad {
		return placeAntHelper(ant, paths, pathAssignments, currentPath+1)
	}

	// Assign ant to the current path
	pathAssignments[currentPath] = append(pathAssignments[currentPath], ant)
	return true
}
