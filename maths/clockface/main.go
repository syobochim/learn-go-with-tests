package main

import (
	"fmt"
	"learn-go/maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	maths.SVGWriter(os.Stdout, t)
}

func secondHandTag(p maths.Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
