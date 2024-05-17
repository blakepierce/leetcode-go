package leetcode

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 2, 1}},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "example2",
			args: args{nums: []int{1, 3, 2, 1}},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenation2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getConcatenation([]int{1, 2, 3, 4, 5})
	}
}

func Benchmark_getConcatenation2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getConcatenation2([]int{1, 2, 3, 4, 5})
	}
}
