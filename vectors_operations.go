package vectorize

import (
	"math"
	"sync"
)

func vector1DTo1D(output, inputA, inputB <-chan float64, operation func(float64) float64, wg sync.WaitGroup) {
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

func Vector1DTo1D(inputA, inputB <-chan float64, size int, operation func(float64, float64) float64) (output chan float64) {
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

func vector2DTo2D(output, inputA, inputB <-chan <-chan float64, size int, operation func(float64, float64) float64, wg sync.WaitGroup) {
	var a, b float64
	for {
		select {
		case a <- inputA:
		case b <- inputB:
		default:
			break
		}
		output <- Vector1DTo1D(a, b, size)
		wg.Done()
	}
}

func Vector2DTo2D(inputA, inputB chan chan float64, shape [2]int, operation func(float64, float64) float64) (output chan chan float64) {
	var wg sync.WaitGroup
	wg.Add(shape[0])
	output := make(chan chan float64, int(math.Log(shape[0])))
	go vector2DTo2D(output, inputA, inputB, shape[1], operation, wg)
	go func() {
		defer close(inputA)
		defer close(inputB)
		wg.Wait()
	}()
	return
}

func vector1DStaticTo1D(output, inputA chan float64, inputB []float64, operation func(float64, float64) float64, wg sync.WaitGroup) {
	for b := range inputB {
	LOOP:
		for {
			select {
			case a, ok := <-inputA:
				if ok == true {
					output <- operation(a, b)
				} else {
					break LOOP
				}
			default:
				break LOOP
			}
		}
	}
}

func Vector1DStaticTo1D(inputA chan float64, inputB []float64, operation func(float64, float64) float64) (output chan float64) {
	var wg sync.WaitGroup
	output = make(chan float64, math.Log(len(inputB)))
	wg.Add(len(inputB))
	go vector1DStaticTo1D(output, inputA, inputB, operation, wg)
	go func() {
		defer close(inputA)
		wg.Wait()
	}()
	return
}

type LockingSlice struct {
	sync.Mutex
	Vector []float64
}

func mutate1D(inputA chan float64, inputB *LockingSlice, operation func(float64, float64) float64, wg sync.WaitGroup) {
LOOP:
	for i, v := range inputB {
		for {
			select {
			case j, ok := <-inputA:
				if ok == false {
					break LOOP
				} else {
					inputB[i] = operation(j, v)
				}
			default:
				break LOOP
			}
			wg.Done()
		}
	}

}

func Mutate1D(inputA chan float64, inputB *LockingSlice, operation func(float64, float64) float64) *LockingSlice {
	var wg sync.WaitGroup
	wg.Add(math.Log(len(inputB)))
	inputB.Lock()
	go mutate1D(inputA, inputB, operation, wg)
	go func() {
		defer close(inputA)
		defer inputB.Unlock()
		wg.Wait()
	}()
	return inputB
}
