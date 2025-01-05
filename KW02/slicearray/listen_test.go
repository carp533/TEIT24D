package main

import (
	"reflect"
	"testing"
)

func TestListMax(t *testing.T) {
	tests := []struct {
		name        string
		list        []int
		expectedVal int
		expectedPos int
	}{
		{"Empty list", []int{}, 0, -1},
		{"Single element", []int{5}, 5, 0},
		{"Multiple elements", []int{1, 3, 5, 7, 9}, 9, 4},
		{"Negative values", []int{-1, -3, -5}, -1, 0},
		{"First element is max", []int{10, 3, 2}, 10, 0},
		{"Last element is max", []int{1, 3, 20}, 20, 2},
		{"Duplicates max value", []int{2, 5, 5, 1}, 5, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, pos := listMax(test.list)
			if val != test.expectedVal || pos != test.expectedPos {
				t.Errorf("%v: listMax(%v) = %v, %v; want %v, %v", test.name, test.list, val, pos, test.expectedVal, test.expectedPos)
			}
		})
	}
}
func TestDivisors(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{"Divisors of 10", 10, []int{1, 2, 5, 10}},
		{"Divisors of 60", 60, []int{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60}},
		{"Divisors of 37", 37, []int{1, 37}},
		{"Divisors of 1", 1, []int{1}},
		{"Divisors of 0", 0, []int{}}, // Annahme: Leere Liste für 0
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := divisors(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("divisors(%d) = %v, want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Prime number 2", 2, true},
		{"Prime number 3", 3, true},
		{"Non-prime number 4", 4, false},
		{"Prime number 5", 5, true},
		{"Non-prime number 10", 10, false},
		{"Non-prime number 0", 0, false},
		{"Non-prime number 1", 1, false},
		{"Negative number -7", -7, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPrime(test.input)
			if result != test.expected {
				t.Errorf("isPrime(%d) = %v, want %v", test.input, result, test.expected)
			}
		})
	}
}

func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(i)
	}
}
