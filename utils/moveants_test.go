package utils

import (
	"lem-in/models"
	"reflect"
	"testing"
)

func TestMoveAnts(t *testing.T) {
	type args struct {
		paths       []models.Path
		antsperroom map[int][]int
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
				antsperroom: map[int][]int{
					0: {1},
				},
				turns: 2,
			},
			want: []string{"L1-1 ", "L1-end "},
		},
		{
			name: "Two paths with two ants",
			args: args{
				paths: []models.Path{
					{Rooms: []string{"start", "1", "end"}},
					{Rooms: []string{"start", "2", "end"}},
				},
				antsperroom: map[int][]int{
					0: {1},
					1: {2},
				},
				turns: 2,
			},
			want: []string{"L1-1 L2-2 ", "L1-end L2-end "},
		},
		{
			name: "Single path with multiple ants",
			args: args{
				paths: []models.Path{
					{Rooms: []string{"start", "1", "2", "end"}},
				},
				antsperroom: map[int][]int{
					0: {1, 2, 3},
				},
				turns: 5,
			},
			want: []string{"L1-1 ", "L1-2 L2-1 ", "L1-end L2-2 L3-1 ", "L2-end L3-2 ", "L3-end "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveAnts(tt.args.paths, tt.args.antsperroom, tt.args.turns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
