package key

import "github.com/hajimehoshi/ebiten"

var Up = ebiten.KeyUp
var Enter = ebiten.KeyEnter
var Right = ebiten.KeyRight
var Left = ebiten.KeyLeft
var Down = ebiten.KeyDown
var Space = ebiten.KeySpace

func IsPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}
