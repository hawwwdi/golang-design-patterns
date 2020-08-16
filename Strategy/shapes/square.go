package shapes

import (
	"fmt"
	"github.com/hawwwdi/golang-design-patterns/Strategy/models"
	"os"
)

type Square struct {
	models.Output
	Side int
}

func (square *Square) Draw() error{
	_, err := fmt.Fprintf(square.Writer, "Square Side=%v", square.Side)
	if err != nil {
		_, _ = fmt.Fprint(square.Logger, err)
		return err
	}
	return nil
}

func NewSquare(side int) *Square {
	return &Square{
		Output: models.Output{
			Writer: os.Stdout,
			Logger: os.Stdout,
		},
		Side: side,
	}
}
