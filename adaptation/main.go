package main
import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"image/color"
)
const (
	screenWidth  = 640
	screenHeight = 480
)
var (
	sliderValue float64
)
func update(screen *ebiten.Image) error {
	// Получаем текущее значение ползунка
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, _ := ebiten.CursorPosition()
		sliderValue = float64(x) / screenWidth
	}
	// Отображаем текущее значение ползунка
	screen.Fill(color.White)
	ebitenutil.DrawLine(screen, screenWidth*sliderValue, 0, screenWidth*sliderValue, screenHeight, color.Black)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Slider Value: %.2f", sliderValue))
	return nil
}
func main() {
	ebiten.Run(update, screenWidth, screenHeight, 1, "Slider Example")
}