package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
	colorBeat    = 333
)

var (
	img          *ebiten.Image
	p            = newPalette(Colors)
	currentColor = p.random()
)

func init() {
	img, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
	img.Fill(currentColor)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	img.Fill(currentColor)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%2.f FPS", ebiten.CurrentFPS()))
	return nil
}

func main() {
	go func() {
		for range time.Tick(colorBeat * time.Millisecond) {
			currentColor = p.next()
		}
	}()

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Pouet pouet"); err != nil {
		log.Fatal(err)
	}
}
