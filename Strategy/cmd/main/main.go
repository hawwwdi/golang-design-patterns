package main

import (
	"flag"
	"fmt"
	"github.com/hawwwdi/golang-design-patterns/Strategy/utils"
	"log"
)

func main() {
	in := flag.String("shape", "circle", "use for shape name")
	size := flag.Int("size", 4, "use for shape size")
	flag.Parse()
	fmt.Println(*in , *size)
	shape, err:= utils.MakeShape(*in, *size)
	handleError(err)
	err  = shape.Draw()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
