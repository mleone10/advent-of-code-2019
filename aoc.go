// Package aoc contains various utility methods useful throughout Advent of Code problems
package aoc

import "math"

// Bearing returns the degrees clockwise from the +Y axis to a given point
func Bearing(x, y int) float64 {
	return math.Mod(360-math.Atan2(float64(-x), float64(-y))*(180/math.Pi), 360)
}
