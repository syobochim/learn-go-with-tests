package maths

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * math.Pi / 30 // 2π / 60
	return math.Pi / (30 / float64(t.Second())) // For floats horrible
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// SecondHand は t.Time でのアナログ時計の秒針の単位ベクトルをPointとして表現したものです。
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // Scale it to the length of the hand
	p = Point{p.X, -p.Y}                                      // Flip it over the X axis to account for the SVG having an origin in the top left hand corner
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // Translate it to the right position (so that it's coming from an origin of (150,150))
	return p
}
