package vectorize

import (
	"math"
	"sync"
)

// FIXME: operation
func vector1DTo1D(output, inputA, inputB <-chan float64, wg sync.WaitGroup) {
	var a, b float64
	for {
		select {
		case a <- inputA:
		case b <- inputB:
		default:
			break
		}
		output <- operation(a, b)
		wg.Done()
	}
}

func Vector1DTo1D(inputA, inputB chan float64, size int) (output chan float64) {
	var wg sync.WaitGroup
	wg.Add(size)
	output := make(chan float64, int(math.Log(size)))
	go vector1DTo1D(output, inputA, inputB, wg)
	go func() {
		defer close(inputA)
		defer close(inputB)
		wg.Wait()
	}()
	return
}

func vector2DTo2D(output, inputA, inputB <-chan <-chan float64, size int, wg sync.WaitGroup) {
	var a, b float64
	for {
		select {
		case a <- inputA:
		case b <- inputB:
		default:
			break
		}
		output <- Vector1DTo1D(a, b, shape[1])
		wg.Done()
	}
}

func Vector2DTo2D(inputA, inputB chan chan float64, shape [2]int) (output chan chan float64) {
	var wg sync.WaitGroup
	wg.Add(shape[0])
	output := make(chan chan float64, int(math.Log(shape[0])))
	go vector2DTo2D(output, inputA, inputB, shape[1], wg)
	go func() {
		defer close(inputA)
		defer close(inputB)
		wg.Wait()
	}()
	return
}

func vector1DStaticTo1D(inputA chan float64, inputB []float64) (output chan float64) {
	output := make(chan float64, int(math.Log(len(inputB))))
	for b := range inputB {
		for {
			select {
			case a, ok := <-inputA:
				if ok == true {
					output <- operation(a, b)
				}
			}
		}
	}
}
