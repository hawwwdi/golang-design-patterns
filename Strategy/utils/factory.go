package utils

import (
	"fmt"
	"github.com/hawwwdi/golang-design-patterns/Strategy/models"
	"github.com/hawwwdi/golang-design-patterns/Strategy/shapes"
)

const (
	CIRCLE = "circle"
	SQUARE = "square"
)

func MakeShape(name string, size int) (models.Shape, error){
	switch name {
	case CIRCLE:
		return shapes.NewCircle(size), nil
	case SQUARE:
		return shapes.NewSquare(size), nil
	default:
		return nil, fmt.Errorf("shape %v not found!", name)
	}
}
