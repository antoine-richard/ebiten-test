package main

import (
	"image/color"
	"math/rand"
)

// from https://github.com/peterhellberg/pixel-experiments/blob/master/bouncing/bouncing-gradient.go
type Palette struct {
	colors []color.RGBA
	size   int
	index  int
}

func newPalette(cc []color.Color) *Palette {
	colors := []color.RGBA{}

	for _, v := range cc {
		if c, ok := v.(color.RGBA); ok {
			colors = append(colors, c)
		}
	}

	return &Palette{colors, len(colors), 0}
}

func (p *Palette) clone() *Palette {
	return &Palette{p.colors, p.size, p.index}
}

func (p *Palette) next() color.RGBA {
	if p.index++; p.index >= p.size {
		p.index = 0
	}

	return p.colors[p.index]
}

func (p *Palette) color() color.RGBA {
	return p.colors[p.index]
}

func (p *Palette) random() color.RGBA {
	p.index = rand.Intn(p.size)

	return p.colors[p.index]
}

var Colors = []color.Color{
	color.RGBA{190, 38, 51, 255},
	color.RGBA{224, 111, 139, 255},
	color.RGBA{73, 60, 43, 255},
	color.RGBA{164, 100, 34, 255},
	color.RGBA{235, 137, 49, 255},
	color.RGBA{247, 226, 107, 255},
	color.RGBA{47, 72, 78, 255},
	color.RGBA{68, 137, 26, 255},
	color.RGBA{163, 206, 39, 255},
	color.RGBA{0, 87, 132, 255},
	color.RGBA{49, 162, 242, 255},
	color.RGBA{178, 220, 239, 255},
}
