package shape

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
)

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw() {
	spew.Dump("Draw Circle")
}

type Square struct {
}

func (s *Square) Draw() {
	spew.Dump("Draw Square")
}

type Rec struct {
}

func (r *Rec) Draw() {
	spew.Dump("Draw Rec")
}

func GetShape(shape string) (Shape, error) {
	if shape == "Circle" {
		return &Circle{}, nil
	} else if shape == "Square" {
		return &Square{}, nil
	} else if shape == "Rec" {
		return &Rec{}, nil
	}
	return nil, errors.New("Shape not found")
}
