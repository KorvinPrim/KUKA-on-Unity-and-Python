package main

/*
The code is a part of a larger program and it defines
functions related to drawing text on the screen using the Ebiten library in Go.

The code begins by importing necessary packages and
libraries. It then declares two font variables -
mplusNormalFont and mplusBigFont.

The init() function is called during the initialization
phase of the program. It parses a TrueType font file
(MPlus1pRegular_ttf) using the opentype library and
assigns it to the mplusNormalFont variable. It also
sets the size, DPI, and hinting options for the font.
Then, it adjusts the line height for the mplusBigFont.

The draw_interface() function takes in some data and
a screen image as parameters. It uses the text.Draw()
function from the Ebiten library to draw text on the
screen using the mplusNormalFont. It formats the text
based on the provided data and draws it at specific
coordinates on the screen.

The repeatingKeyPressed() function checks if a key is
pressed repeatedly. It takes in a key parameter and
returns a boolean value indicating whether the key is
being pressed repeatedly based on a delay and interval.

The draw_text() function is not fully implemented in
the provided code. It is supposed to draw text on the
screen, including a blinking cursor. It uses the
mplusNormalFont and the text.Draw() function to
draw the text at specific coordinates on the screen.

Overall, the code initializes fonts, provides functions
to draw text on the screen using those fonts, and checks
for repeated key presses.
*/

import (
	"fmt"
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Adjust the line height.
	mplusBigFont = text.FaceWithLineHeight(mplusBigFont, 54)
}

func (a *App) draw_iterface(
	directdata map[string]float64,
	workarea map[string]float64,
	screen *ebiten.Image,
) {

	const x = 20
	DirecktText := fmt.Sprintf("Solving a direct kinematics problem - X:%.1f Y:%.1f Q:%.1f",
		directdata["xHand"],
		directdata["yHand"],
		directdata["QHand"])
	text.Draw(screen, DirecktText, mplusNormalFont, x, 50, color.Black)
	WorkingRangeText := fmt.Sprintf(
		"The range of operation of the robo-hand - \nDistance:%.1f Xmin:%.1f Xmax:%.1f Ymin:%.1f Ymax:%.1f",
		workarea["Distance"],
		workarea["Xmin"],
		workarea["Xmax"],
		workarea["Ymin"],
		workarea["Ymax"],
	)
	text.Draw(screen, WorkingRangeText, mplusNormalFont, x, 80, color.Black)
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (a *App) draw_text(screen *ebiten.Image) {
	// Blink the cursor.
	t := a.text
	if a.counter%60 < 30 {
		t += "_"
	}
	const x = 20

	//text.Draw(screen, t, mplusNormalFont, x, 200, color.Black)

	//ebitenutil.DebugPrint(screen, t)
}
