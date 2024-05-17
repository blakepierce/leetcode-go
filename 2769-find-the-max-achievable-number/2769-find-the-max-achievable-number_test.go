package leetcode

import "testing"

func Test_theMaximumAchievableX(t *testing.T) {
	type args struct {
		num int
		t   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{num: 4, t: 1},
			want: 6,
		},
		{

			name: "example2",
			args: args{num: 3, t: 2},
			want: 7,
		},
		{

			name: "example3",
			args: args{num: 30, t: 5},
			want: 40,
		},
		{

			name: "example4",
			args: args{num: 50, t: 50},
			want: 150,
		},
		{

			name: "example5",
			args: args{num: 1, t: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theMaximumAchievableX(tt.args.num, tt.args.t); got != tt.want {
				t.Errorf("theMaximumAchievableX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_theMaximumAchievableX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		theMaximumAchievableX(10, 10)
	}
}
