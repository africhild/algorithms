// create test for containsDuplicate
package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "test1",
			args:       args{nums: []int{1, 2, 3, 1}},
			wantResult: true,
		},
		{
			name:       "test2",
			args:       args{nums: []int{1, 2, 3, 4}},
			wantResult: false,
		},
		{
			name:       "test3",
			args:       args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := containsDuplicate(tt.args.nums); gotResult != tt.wantResult {
				t.Errorf("containsDuplicate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
