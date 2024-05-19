package leetcode

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 2, 3, 1, 1, 3}},
			want: 4,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairsBest(tt.args.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numIdenticalPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numIdenticalPairsWorse([]int{1, 1, 2, 3, 2, 1, 5})
	}
}

func Benchmark_numIdenticalPairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numIdenticalPairsBest([]int{1, 1, 2, 3, 2, 1, 5})
	}
}
