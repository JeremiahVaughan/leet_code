package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
		{
			name: "3",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums: []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
			},
			want: 7,
		},
		{
			name: "5",
			args: args{
				nums: []int{1, 2, 0, 1},
			},
			want: 3,
		},
		{
			name: "6",
			args: args{
				nums: []int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
