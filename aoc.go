// Package aoc contains various utility methods useful throughout Advent of Code problems
package aoc

import "math"

// Bearing returns the degrees clockwise from the +Y axis to a given point
func Bearing(x, y int) float64 {
	return math.Mod(360-math.Atan2(float64(-x), float64(-y))*(180/math.Pi), 360)
}

// Max returns the maximum value of two given integers
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the minimum value of two given integers
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Abs returns the absolute value of a given integer
func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
