package Toxel

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

type Button struct {
	gameObject GameObject
	text       string
	atlas      *text.Atlas
	onPress    func()
}

//TODO: implement function for checking if user pressed button.

//newButton creates new Button with simple hitboxes
func newButton(buttonText string, position pixel.Vec, scale float64, atlas *text.Atlas, onPress func()) *Button {
	button := &Button{
		gameObject: emptyGameObject,
		text:       buttonText,
		atlas:      atlas,
		onPress:    onPress,
	}
	button.gameObject.Scale = scale
	button.gameObject.Pos = position
	textProperties := text.New(button.gameObject.Pos, button.atlas)
	fmt.Fprintln(textProperties, button.text)
	unscaledHitbox := textProperties.Bounds()
	button.gameObject.Hitbox = rectToHitbox(unscaledHitbox, scale, button.gameObject.Pos)
	return button
}

//setScale sets button scale without messing up the hitboxes (can be buggy)
func (b *Button) setScale(scale float64) {
	b.gameObject.Hitbox.minX += 6*scale - (6 * b.gameObject.Scale)
	b.gameObject.Hitbox.minY += 9*scale - (9 * b.gameObject.Scale)
	b.gameObject.Hitbox.maxY *= scale / b.gameObject.Scale
	b.gameObject.Hitbox.maxX *= scale / b.gameObject.Scale
}

//TODO: support matrix
//Draw draw button in given window with given text color
func (button *Button) Draw(win pixel.Target, color color.Color) {
	textProperties := text.New(button.gameObject.Pos, button.atlas)
	textProperties.Color = color
	fmt.Fprintln(textProperties, button.text)
	textProperties.Draw(win, pixel.IM.Scaled(textProperties.Orig, button.gameObject.Scale))
	textProperties.Clear()
}
