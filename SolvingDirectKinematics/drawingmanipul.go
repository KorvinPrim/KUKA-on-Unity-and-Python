package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"fmt"
	"image/color"
)

func (a *App) draw_manip(screen *ebiten.Image, robo *Manipulator) {
	for _, value := range robo.robotStruct {
		//fmt.Println(index,value)
		ebitenutil.DrawLine(screen,
			value["beginX"],
			value["beginY"],
			value["endX"],
			value["endY"],
			color.Black)
	}
	//draw a working area
	purpleClr := color.RGBA{255, 0, 0, 255}
	a.drawCircle(screen, int(robo.centerX), int(robo.centerX), int(robo.working_area["Distance"]), purpleClr)
}
