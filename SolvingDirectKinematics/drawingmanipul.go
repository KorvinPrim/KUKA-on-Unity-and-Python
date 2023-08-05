package main

import "image/color"

/*
The code is a part of a larger program and is
responsible for drawing a manipulator (robot arm) and its
working area on the screen.

In summary, the code iterates over the line segments
of a robot arm and draws them on the screen using the
"ebitenutil.DrawLine" function. It then draws the working
area of the manipulator as a circle using the "drawCircle" function.
*/

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
