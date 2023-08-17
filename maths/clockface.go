package maths

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	return math.Pi
}

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
