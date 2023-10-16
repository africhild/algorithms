// write test for isAnagram

package main

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "test1",
			args: args{
				s: "anagram",
				t: "nagaram",
		    },
			wantResult: true,
		},
		{
			name: "test2",
			args: args{
				s: "rat",
				t: "car",
			},
			wantResult: false,
		},
		{
			name: "test3",
			args: args{
				s: "",
				t: "",
			},
			wantResult: true,
		},
		{
			name: "test4",
			args: args{
				s: "a",
				t: "a",
			},
			wantResult: true,
		},
		{
			name: "test5",
			args: args{
				s: "a",
				t: "b",
			},
			wantResult: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := isAnagram(tt.args.s, tt.args.t); gotResult != tt.wantResult {
				t.Errorf("containsDuplicate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}