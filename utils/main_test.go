package utils

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
					{Rooms: []string{"1", "2", "3"}},
				},
			},
			want: 2,
		},
		{
			name: "Multiple paths",
			args: args{
				option: map[int][]int{0: {2}, 1: {1}},
				paths: []resources.Path{
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
				paths: []resources.Path{
					{Rooms: []string{"1", "2", "3"}},
					{Rooms: []string{"1", "4", "5", "3"}},
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
		colony  *resources.AntColony
		want    string
		wantErr bool
	}{
		{
			name: "valid room",
			line: "room1 23 45",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "room1",
			wantErr: false,
		},
		{
			name: "invalid format - too few parts",
			line: "room1 23",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid format - too many parts",
			line: "room1 23 45 67",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid X coordinate",
			line: "room1 abc 45",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid Y coordinate",
			line: "room1 23 def",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "duplicate coordinates",
			line: "room2 23 45",
			colony: &resources.AntColony{
				Rooms: []resources.Room{
					{Name: "room1", Coord_X: 23, Coord_Y: 45},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "multiple valid rooms",
			line: "room2 24 46",
			colony: &resources.AntColony{
				Rooms: []resources.Room{
					{Name: "room1", Coord_X: 23, Coord_Y: 45},
				},
			},
			want:    "room2",
			wantErr: false,
		},
		{
			name: "invalid room name starts with L",
			line: "L1 1 2",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid room name starts with #",
			line: "#room 1 2",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid room name with space",
			line: "room one 1 2",
			colony: &resources.AntColony{
				Rooms: make([]resources.Room, 0),
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
			name:    "Room name starts with L",
			input:   "Lroom",
			wantErr: true,
		},
		{
			name:    "Room name starts with #",
			input:   "#room",
			wantErr: true,
		},
		{
			name:    "Room name with spaces",
			input:   "room name",
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

