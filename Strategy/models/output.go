package models

import (
	"io"
)

type Shape interface {
	Draw() error
	SetWriter(io.Writer)
	SetLogger(io.Writer)
}

type Output struct {
	Writer io.Writer
	Logger io.Writer
}

func (output *Output) SetWriter(writer io.Writer) {
	output.Writer = writer
}

func (output *Output) SetLogger(logger io.Writer) {
	output.Logger = logger
}
