package main

import (
	"github.com/blusewang/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	const S = 400
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	text := "Hello, world!"
	w, h := dc.MeasureString(text)

	dc.Fill()
	dc.Push()
	dc.Rotate(gg.Radians(10))
	dc.DrawStringAnchored(text, 100, 180, 0.0, 0.0)
	dc.DrawRectangle(100, 180, w, h)
	dc.Stroke()
	dc.Pop()
	dc.SavePNG("out.png")
}
