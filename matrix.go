package goVectorize

import "fmt"

func Dot1D2D(d1 []float64, d2 [][]float64) (output []float64, err error) {
	output = make([]float64, len(d2))
	for i, v := range d2 {
		if len(v) != len(d1) {
			err = fmt.Errorf(
				"Wrong size of a second dimension scalar: d2[i] = %d, d1 = %d",
				len(v),
				len(d1),
			)
			return
		}
		for j, k := range d1 {
			output[i] += v[j] * k
		}
	}
	return
}

func ApplyFunction(f func(float64) (float64, error), scalr []float64) (output []float64, err error) {
	var applied float64
	for _, v := range scalr {
		applied, err = f(v)
		if err != nil {
			return
		}
		output = append(output, applied)
	}
	return
}

func Add(a, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf(
			"Wrong size of arrays: a = %d, b = %d", len(a), len(b),
		)
	}

	for i, v := range b {
		a[i] += v
	}
	return a, nil
}

func EntrywiseSum(a, b [][]float64) ([][]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf(
			"Wrong size of matrices: a = %d, b = %d", len(a), len(b),
		)
	}

	for i, v := range b {
		if len(v) != len(a[i]) {
			return nil, fmt.Errorf(
				"Wrong size of arrays: a[i] = %d, b[i] = %d", len(a[i]), len(b),
			)
		}
		for j, k := range v {
			a[i][j] += k
		}
	}
	return a, nil
}

func OuterProduct(column, row []float64) (output [][]float64) {
	var rowOut []float64

	for _, i := range column {
		rowOut = nil
		for _, j := range row {
			rowOut = append(rowOut, i*j)
		}
		output = append(output, rowOut)
	}
	return
}

func Transpose(matrix [][]float64, lineSize int) (output [][]float64, err error) {
	output = make([][]float64, lineSize)
	for i, v := range matrix {
		if len(v) != lineSize {
			err = fmt.Errorf(
				"Inconsistent matrix line size. Expected: %d. Actual: %d",
				lineSize, len(v),
			)
			return
		}
		for j := range v {
			// To do transpose in place switch matrix items
			//matrix[i][j], matrix[j][i] = matrix[j][j], matrix[i][j]
			output[j] = append(output[j], matrix[i][j])
		}
	}
	return
}
