package main
import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
const (
	cscreenWidth  = 600
	cscreenHeight = 600
)
type App struct {
	angle float64
	screenWidth int
	screenHeight int
}

func (a *App) Update() error {
	a.angle += 0.1
	return nil
}
func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	a.draw_grid(screen)

	centerX := a.screenWidth / 2
	centerY := a.screenHeight / 2
	lineLength := 50

	//fmt.Println(g.angle)

	endX1 := centerX + int(math.Cos(a.angle)*float64(lineLength))
	endY1 := centerY + int(math.Sin(a.angle)*float64(lineLength))

	endX2 := endX1 + int(math.Cos(a.angle/2)*float64(lineLength))
	endY2 := endY1 + int(math.Sin(a.angle/2)*float64(lineLength))

	endX3 := endX2 + int(math.Cos(a.angle/4)*float64(lineLength))
	endY3 := endY2 + int(math.Sin(a.angle/4)*float64(lineLength))

	ebitenutil.DrawLine(screen, float64(centerX), float64(centerY), float64(endX1), float64(endY1), color.Black)

	ebitenutil.DrawLine(screen, float64(endX1), float64(endY1), float64(endX2), float64(endY2), color.Black)

	ebitenutil.DrawLine(screen, float64(endX2), float64(endY2), float64(endX3), float64(endY3), color.Black)
}
func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.screenWidth, a.screenHeight
}
func main() {
	ebiten.SetWindowSize(cscreenWidth, cscreenHeight)
	ebiten.SetWindowTitle("Rotation Video")
	App := &App{0,cscreenWidth, cscreenHeight}

	RoboManip := NewManipulator(cscreenWidth,cscreenHeight,300,300,3,[]int{10,20,30},[]float64{0,1,2})
	fmt.Println(RoboManip.robotStruct)
	fmt.Println(RoboManip.angleList)

	if err := ebiten.RunGame(App); err != nil {
		log.Fatal(err)
	}
}