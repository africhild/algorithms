package main

import (
	"reflect"
	"testing"
)

func Test_twoSums(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
	},
	{
		name: "Example 3",
		args: args{
			nums: []int{3,3},
			target:  6,
		},
		want: []int{0, 1},
	},
}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := twoSums(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
