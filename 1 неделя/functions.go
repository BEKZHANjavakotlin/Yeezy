package main

import (
	"errors"
	"math"
	"strconv"
)

func add(a, b int) int {
	return a + b
}
func greet(name string) string {
	return "Hello " + name + "!"
}
func isAdult(age int) bool {
	if age >= 18 {
		return true
	} else {
		return false
	}
}
func divide(a, b float64) (float64, error) {
	if b == 0. {
		return 0., errors.New("Cannot divide by zero")
	} else {
		return a / b, nil
	}
}
func minMax(nums []int) (min, max int) {
	for _, n := range nums {
		min = nums[0]
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return min, max
}
func circleStats(radius float64) (area, perimeter float64) {
	area = radius * math.Pi * radius
	perimeter = 2 * math.Pi * radius
	return area, perimeter
}
func parsenAndDouble(s string) (int, error) {
	convert, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return convert * 2, nil
}
