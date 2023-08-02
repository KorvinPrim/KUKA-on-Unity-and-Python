package main
import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
const (
	screenWidth  = 640
	screenHeight = 480
)
type Game struct {
	angle float64
}
func (g *Game) Update() error {
	g.angle += 0.1
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	centerX := screenWidth / 2
	centerY := screenHeight / 2
	lineLength := 50
	endX1 := centerX + int(math.Cos(g.angle)*float64(lineLength))
	endY1 := centerY + int(math.Sin(g.angle)*float64(lineLength))

	endX2 := endX1 + int(math.Cos(g.angle/2)*float64(lineLength))
	endY2 := endY1 + int(math.Sin(g.angle/2)*float64(lineLength))

	endX3 := endX2 + int(math.Cos(g.angle/4)*float64(lineLength))
	endY3 := endY2 + int(math.Sin(g.angle/4)*float64(lineLength))

	ebitenutil.DrawLine(screen, float64(centerX), float64(centerY), float64(endX1), float64(endY1), color.Black)

	ebitenutil.DrawLine(screen, float64(endX1), float64(endY1), float64(endX2), float64(endY2), color.Black)

	ebitenutil.DrawLine(screen, float64(endX2), float64(endY2), float64(endX3), float64(endY3), color.Black)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rotation Video")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}