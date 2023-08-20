package maths

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * math.Pi / 30 // 2π / 60
	return math.Pi / (30 / float64(t.Second())) // For floats horrible
}

func secondHandPoint(t time.Time) Point {
	return Point{0, -1}
}

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
