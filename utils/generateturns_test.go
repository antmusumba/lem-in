package utils

import (
	"lem-in/models"
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
			if got := GenerateTurns(tt.args.option, tt.args.paths); got != tt.want {
				t.Errorf("GenerateTurns() = %v, want %v", got, tt.want)
			}
		})
	}
}
