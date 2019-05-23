package main

import (
	"errors"
	"fmt"
)

func main() {
	//Examples
	testCases := [][]float32{
		[]float32{},
		[]float32{1.0},
		[]float32{1.32, 1.32},
		[]float32{1.32, 2.3323},
		[]float32{21.3, 3.3232},
		[]float32{2.0, 2.0, 2.0, 2.0, 2.0, 2.0},
		[]float32{12.0, 4.0, 8.1, 4.4, 6.767},
		[]float32{4.4, 54.5, 12.0, 43.3, 63.0, 12.2},
		[]float32{5.0, 5.0, 12.2, 3.0, 3.33, 4.99},
		[]float32{34.5, 21.2, 1.0, 43.4},
		[]float32{1.0, 1.0, 2.0, 2.0, 1.0},
		[]float32{-34.5, 21.2, 1.0, -43.4},
		[]float32{-1.0, 0, -2.0, 0, 1.0},
		[]float32{-1.0, 0},
	}

	for _, tc := range testCases {
		max1, max2, err := twoBiggestEls(tc)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("The biggest nubmber is: %v; the number next to it: %v\n", max2, max1)
		}
	}
}

func twoBiggestEls(numbers []float32) (float32, float32, error) {
	if len(numbers) < 2 {
		return 0, 0, errors.New("Error! Input slice is too short")
	}

	previos, current := numbers[0], numbers[1]
	if previos > current {
		previos, current = current, previos
	}

	if len(numbers[2:]) == 0 {
		return previos, current, nil
	}

	for _, n := range numbers[2:] {
		if n == current {
			continue
		}

		if n > current {
			previos, current = current, n
			continue
		}

		if n > previos {
			previos = n
		}
	}

	if previos == current {
		return 0, 0, errors.New("Error! Input slice didn't include two different elements")
	}

	return previos, current, nil
}
