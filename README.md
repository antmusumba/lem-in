# Lem-in

A Go implementation of an ant colony pathfinding simulator. The program finds the most efficient way to move ants through a colony from start to end room, handling multiple paths and traffic optimization.

## Overview

The program reads a colony description from a file and determines the optimal paths for ants to traverse from the start room to the end room. It handles various colony configurations, including multiple paths, room connections, and validates the input format.

## Features

- Reads colony configuration from a file
- Validates input format and colony structure
- Finds optimal paths using breadth-first search
- Handles multiple possible paths
- Optimizes ant distribution across paths
- Detects and handles invalid configurations
- Provides detailed error messages

## Input Format

The input file should follow this format:

```
number_of_ants
##start
start_room x y
room1 x y
room2 x y
##end
end_room x y
room1-room2
start_room-room1
room2-end_room
```

### Rules

- Room names cannot:
  - Start with 'L' or '#'
  - Contain spaces
- Each room must have integer coordinates
- Two rooms cannot share the same coordinates
- Each tunnel connects exactly two rooms
- A room can connect to multiple other rooms
- Comments start with '#'

## Usage

```bash
go run . [filename]
```

Example:
```bash
go run . example.txt
```

### Example Input File
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

### Example Output
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

## Error Handling

The program provides specific error messages for various invalid scenarios:
- Invalid number of ants
- Missing start/end rooms
- Invalid room names
- Invalid coordinates
- Invalid connections
- Duplicate rooms/coordinates
- Invalid file format

## Project Structure

```
lem-in/
├── cmd/
│   └── main.go           # Main entry point
├── models/
│   └── models.go         # Data structures
├── utils/
│   ├── findpaths.go      # Path finding logic
│   ├── generateturns.go  # Turn generation
│   ├── parseFile.go      # File parsing
│   └── placeants.go      # Ant placement logic
└── README.md
```

## Testing

The project includes comprehensive unit tests. Run them using:

```bash
go test ./...
```

## Implementation Details

1. **File Parsing**: Validates input format and builds colony structure
2. **Path Finding**: Uses BFS to find all possible paths from start to end
3. **Path Optimization**: Selects optimal paths based on length and congestion
4. **Ant Distribution**: Distributes ants across paths to minimize total moves
5. **Move Generation**: Generates valid moves for each turn

## Error Messages

The program returns specific error messages in the format:
```
ERROR: invalid data format, [specific reason]
```

Examples:
- `ERROR: invalid data format, no start room found`
- `ERROR: invalid data format, no end room found`
- `ERROR: invalid data format, invalid room name: L1`

## Contributors
## Contributing

Feel free to submit issues and enhancement requests.

