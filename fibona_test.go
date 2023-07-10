package fibona

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"insert 0", args{0}, 0},
		{"insert 1", args{1}, 1},
		{"insert 2", args{2}, 1},
		{"insert 3", args{3}, 2},
		{"insert 10", args{10}, 55},
		{"insert 20", args{20}, 6765},
		{"insert 30", args{30}, 832040},
		{"insert 40", args{40}, 102334155},
		{"insert 50", args{50}, 12586269025},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Fibonacci(tt.args.n), "fibonacci(%v)", tt.args.n)
		})
	}

}
