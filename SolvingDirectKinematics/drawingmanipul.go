package main

import (
	//"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (a *App) draw_manip(screen *ebiten.Image, robo *Manipulator){
	for _, value := range robo.robotStruct{
		//fmt.Println(index,value)
		ebitenutil.DrawLine(screen,
			value["beginX"],
			value["beginY"],
			value["endX"],
			value["endY"],
			color.Black)
	}
}
