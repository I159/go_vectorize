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
