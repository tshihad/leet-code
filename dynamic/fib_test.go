package dynamic

import "testing"

func Test_normalFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 5,
			},
			want: 5,
		},
		{
			name: "1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				n: 50,
			},
			want: 12586269025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("normalFib() = %v, want %v", got, tt.want)
			}
		})
	}
}
