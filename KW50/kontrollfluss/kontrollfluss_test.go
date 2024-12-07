// Aufgaben zu Funktionen, Schleifen und If-Then-Else
package main

import (
	"fmt"
	"math"
	"testing"
)

func Examplesum() {
	//Output:
	//15
	fmt.Println(sum(5))
}
func Exampledouble() {
	//Output:
	//10
	//30
	fmt.Println(double(5))
	fmt.Println(double(15))
}

func TestDouble(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{0, 0},
		{1, 2},
		{-2, -4},
		{-1, -2},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := double(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, 1},
		{2, 3},
		{5, 15},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sum(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSumMultiples(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{3, 10, 18},
		{2, 10, 20},
		{10, 2, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := sumMultiples(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{31, true},
		{101, true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isPrime(tt.a)
			if ans != tt.want {

				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{3, 5, 15},
		{2, 5, 10},
		{4, 10, 20},
		{20, 5, 20},
		{25, 10, 50},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := lcm(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestBinrep(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"Binary of 0", 0, "0"},
		{"Binary of 1", 1, "1"},
		{"Binary of 2", 2, "10"},
		{"Binary of 5", 5, "101"},
		{"Binary of 10", 10, "1010"},
		{"Binary of 16", 16, "10000"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := binrep(test.input)
			if result != test.expected {
				t.Errorf("binrep(%d) = %s, want %s", test.input, result, test.expected)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	// Test cases als Array von anonymen Structs
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1, 1},
		{4, 2},
		{16, 4},
		{2, 1.4142135623730951},
		{3, 1.7320508075688772},
		{25, 5},
		{0.25, 0.5},
	}

	for _, tc := range testCases {
		result := sqrt(tc.input)
		// Vergleiche mit einer Toleranz von 1e-10
		if math.Abs(result-tc.expected) > 1e-10 {
			t.Errorf("sqrt(%f) = %f; expected %f",
				tc.input, result, tc.expected)
		}
	}

	// Test für negative Zahlen
	result := sqrt(-1)
	if !math.IsNaN(result) {
		t.Errorf("sqrt(-1) = %f; expected NaN", result)
	}
}
