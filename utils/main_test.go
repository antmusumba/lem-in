package utils

import (
	"lem-in/models"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestGenerateTurns(t *testing.T) {
	type args struct {
		option map[int][]int
		paths  []models.Path
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single path",
			args: args{
				option: map[int][]int{0: {3}},
				paths: []models.Path{
					{Rooms: []string{"1", "2", "3"}},
				},
			},
			want: 2,
		},
		{
			name: "Multiple paths",
			args: args{
				option: map[int][]int{0: {2}, 1: {1}},
				paths: []models.Path{
					{Rooms: []string{"1", "2", "3"}},
					{Rooms: []string{"1", "4", "3"}},
				},
			},
			want: 2,
		},
		{
			name: "Uneven distribution",
			args: args{
				option: map[int][]int{0: {3}, 1: {1}},
				paths: []models.Path{
					{Rooms: []string{"1", "2", "3"}},
					{Rooms: []string{"1", "4", "5", "3"}},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateTurns(tt.args.option, tt.args.paths); got != tt.want {package utils

import (
	"lem-in/resources"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateTurns(t *testing.T) {
	type args struct {
		option map[int][]int
		paths  []resources.Path
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single path",
			args: args{
				option: map[int][]int{0: {3}},
				paths: []resources.Path{
					{RoomsInThePath: []string{"1", "2", "3"}},
				},
			},
			want: 2,
		},
		{
			name: "Multiple paths",
			args: args{
				option: map[int][]int{0: {2}, 1: {1}},
				paths: []resources.Path{
					{RoomsInThePath: []string{"1", "2", "3"}},
					{RoomsInThePath: []string{"1", "4", "3"}},
				},
			},
			want: 2,
		},
		{
			name: "Uneven distribution",
			args: args{
				option: map[int][]int{0: {3}, 1: {1}},
				paths: []resources.Path{
					{RoomsInThePath: []string{"1", "2", "3"}},
					{RoomsInThePath: []string{"1", "4", "5", "3"}},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTurns(tt.args.option, tt.args.paths)
			if got != tt.want {
				t.Errorf("GenerateTurns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileContents(t *testing.T) {
	// Create a temporary directory for test files

				t.Errorf("GenerateTurns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileContents(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "filecontents_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up after the test

	tests := []struct {
		name     string
		content  string
		expected []string
		wantErr  bool
	}{
		{
			name: "valid file with all types of lines",
			content: `3
##start
1 23 3
#comment to ignore
2 16 7
##end
0 9 5
`,
			expected: []string{"3", "##start", "1 23 3", "2 16 7", "##end", "0 9 5"},
			wantErr:  false,
		},
		{
			name: "empty file",
			content: ``,
			expected: []string{},
			wantErr:  false,
		},
		{
			name: "only comments",
			content: `#comment1
#comment2
#comment3`,
			expected: []string{},
			wantErr:  false,
		},
		{
			name: "special commands and connections",
			content: `##start
##end
1-2
2-3`,
			expected: []string{"##start", "##end", "1-2", "2-3"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for this test case
			tmpFile := filepath.Join(tmpDir, "test.txt")
			err := os.WriteFile(tmpFile, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Test the fileContents function
			got, err := fileContents(tmpFile)
			
			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("fileContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Compare results
			if len(got) != len(tt.expected) {
				t.Errorf("fileContents() got %d lines, want %d lines", len(got), len(tt.expected))
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("fileContents() line %d = %q, want %q", i, got[i], tt.expected[i])
				}
			}
		})
	}

	// Test non-existent file
	t.Run("non-existent file", func(t *testing.T) {
		_, err := fileContents("/non/existent/file.txt")
		if err == nil {
			t.Error("fileContents() expected error for non-existent file, got nil")
		}
	})
}

func TestParseRoom(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		colony  *models.AntColony
		want    string
		wantErr bool
	}{
		{
			name: "valid room",
			line: "room1 23 45",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "room1",
			wantErr: false,
		},
		{
			name: "invalid format - too few parts",
			line: "room1 23",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid format - too many parts",
			line: "room1 23 45 67",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid X coordinate",
			line: "room1 abc 45",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid Y coordinate",
			line: "room1 23 def",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "duplicate coordinates",
			line: "room2 23 45",
			colony: &models.AntColony{
				Rooms: []models.Room{
					{Name: "room1", Coord_X: 23, Coord_Y: 45},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "multiple valid rooms",
			line: "room2 24 46",
			colony: &models.AntColony{
				Rooms: []models.Room{
					{Name: "room1", Coord_X: 23, Coord_Y: 45},
				},
			},
			want:    "room2",
			wantErr: false,
		},
		{
			name: "invalid room name starts with L",
			line: "L1 1 2",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid room name starts with #",
			line: "#room 1 2",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid room name with space",
			line: "room one 1 2",
			colony: &models.AntColony{
				Rooms: make([]models.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRoom(tt.line, tt.colony)
			
			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check return value
			if got != tt.want {
				t.Errorf("parseRoom() = %v, want %v", got, tt.want)
			}

			// For valid cases, check if room was actually added to colony
			if err == nil {
				found := false
				for _, room := range tt.colony.Rooms {
					if room.Name == got {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("parseRoom() room %s was not added to colony", got)
				}
			}
		})
	}
}

func TestValidateRoomName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid room name",
			input:   "room1",
			wantErr: false,
		},
		{
			name:    "Room name starting with L",
			input:   "L1",
			wantErr: true,
		},
		{
			name:    "Room name starting with #",
			input:   "#room",
			wantErr: true,
		},
		{
			name:    "Room name with space",
			input:   "room 1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRoomName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateRoomName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseConnection(t *testing.T) {
	// Reset the Existinglink map before each test
	models.Existinglink = make(map[string]bool)

	tests := []struct {
		name    string
		line    string
		colony  *models.AntColony
		wantErr bool
	}{
		{
			name: "valid connection",
			line: "room1-room2",
			colony: &models.AntColony{
				Links: map[string][]string{
					"room1": {},
					"room2": {},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid format - no hyphen",
			line: "room1room2",
			colony: &models.AntColony{
				Links: map[string][]string{},
			},
			wantErr: true,
		},
		{
			name: "invalid format - multiple hyphens",
			line: "room1-room2-room3",
			colony: &models.AntColony{
				Links: map[string][]string{},
			},
			wantErr: true,
		},
		{
			name: "self connection",
			line: "room1-room1",
			colony: &models.AntColony{
				Links: map[string][]string{
					"room1": {},
				},
			},
			wantErr: true,
		},
		{
			name: "first room doesn't exist",
			line: "nonexistent-room2",
			colony: &models.AntColony{
				Links: map[string][]string{
					"room2": {},
				},
			},
			wantErr: true,
		},
		{
			name: "second room doesn't exist",
			line: "room1-nonexistent",
			colony: &models.AntColony{
				Links: map[string][]string{
					"room1": {},
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate connection",
			line: "room1-room2",
			colony: &models.AntColony{
				Links: map[string][]string{
					"room1": {},
					"room2": {},
				},
			},
			wantErr: true,
			// This will error because we'll try to add the same connection twice
		},
		{
			name: "valid reverse connection",
			line: "room2-room1",
			colony: &models.AntColony{
				Links: map[string][]string{
					"room1": {},
					"room2": {},
				},
			},
			wantErr: true, // Should error because room1-room2 was already added
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For duplicate connection test, add the first connection
			if tt.name == "duplicate connection" || tt.name == "valid reverse connection" {
				models.Existinglink["room1room2"] = true
				models.Existinglink["room2room1"] = true
			}

			err := parseConnection(tt.line, tt.colony)
			
			// Check error cases
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// For valid cases, check if connection was actually added to colony
			if err == nil {
				parts := strings.Split(tt.line, "-")
				
				// Check forward connection
				found := false
				for _, conn := range tt.colony.Links[parts[0]] {
					if conn == parts[1] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("parseConnection() forward connection %s->%s was not added", parts[0], parts[1])
				}

				// Check reverse connection
				found = false
				for _, conn := range tt.colony.Links[parts[1]] {
					if conn == parts[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("parseConnection() reverse connection %s->%s was not added", parts[1], parts[0])
				}

				// Check if links were added to Existinglink map
				link := parts[0] + parts[1]
				link2 := parts[1] + parts[0]
				if !models.Existinglink[link] || !models.Existinglink[link2] {
					t.Errorf("parseConnection() links not properly added to Existinglink map")
				}
			}
		})
	}
}

func TestMoveAnts(t *testing.T) {
	type args struct {
		paths       []models.Path
		antsPerRoom map[int][]int
		turns       int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Single path with one ant",
			args: args{
				paths: []models.Path{
					{Rooms: []string{"start", "1", "end"}},
				},
				antsPerRoom: map[int][]int{
					0: {1},
				},
				turns: 2,
			},
			want: []string{"L1-1", "L1-end"},
		},
		{
			name: "Two paths with two ants",
			args: args{
				paths: []models.Path{
					{Rooms: []string{"start", "1", "end"}},
					{Rooms: []string{"start", "2", "end"}},
				},
				antsPerRoom: map[int][]int{
					0: {1},
					1: {2},
				},
				turns: 2,
			},
			want: []string{"L1-1 L2-2", "L1-end L2-end"},
		},
		{
			name: "Single path with multiple ants",
			args: args{
				paths: []models.Path{
					{Rooms: []string{"start", "1", "2", "end"}},
				},
				antsPerRoom: map[int][]int{
					0: {1, 2, 3},
				},
				turns: 5,
			},
			want: []string{"L1-1", "L1-2 L2-1", "L1-end L2-2 L3-1", "L2-end L3-2", "L3-end"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MoveAnts(tt.args.paths, tt.args.antsPerRoom, tt.args.turns)

			// Trim trailing spaces from both `got` and `want` for consistency
			for i := range got {
				got[i] = strings.TrimSpace(got[i])
			}
			for i := range tt.want {
				tt.want[i] = strings.TrimSpace(tt.want[i])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPaths(t *testing.T) {
	tests := []struct {
		name          string
		colony        *models.AntColony
		wantPaths    []models.Path
		wantAntsPerPath map[int][]int
		wantTurns    int
	}{
		{
			name: "Simple path with one route",
			colony: &models.AntColony{
				NumberOfAnts: 3,
				Start:       "1",
				End:         "0",
				Rooms: []models.Room{
					{Name: "1"},
					{Name: "2"},
					{Name: "0"},
				},
				Links: map[string][]string{
					"1": {"2"},
					"2": {"1", "0"},
					"0": {"2"},
				},
			},
			wantPaths: []models.Path{
				{Rooms: []string{"1", "2", "0"}},
			},
			wantAntsPerPath: map[int][]int{
				0: {1, 2, 3},
			},
			wantTurns: 4,
		},
		{
			name: "Multiple possible paths",
			colony: &models.AntColony{
				NumberOfAnts: 4,
				Start:       "start",
				End:         "end",
				Rooms: []models.Room{
					{Name: "start"},
					{Name: "room1"},
					{Name: "room2"},
					{Name: "room3"},
					{Name: "end"},
				},
				Links: map[string][]string{
					"start": {"room1", "room2"},
					"room1": {"start", "room3"},
					"room2": {"start", "end"},
					"room3": {"room1", "end"},
					"end":   {"room2", "room3"},
				},
			},
			wantPaths: []models.Path{
				{Rooms: []string{"start", "room2", "end"}},
				{Rooms: []string{"start", "room1", "room3", "end"}},
			},
			wantAntsPerPath: map[int][]int{
				0: {1, 2, 4},
				1: {3},
			},
			wantTurns: 4,
		},
		{
			name: "Path with cycle detection",
			colony: &models.AntColony{
				NumberOfAnts: 2,
				Start:       "start",
				End:         "end",
				Rooms: []models.Room{
					{Name: "start"},
					{Name: "A"},
					{Name: "B"},
					{Name: "end"},
				},
				Links: map[string][]string{
					"start": {"A", "B"},
					"A":     {"start", "B"},
					"B":     {"A", "start", "end"},
					"end":   {"B"},
				},
			},
			wantPaths: []models.Path{
				{Rooms: []string{"start", "B", "end"}},
			},
			wantAntsPerPath: map[int][]int{
				0: {1, 2},
			},
			wantTurns: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPaths, gotAntsPerPath, gotTurns := FindPaths(tt.colony)

			// Check number of paths
			if len(gotPaths) != len(tt.wantPaths) {
				t.Errorf("FindPaths() got %v paths, want %v paths", len(gotPaths), len(tt.wantPaths))
				return
			}

			// Check each path
			for i, wantPath := range tt.wantPaths {
				if !reflect.DeepEqual(gotPaths[i].Rooms, wantPath.Rooms) {
					t.Errorf("FindPaths() path %d = %v, want %v", i, gotPaths[i].Rooms, wantPath.Rooms)
				}
			}

			// Check ants per path distribution
			if !reflect.DeepEqual(gotAntsPerPath, tt.wantAntsPerPath) {
				t.Errorf("FindPaths() ants distribution = %v, want %v", gotAntsPerPath, tt.wantAntsPerPath)
			}

			// Check number of turns
			if gotTurns != tt.wantTurns {
				t.Errorf("FindPaths() turns = %v, want %v", gotTurns, tt.wantTurns)
			}
		})
	}
}

func TestContainsRoom(t *testing.T) {
	tests := []struct {
		name string
		path []string
		room string
		want bool
	}{
		{
			name: "Room exists in path",
			path: []string{"start", "room1", "room2", "end"},
			room: "room1",
			want: true,
		},
		{
			name: "Room does not exist in path",
			path: []string{"start", "room1", "room2", "end"},
			room: "room3",
			want: false,
		},
		{
			name: "Empty path",
			path: []string{},
			room: "room1",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsRoom(tt.path, tt.room); got != tt.want {
				t.Errorf("containsRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptimizedPaths1(t *testing.T) {
	tests := []struct {
		name  string
		paths []models.Path
		want  []models.Path
	}{
		{
			name: "Non-overlapping paths",
			paths: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
				{Rooms: []string{"start", "B", "end"}},
			},
			want: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
				{Rooms: []string{"start", "B", "end"}},
			},
		},
		{
			name: "Overlapping paths",
			paths: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
				{Rooms: []string{"start", "A", "B", "end"}},
			},
			want: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OptimizedPaths1(tt.paths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptimizedPaths1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptimizedPaths2(t *testing.T) {
	tests := []struct {
		name   string
		paths  []models.Path
		colony *models.AntColony
		want   []models.Path
	}{
		{
			name: "Paths within ant count limit",
			paths: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
				{Rooms: []string{"start", "B", "end"}},
			},
			colony: &models.AntColony{
				NumberOfAnts: 4,
			},
			want: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
				{Rooms: []string{"start", "B", "end"}},
			},
		},
		{
			name: "Paths exceeding ant count limit",
			paths: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
				{Rooms: []string{"start", "B", "C", "D", "end"}},
			},
			colony: &models.AntColony{
				NumberOfAnts: 2,
			},
			want: []models.Path{
				{Rooms: []string{"start", "A", "end"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OptimizedPaths2(tt.paths, tt.colony)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptimizedPaths2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaceAnts(t *testing.T) {
	tests := []struct {
		name    string
		colony  *models.AntColony
		paths   []models.Path
		want    map[int][]int
	}{
		{
			name: "Single path",
			colony: &models.AntColony{
				NumberOfAnts: 3,
				Start:       "start",
				End:         "end",
			},
			paths: []models.Path{
				{Rooms: []string{"start", "1", "end"}},
			},
			want: map[int][]int{
				0: {1, 2, 3},
			},
		},
		{
			name: "Two equal length paths",
			colony: &models.AntColony{
				NumberOfAnts: 4,
				Start:       "start",
				End:         "end",
			},
			paths: []models.Path{
				{Rooms: []string{"start", "1", "end"}},
				{Rooms: []string{"start", "2", "end"}},
			},
			want: map[int][]int{
				0: {1, 3},
				1: {2, 4},
			},
		},
		{
			name: "Two paths with different lengths",
			colony: &models.AntColony{
				NumberOfAnts: 5,
				Start:       "start",
				End:         "end",
			},
			paths: []models.Path{
				{Rooms: []string{"start", "1", "end"}},
				{Rooms: []string{"start", "2", "3", "end"}},
			},
			want: map[int][]int{
				0: {1, 2, 4},
				1: {3, 5},
			},
		},
		{
			name: "Three paths with varying lengths",
			colony: &models.AntColony{
				NumberOfAnts: 6,
				Start:       "start",
				End:         "end",
			},
			paths: []models.Path{
				{Rooms: []string{"start", "1", "end"}},
				{Rooms: []string{"start", "2", "3", "end"}},
				{Rooms: []string{"start", "4", "5", "6", "end"}},
			},
			want: map[int][]int{
				0: {1, 2, 4, 6},
				1: {3, 5},
			},
		},
		{
			name: "No ants",
			colony: &models.AntColony{
				NumberOfAnts: 0,
				Start:       "start",
				End:         "end",
			},
			paths: []models.Path{
				{Rooms: []string{"start", "1", "end"}},
			},
			want: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlaceAnts(tt.colony, tt.paths)
			
			// Check if the number of paths with assignments matches
			if len(got) != len(tt.want) {
				t.Errorf("PlaceAnts() got %d paths with assignments, want %d", len(got), len(tt.want))
				return
			}

			// Check each path's ant assignments
			for pathIndex, wantAnts := range tt.want {
				gotAnts, exists := got[pathIndex]
				if !exists {
					t.Errorf("PlaceAnts() missing assignments for path %d", pathIndex)
					continue
				}

				// Check number of ants assigned to this path
				if len(gotAnts) != len(wantAnts) {
					t.Errorf("PlaceAnts() path %d got %d ants, want %d ants", pathIndex, len(gotAnts), len(wantAnts))
					continue
				}

				// Check if all expected ants are present
				for i, wantAnt := range wantAnts {
					if gotAnts[i] != wantAnt {
						t.Errorf("PlaceAnts() path %d, ant index %d got ant %d, want ant %d", 
							pathIndex, i, gotAnts[i], wantAnt)
					}
				}
			}
		})
	}
}
