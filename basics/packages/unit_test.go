package main

import "testing"

func Average(arr []float64) float64 {
	if len(arr) == 0 {
		return 0.0 // empty array handling
	}
	total := 0.0
	for _, element := range arr {
		total += element
	}
	return total / float64(len(arr))
}

// run with `go test unit_test.go`

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

// test pair
type AverageTestPair struct {
	values           []float64
	expected_average float64
}

var tests = []AverageTestPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverageBatched(t *testing.T) {
	for _, test_pair := range tests {
		avg := Average(test_pair.values)
		if avg != test_pair.expected_average {
			t.Errorf("Expected %f, but got %v", test_pair.expected_average, avg)
		}
	}
}
