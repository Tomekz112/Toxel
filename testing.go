package Toxel

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var buttonA, buttonB, buttonC, buttonD *Button

var Frames float64

func main() {
	Frames = 0
	buttonA = newButton("1", pixel.V(0, -100), 2, basicAtlas, test)
	buttonD = newButton("1", pixel.V(12, -100), 2, basicAtlas, test)
	buttonB = newButton("1", pixel.V(24, -100), 2, basicAtlas, test)
	buttonC = newButton("1", pixel.V(-1500, -100), 2, basicAtlas, test)

	fmt.Println(buttonA.gameObject.Hitbox.minX)

	fmt.Println(buttonB.gameObject.Hitbox.maxY)
	pixelgl.Run(game)
}

func masno(FPS float64) {
	fmt.Println("witam witam tu fps masne som ogolnie takie jak:", FPS)
}

func game() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	e := 40.0
	s := 40.0
	a := 1.0

	FPSCounter(&Frames, masno)
	for !win.Closed() {
		win.SetMatrix(pixel.IM.Moved(pixel.V(e, s)).ScaledXY(pixel.ZV, pixel.V(a, a)))
		if win.Pressed(pixelgl.KeyLeft) {
			e++
		} else if win.Pressed(pixelgl.KeyRight) {
			e--
		}
		if win.Pressed(pixelgl.KeyUp) {
			s--
		} else if win.Pressed(pixelgl.KeyDown) {
			s++
		}

		if win.JustPressed(pixelgl.KeyEqual) {
			a *= 1.25
			fmt.Println("current scale: ", a)
		} else if win.JustPressed(pixelgl.KeyMinus) {
			a /= 1.25
			fmt.Println("current scale: ", a)
		}

		if win.JustPressed(pixelgl.Key0) {
			buttonA.text = "1"
		}
		if win.JustPressed(pixelgl.Key1) {
			buttonA.text = " "
		}
		win.Clear(colornames.Skyblue)
		buttonD.Draw(win, colornames.Salmon)
		buttonA.Draw(win, colornames.Salmon)
		buttonB.Draw(win, colornames.Salmon)
		buttonC.Draw(win, colornames.Salmon)
		win.Update()
		Frames++
	}
}

func test() {

}
