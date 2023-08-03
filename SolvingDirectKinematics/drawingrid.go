package main

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (a *App) draw_grid(screen *ebiten.Image){
	var w  = float64(a.screenWidth)
	var h  = float64(a.screenHeight)

	ebitenutil.DrawRect(screen, 0, 0, w, h, color.White)

	// Рисуем сетку (32x32 и 64x64)
	gridColor64 := &color.RGBA{A: 50}
	gridColor32 := &color.RGBA{A: 20}
	for y := 0.0; y < h; y += 25 {
		ebitenutil.DrawLine(screen, 0, y, w, y, gridColor32)
	}
	for y := 0.0; y < h; y += 50 {
		ebitenutil.DrawLine(screen, 0, y, w, y, gridColor64)
	}
	for x := 0.0; x < w; x += 25 {
		ebitenutil.DrawLine(screen, x, 0, x, h, gridColor32)
	}
	for x := 0.0; x < w; x += 50 {
		ebitenutil.DrawLine(screen, x, 0, x, h, gridColor64)
	}
}
