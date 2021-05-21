package v2

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/algosup/game/v2/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var surface *ebiten.Image

type Bitmap struct {
	*ebiten.Image
}

func Run(title string, width int, height int, draw func()) {
	ebiten.Run(func(s *ebiten.Image) error {
		surface = s
		if draw != nil {
			draw()
		}
		return nil
	}, width, height, 1.0, title)

}

func DrawRect(x int, y int, width int, height int, color color.Color) {
	ebitenutil.DrawRect(surface, float64(x), float64(y), float64(width), float64(height), color)
}

func DrawText(text string, x int, y int) {
	ebitenutil.DebugPrintAt(surface, text, x, y)
}

func LoadBitmap(name string) (Bitmap, error) {
	ei, _, e := ebitenutil.NewImageFromFile(name, ebiten.FilterDefault)
	return Bitmap{ei}, e
}

func DrawBitmap(x int, y int, bitmap Bitmap) {
	surface.DrawImage(bitmap.Image, &ebiten.DrawImageOptions{GeoM: ebiten.TranslateGeo(float64(x), float64(y))})
}
