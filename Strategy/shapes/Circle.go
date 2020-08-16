package shapes

import (
	"fmt"
	"github.com/hawwwdi/golang-design-patterns/Strategy/models"
	"os"
)

type Circle struct {
	models.Output
	R int
}

func (circle *Circle) Draw() error{
	_, err := fmt.Fprintf(circle.Writer, "Circle R=%v", circle.R)
	if err != nil {
		_, _ = fmt.Fprint(circle.Logger, err)
		return err
	}
	return nil
}

func NewCircle(r int) *Circle {
	return &Circle{
		Output: models.Output{
			Writer: os.Stdout,
			Logger: os.Stdout,
		},
		R: r,
	}
}
