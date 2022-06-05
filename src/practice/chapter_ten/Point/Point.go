package Point

import "math"

type Point struct {
	X, Y float64
}

type Point3D struct {
	X, Y, Z float64
}

type Point_Polar struct {
	R, T float64
}

func Abs(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func Scale(s float64, p Point) (q Point) {
	q.X = p.X * s
	q.Y = p.Y * s
	return
}
