package Toxel

import (
	_ "image/png"

	"github.com/faiface/pixel"
)

type hitbox struct {
	minX   float64
	minY   float64
	maxX   float64
	maxY   float64
	radius float64
}

var emptyHitbox = &hitbox{0, 0, 0, 0, 0}

//GameObject is gameobject with its sprite, type and hitbox
type GameObject struct {
	Active    bool
	Pos       pixel.Vec
	Hitbox    *hitbox
	Sprite    *pixel.Sprite
	Scale     float64
	Animation []Animator
	Type      int
}

var emptyGameObject = GameObject{true, pixel.ZV, emptyHitbox, nil, 1, []Animator{}, 0}

//convert pixel rectangle to toxel hitbox
func RectToHitbox(rect pixel.Rect, scale float64, pos pixel.Vec) *hitbox {
	return &hitbox{
		minX:   rect.Min.X - pos.X - (6 * scale),
		minY:   rect.Min.Y - pos.Y - (9 * scale),
		maxX:   (rect.Max.X - pos.X) * scale,
		maxY:   (rect.Max.Y - pos.Y) * scale,
		radius: 0,
	}
}

//changes gameObject scale without messing up the hitbox
func (g *GameObject) SetScale(scale float64) {
	g.Hitbox.maxX *= scale / g.Scale
	g.Hitbox.maxY *= scale / g.Scale
	g.Scale = scale
}

//fixes hitbox by removing negatives values from it.
// func (h *hitbox) fix() *hitbox {
// 	if h.minX < 0 {
// 		fmt.Println("b")
// 		h.maxX += h.minX * -1
// 		h.minX = 0
// 	}
// 	if h.minY < 0 {
// 		fmt.Println("a")
// 		h.maxY += h.minY * -1
// 		h.minY = 0
// 	}
// 	return h
// }

//hitboxesCollide Checks whenever any hitbox collide with any object
//It uses
func AnyHitboxesCollide(gameObjects []GameObject) [][]int {
	var colliders [][]int
	for i := 0; i != len(gameObjects); i++ {
		for j := 0; j != len(gameObjects); j++ {
			if i != j && HitboxCollides(gameObjects[i], gameObjects[j]) {
				colliders = append(colliders, []int{i, j})
			}
		}
	}
	return colliders
}

func (g *GameObject) Collide(gameObjects []GameObject) []int {
	var colliders []int
	for j := 0; j != len(gameObjects); j++ {
		if HitboxCollides(*g, gameObjects[j]) {
			colliders = append(colliders, j)
		}
	}
	return colliders
}

func HitboxCollides(a, b GameObject) bool {
	return InBetween(a.Pos.X, a.Hitbox.minX-b.Hitbox.minX, a.Hitbox.maxX+b.Hitbox.maxX, a.Hitbox.radius+b.Hitbox.radius) && //
		InBetween(a.Pos.Y, b.Hitbox.minY+a.Hitbox.minY, b.Hitbox.maxY+a.Hitbox.maxY, a.Hitbox.radius+b.Hitbox.radius) || //&&
		// inBetween(a.Pos.Y, b.Hitbox.minY+a.Hitbox.minY, b.Hitbox.maxY+a.Hitbox.maxY, b.Hitbox.radius) ||
		InRadius(a.Hitbox.radius, b.Hitbox.radius, a.Pos, b.Pos)
}

func InBetween(num, min, max, difference float64) bool {
	return num > min-difference && num < max+difference
}

func InRadius(aRadius, bRadius float64, aPos, bPos pixel.Vec) bool {
	radius := aRadius + bRadius
	return InBetween(aPos.Sub(bPos).X, radius*-1, radius, 0) &&
		InBetween(aPos.Sub(bPos).Y, radius*-1, radius, 0)
}
