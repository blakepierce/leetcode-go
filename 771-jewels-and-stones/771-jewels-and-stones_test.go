package leetcode

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{jewels: "aA", stones: "aAAbbbb"},
			want: 3,
		},
		{
			name: "example2",
			args: args{jewels: "z", stones: "ZZ"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStones("aAbCDkwox", "aAAbbbbkeWwbBdDooooOxkxkxaAC")
	}
}

func Test_numJewelsInStones2(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "noStones",
			args: args{jewels: "a", stones: ""},
			want: 0,
		},
		{
			name: "1 jewel-1 jewel",
			args: args{jewels: "a", stones: "a"},
			want: 1,
		},
		{
			name: "1 jewel-1 stone",
			args: args{jewels: "a", stones: "A"},
			want: 0,
		},
		{
			name: "1 jewel-2 jewels",
			args: args{jewels: "a", stones: "aa"},
			want: 2,
		},
		{
			name: "2 jewel-2 jewels",
			args: args{jewels: "aA", stones: "aA"},
			want: 2,
		},
		{
			name: "example1",
			args: args{jewels: "aA", stones: "aAAbbbb"},
			want: 3,
		},
		{
			name: "example2",
			args: args{jewels: "z", stones: "ZZ"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones2(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numJewelsInStones2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numJewelsInStones2("aAbCDkwox", "aAAbbbbkeWwbBdDooooOxkxkxaAC")
	}
}
