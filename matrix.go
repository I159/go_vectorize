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
