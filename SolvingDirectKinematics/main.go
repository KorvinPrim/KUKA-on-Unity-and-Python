package main
import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

)
const (
	cscreenWidth  = 600
	cscreenHeight = 600
)

var RoboManip = NewManipulator(
	float64(cscreenWidth),
	float64(cscreenHeight),
	300.0,
	300.0,
	3,
	[]float64{50,50,50}, 						//link lengths
	[]float64{
		math.Pi*1.5,
		math.Pi*1.5,
		math.Pi*1.5},   //starting position
	[]float64{
		math.Pi/128,
		math.Pi/256,
		math.Pi/128},					//rotation speed
	[]float64{0,0,0})							//acceleration of rotation

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

	RoboManip.calculate_movement()
	a.draw_manip(screen, RoboManip)
	a.draw_iterface(screen)



}
func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.screenWidth, a.screenHeight
}
func main() {
	ebiten.SetWindowSize(cscreenWidth, cscreenHeight)
	ebiten.SetWindowTitle("Robo - arm")
	App := &App{0,cscreenWidth, cscreenHeight}


	fmt.Println(RoboManip.robotStruct)


	if err := ebiten.RunGame(App); err != nil {
		log.Fatal(err)
	}
}