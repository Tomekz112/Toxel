package Toxel

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Animator struct {
	FPS            uint64
	AnimationSheet []*pixel.Sprite
}

func (g *GameObject) StartAnimation(animTime int, scale float64, ID uint, position pixel.Vec, win *pixelgl.Window) {
	frameDur := float64(time.Second) / float64(g.Animation[ID].FPS) //Counts time of 1 frame
	idleAnimation := g.Sprite
	go func(g *GameObject, frameDur float64, ID uint, animTime int, scale float64, position pixel.Vec, win *pixelgl.Window) {
		timeStart := time.Now() //beggining of the animation time
		lastFrame := timeStart  //time of last frame
		index := 0
		for int(time.Now().Sub(timeStart)) < animTime {
			if float64(time.Now().Sub(lastFrame)) <= frameDur {
				g.Sprite = g.Animation[ID].AnimationSheet[SaveSliceAccess(index, len(g.Animation[ID].AnimationSheet))]
				//pixel.Draw(a.AnimationSheet[SaveSliceAccess(index,len(a.AnimationSheet))])
				index++
				lastFrame = time.Now()
			}
		}
	}(g, frameDur, ID, animTime, scale, position, win)
	g.Sprite = idleAnimation
}
