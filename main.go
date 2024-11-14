package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Room represents a room in the ant farm
type Room struct {
	Name  string
	X     int
	Y     int
	Links []string // Connected rooms
}

// Graph represents the entire ant farm
type Graph struct {
	Rooms map[string]*Room
	Start string
	End   string
}

// ReadInput reads the input file and builds the graph
func ReadInput(filename string) (int, *Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numberOfAnts int
	graph := &Graph{Rooms: make(map[string]*Room)}

	// Read number of ants
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &numberOfAnts)
	}

	var currentSection string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			continue // Ignore comments
		}

		switch line {
		case "##start":
			currentSection = "start"
			continue
		case "##end":
			currentSection = "end"
			continue
		default:
			if currentSection == "start" {
				var x, y int
				room := &Room{}
				if _, err := fmt.Sscanf(line, "%s %d %d", &room.Name, &x, &y); err != nil || room.Name == "" {
					return 0, nil, fmt.Errorf("invalid room definition at start: %s", line)
				}
				room.X = x
				room.Y = y
				graph.Rooms[room.Name] = room
				graph.Start = room.Name
			} else if currentSection == "end" {
				var x, y int
				room := &Room{}
				if _, err := fmt.Sscanf(line, "%s %d %d", &room.Name, &x, &y); err != nil || room.Name == "" {
					return 0, nil, fmt.Errorf("invalid room definition at end: %s", line)
				}
				room.X = x
				room.Y = y
				graph.Rooms[room.Name] = room
				graph.End = room.Name
			} else {
				// This is a tunnel definition or another room definition.
				if strings.Contains(line, "-") {
					parts := strings.Split(line, "-")
					if len(parts) == 2 {
						room1 := parts[0]
						room2 := parts[1]
						if _, exists := graph.Rooms[room1]; exists {
							graph.Rooms[room1].Links = append(graph.Rooms[room1].Links, room2)
						} else {
							fmt.Printf("WARNING: Room %s does not exist when linking.\n", room1)
						}
						if _, exists := graph.Rooms[room2]; exists {
							graph.Rooms[room2].Links = append(graph.Rooms[room2].Links, room1)
						} else {
							fmt.Printf("WARNING: Room %s does not exist when linking.\n", room2)
						}
					}
				} else {
					var x, y int
					room := &Room{}
					if _, err := fmt.Sscanf(line, "%s %d %d", &room.Name, &x, &y); err != nil || room.Name == "" {
						return 0, nil, fmt.Errorf("invalid room definition: %s", line)
					}
					room.X = x
					room.Y = y
					graph.Rooms[room.Name] = room
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}

	return numberOfAnts, graph, nil
}

// BFS finds the shortest path from start to end using BFS algorithm
func (g *Graph) BFS() []string {
	if g == nil || g.Rooms == nil {
		fmt.Println("ERROR: Graph or Rooms map is not initialized.")
		return nil
	}

	queue := []string{g.Start}
	cameFrom := make(map[string]string)
	cameFrom[g.Start] = ""

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

        // Check if the current room exists in the graph.
        room, exists := g.Rooms[current]
        if !exists {
            fmt.Printf("ERROR: Room %s does not exist.\n", current)
            continue // Skip this iteration if the room doesn't exist.
        }

        for _, neighbor := range room.Links {
            if _, visited := cameFrom[neighbor]; !visited && neighbor != g.Start {
                cameFrom[neighbor] = current
                queue = append(queue, neighbor)

                if neighbor == g.End {
                    path := []string{}
                    for at := g.End; at != ""; at = cameFrom[at] {
                        path = append([]string{at}, path...)
                    }
                    return path // Found a path from start to end.
                }
            }
        }
    }

    return nil // No path found.
}

// SimulateAntMovements simulates moving ants along the found path.
func SimulateAntMovements(numberOfAnts int, path []string) {
	for i := 0; i < numberOfAnts; i++ {
        fmt.Printf("L%d-%s ", i+1, path[0]) // Move ant to start position

        for j := 1; j < len(path); j++ { // Move through the path.
            fmt.Printf("L%d-%s ", i+1, path[j])
            fmt.Println()
        }
        
        fmt.Println() // New line after each ant's movement sequence.
    }
}

func main() {
	if len(os.Args) < 2 {
        fmt.Println("Usage: go run . <input_file>")
        return
    }

	filename := os.Args[1]
	numberOfAnts, graph, err := ReadInput(filename)
	if err != nil {
        fmt.Printf("Error reading input file: %v\n", err)
        return
    }

	path := graph.BFS()
	if len(path) == 0 { // Check only length since nil check is not needed.
        fmt.Println("ERROR: No valid path found from start to end.")
        return
    }

	fmt.Printf("%d\n", numberOfAnts)

	for _, roomName := range path { // Print all rooms in the path.
        fmt.Printf("%s ", roomName)
    }
    fmt.Println()

    SimulateAntMovements(numberOfAnts, path)
}
