package array

import "testing"

func Test_maxSubArrayBruteForce(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayBruteForce(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArrayBruteForce() = %v, want %v", got, tt.want)
			}
			if got := maxSubArrayRecursive(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArrayRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
