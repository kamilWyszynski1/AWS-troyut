package service

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		i1 int
		i2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				i1: 1,
				i2: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.i1, tt.args.i2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
