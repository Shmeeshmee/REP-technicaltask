package calculate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_calculator(t *testing.T) {
	type args struct {
		target  int
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test 2",
			args: args{
				target:  263,
				numbers: []int{23, 31, 53},
			},
			want: []string{"2 X 23", "7 X 31"},
		},
		{
			name: "test 1.1",
			args: args{
				target:  50241,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"10 X 5000", "1 X 250"},
		},
		{
			name: "test 1.2",
			args: args{
				target:  250,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250"},
		},
		{
			name: "test 1.3",
			args: args{
				target:  251,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 500"},
		},
		{
			name: "test 1.4",
			args: args{
				target:  501,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250", "1 X 500"},
		},
		{
			name: "test 1.4",
			args: args{
				target:  12001,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250", "1 X 2000", "2 X 5000"},
		},
		{
			name: "test 3.1 - mil",
			args: args{
				target:  5000000,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1000 X 5000"},
		},
		{
			name: "test 3.2 - mil + 1",
			args: args{
				target:  5000001,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250", "1000 X 5000"},
		},
		{
			name: "test 3.3 - mil + 249",
			args: args{
				target:  5000249,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250", "1000 X 5000"},
		},
		{
			name: "test 3.4 - bil",
			args: args{
				target:  5000000249,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250", "1000000 X 5000"},
		},
		{
			name: "test 3.5 - tril",
			args: args{
				target:  5000000000249,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250", "1000000000 X 5000"},
		},
		{
			name: "test 3.6 - max int",
			args: args{
				target:  9223372036854775750,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1844674407370955 X 5000", "1 X 250", "1 X 500"},
		},
		{
			name: "test 4.1 - high low no's",
			args: args{
				target:  1799,
				numbers: []int{900, 1},
			},
			want: []string{"1 X 900", "899 X 1"},
		},
		{
			name: "test 4.2 - high low no's",
			args: args{
				target:  1799,
				numbers: []int{900, 2, 1},
			},
			want: []string{"1 X 900", "449 X 2", "1 X 1"},
		},
	}
	less := func(a, b string) bool { return a < b }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := calculator(tt.args.target, tt.args.numbers)
			if r := cmp.Diff(gotResult.Packs, tt.want, cmpopts.SortSlices(less)); r != "" {
				t.Errorf("calculator() = %v, want %v", gotResult.Packs, tt.want)
			}
		})
	}
}
