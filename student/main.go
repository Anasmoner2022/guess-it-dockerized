package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var window []float64
	windowSize := 20
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}

		if len(window) < 2 {
			fmt.Println("0 200")
		} else {
			m := CalculateMean(window)
			s := CalculateStdDev(window, m)
			cv := CV(m, s)
			if cv <= 10.0 {
				k := 1.0
				lower := m - (k * s)
				upper := m + (k * s)
			} else if cv > 10.0 && cv < 20.0 {
				k := 2.0
				lower := m - (k * s)
				upper := m + (k * s)
			} else {
				k := 3.0
				lower := m - (k * s)
				upper := m + (k * s)
			}

			fmt.Printf("%d %d\n", int(math.Round(lower)), int(math.Round(upper)))
		}
		window = append(window, value)
		if len(window) > windowSize {
			window = window[1:]
		}
	}
}

func CalculateMean(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

func CalculateVariance(data []float64, m float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	total := 0.0
	for _, num := range data {
		diff := m - num
		total += diff * diff
	}

	// Use n-1 for sample variance (more accurate)
	n := len(data)
	if n <= 1 {
		return 0.0
	}

	return total / float64(n-1)
}

func CalculateStdDev(data []float64, m float64) float64 {
	return math.Sqrt(CalculateVariance(data, m))
}

func CV(mean float64, std float64) float64 {
	if mean == 0 {
		return 0.0
	}
	return (std / mean) * 100
}
