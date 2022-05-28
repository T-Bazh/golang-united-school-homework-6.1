package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return fmt.Errorf("this shape goes out of the shapesCapacity range")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) || i < 0 {
		return nil, fmt.Errorf("index is out of range. Index is %d with the list length %d", i, len(b.shapes))
	}
	shape := b.shapes[i]
	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) || i < 0 {
		return nil, fmt.Errorf("index is out of range. Index is %d with the list length %d", i, len(b.shapes))
	}
	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) || i < 0 {
		return nil, fmt.Errorf("index is out of range. Index is %d with the list length %d", i, len(b.shapes))
	}
	replacedShape := b.shapes[i]
	b.shapes[i] = shape
	return replacedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	total := 0.0
	for _, shape := range b.shapes {
		total += shape.CalcPerimeter()
	}
	return total
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	total := 0.0
	for _, shape := range b.shapes {
		total += shape.CalcArea()
	}
	return total
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var circleFound bool
	shapes := []Shape{}
	for _, shape := range b.shapes {
		if shape.GetType() != "circle" {
			shapes = append(shapes, shape)
			circleFound = true
		}
	}
	b.shapes = shapes
	if !circleFound {
		return fmt.Errorf("circles are not exist in the list")
	}
	return nil
}
