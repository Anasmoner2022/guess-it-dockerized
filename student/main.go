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
	windowSize := 20 // Smaller window for faster adaptation

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}

		// PREDICT FIRST (before adding current value)
		if len(window) < 2 {
			// Not enough data for good prediction
			fmt.Println("0 300")
		} else {
			m := CalculateMean(window)
			s := CalculateStdDev(window, m)

			// Ensure minimum stddev to avoid zero ranges
			if s < 10 {
				s = 10
			}

			k := 2.5 // Increased from 2.0 for better coverage
			lower := m - (k * s)
			upper := m + (k * s)

			fmt.Printf("%d %d\n", int(math.Round(lower)), int(math.Round(upper)))
		}

		// THEN ADD to window (for next prediction)
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
