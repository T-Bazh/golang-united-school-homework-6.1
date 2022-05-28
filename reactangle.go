package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rect *Rectangle) GetType() string {
	return "rectangle"
}
func (rect *Rectangle) CalcPerimeter() float64 {
	return (rect.Weight + rect.Height) * 2
}
func (rect *Rectangle) CalcArea() float64 {
	return rect.Weight * rect.Height
}
