package goVectorize

import (
	"reflect"
	"testing"
)

func TestDot1D2D(t *testing.T) {
	type args struct {
		d1 []float64
		d2 [][]float64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []float64
		wantErr    bool
	}{
		{
			name: "dotProduct",
			args: args{
				[]float64{0, 1, 2, 3},
				[][]float64{
					{0, 1, 2, 3},
					{4, 5, 6, 7},
					{8, 9, 0, 1},
				},
			},
			wantOutput: []float64{14, 38, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := Dot1D2D(tt.args.d1, tt.args.d2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dot1D2D() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Dot1D2D() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestApplyFunction(t *testing.T) {
	type args struct {
		f     func(float64) (float64, error)
		scalr []float64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []float64
		wantErr    bool
	}{
		{
			name: "successFunction",
			args: args{
				f:     func(x float64) (float64, error) { return x / 2, nil },
				scalr: []float64{14, 38, 12},
			},
			wantOutput: []float64{7, 19, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := ApplyFunction(tt.args.f, tt.args.scalr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("ApplyFunction() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		a []float64
		b []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "add_1d_to_1d",
			args: args{
				[]float64{1, 2, 3},
				[]float64{2, 3, 1},
			},
			want: []float64{3, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOuterProduct(t *testing.T) {
	type args struct {
		column []float64
		row    []float64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput [][]float64
	}{
		{
			name: "outer_product",
			args: args{
				[]float64{1, 3, 5},
				[]float64{2, 4, 6, 8},
			},
			wantOutput: [][]float64{
				{2, 4, 6, 8},
				{6, 12, 18, 24},
				{10, 20, 30, 40},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := OuterProduct(tt.args.column, tt.args.row); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("OuterProduct() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestEntrywiseSum(t *testing.T) {
	type args struct {
		a [][]float64
		b [][]float64
	}
	tests := []struct {
		name    string
		args    args
		want    [][]float64
		wantErr bool
	}{
		{
			name: "add_2d",
			args: args{
				[][]float64{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
				[][]float64{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
			want: [][]float64{
				{2, 4, 6},
				{2, 4, 6},
				{2, 4, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EntrywiseSum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntrywiseSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntrywoseSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	type args struct {
		matrix   [][]float64
		lineSize int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput [][]float64
		wantErr    bool
	}{
		{
			name: "transpose",
			args: args{
				[][]float64{
					{1, 2, 3, 4},
					{1, 2, 3, 4},
					{1, 2, 3, 4},
				},
				4,
			},
			wantOutput: [][]float64{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
				{4, 4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := Transpose(tt.args.matrix, tt.args.lineSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transpose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Transpose() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMultiplyArrays(t *testing.T) {
	type args struct {
		a []float64
		b []float64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []float64
		wantErr    bool
	}{
		{
			name: "multiply_arrays",
			args: args{
				[]float64{1, 2, 3},
				[]float64{2, 3, 4},
			},
			wantOutput: []float64{2, 6, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := MultiplyArrays(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiplyArrays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MultiplyArrays() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
