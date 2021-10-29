package Toxel

import (
	"fmt"
	"image/color"

	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type ToxelText struct {
	pos   pixel.Vec
	scale float64
	// gameObject GameObject
	text  string
	atlas *text.Atlas
}

var basicColor = pixel.Alpha(1)
var basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
var emptyButton = Button{emptyGameObject, "", basicAtlas, nil}

//Creates new text
func newText(Text string, position pixel.Vec, scale float64, atlas *text.Atlas) *ToxelText {
	newText := &ToxelText{
		pos:   position,
		scale: scale,
		atlas: atlas,
	}
	return newText
}

//autoLineBreak creates linebreak if text crosses given point
//Be aware that autoLineBreak delets all old endlines from text
//Warning: splits words in dumb way
func (t *ToxelText) autoLineBreak(maxX float64) {
	t.removeEndLines()
	var finalText = t.text
	textlngth := float64(len(t.text)) * 6 * t.scale //gets text length by mulptiplying characters count by size of 1 character * scale
	if textlngth > maxX {                           //if textlength is crossing given point
		maxChars := maxX / 6 * t.scale //get max number of characters in 1 line
		splittedText := strings.Split(t.text, "")
		for i := 0; i < len(splittedText)/int(maxChars); i++ { //repeat till there are less characters than maximum in 1 line
			splittedText[int(maxChars)*i] += "\n"
		}
		finalText = strings.Join(splittedText, "")
	}
	t.text = finalText
}

func (t *ToxelText) removeEndLines() {
	t.text = strings.ReplaceAll(t.text, "\n", "")
}

//Draws text in given window
//TODO: support matrix
func (t *ToxelText) Draw(win pixel.Target, color color.Color) {
	textProperties := text.New(t.pos, t.atlas)
	textProperties.Color = color
	fmt.Fprintln(textProperties, t.text)
	textProperties.Draw(win, pixel.IM.Scaled(textProperties.Orig, t.scale))
	textProperties.Clear()
}
