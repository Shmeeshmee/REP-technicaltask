package calculate

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var maxInt = math.MaxInt64

func Test_findCombinationsHelper(t *testing.T) {
	type args struct {
		target             int
		numbers            []int
		startIndex         int
		currentCombination []int
		shortestResult     []int
		shortestLength     *int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 2",
			args: args{
				target:         263,
				numbers:        []int{23, 31, 53},
				startIndex:     0,
				shortestLength: &maxInt,
			},
			want: []int{23, 23, 31, 31, 31, 31, 31, 31, 31},
		},
		{
			name: "test 1.1",
			args: args{
				target:         1,
				numbers:        []int{250, 500, 1000, 2000, 5000},
				startIndex:     0,
				shortestLength: &maxInt,
			},
		},
		{
			name: "test 1.2",
			args: args{
				target:         250,
				numbers:        []int{250, 500, 1000, 2000, 5000},
				startIndex:     0,
				shortestLength: &maxInt,
			},
			want: []int{250},
		},
		{
			name: "test 1.3",
			args: args{
				target:         251,
				numbers:        []int{250, 500, 1000, 2000, 5000},
				startIndex:     0,
				shortestLength: &maxInt,
			},
		},
		{
			name: "test 1.4",
			args: args{
				target:         501,
				numbers:        []int{250, 500, 1000, 2000, 5000},
				startIndex:     0,
				shortestLength: &maxInt,
			},
		},
		{
			name: "test 1.4",
			args: args{
				target:         120001,
				numbers:        []int{250, 500, 1000, 2000, 5000},
				startIndex:     0,
				shortestLength: &maxInt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCombinationsHelper(tt.args.target, tt.args.numbers, tt.args.startIndex, tt.args.currentCombination, tt.args.shortestResult, tt.args.shortestLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCombinationsHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringify(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name:       "test.1",
			args:       args{a: []int{23, 23, 31, 31, 31, 31, 31, 31, 31}},
			wantResult: []string{"2 X 23", "7 X 31"},
		},
		{
			name:       "test.2",
			args:       args{a: []int{5000, 2000, 5000, 250, 2000}},
			wantResult: []string{"2 X 5000", "2 X 2000", "1 X 250"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := stringify(tt.args.a); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("stringify() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_exact(t *testing.T) {
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
				target:  1,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
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
		},
		{
			name: "test 1.4",
			args: args{
				target:  501,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
		},
		{
			name: "test 1.4",
			args: args{
				target:  120001,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exact(tt.args.target, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exact() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				target:  1,
				numbers: []int{250, 500, 1000, 2000, 5000},
			},
			want: []string{"1 X 250"},
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
	}
	less := func(a, b string) bool { return a < b }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := calculator(tt.args.target, tt.args.numbers)
			if cmp.Diff(gotResult, tt.want, cmpopts.SortSlices(less)) != "" {
				t.Errorf("calculator() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
