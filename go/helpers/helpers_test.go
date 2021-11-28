package helpers

import (
	"testing"
)

func TestSumOfFloat64Array(t *testing.T) {
	type args struct {
		test []float64
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{"single", args{[]float64{1.0}}, 1.0},
		{"2 values", args{[]float64{1.0, 10.0}}, 11.0},
		{"multiple values", args{[]float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}}, 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SumOfFloat64Array(tt.args.test); gotResult != tt.wantResult {
				t.Errorf("SumOfFloat64Array() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSumOfIntArray(t *testing.T) {
	type args struct {
		test []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"single", args{[]int{1}}, 1},
		{"2 values", args{[]int{1, 10}}, 11},
		{"multiple values", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SumOfIntArray(tt.args.test); gotResult != tt.wantResult {
				t.Errorf("SumOfIntArray() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_Abs(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"positive", args{3000}, 3000},
		{"negative", args{-9000}, 9000},
		{"zero", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
