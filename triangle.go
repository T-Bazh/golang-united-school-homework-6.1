package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (triangle *Triangle) GetType() string {
	return "triangle"
}
func (triangle *Triangle) CalcPerimeter() float64 {
	return 3 * triangle.Side
}
func (triangle *Triangle) CalcArea() float64 {
	return triangle.Side * math.Sqrt(float64(3)) / 4
}
