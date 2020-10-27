package game

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/algosup/game/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Surface struct {
	*ebiten.Image
}

type Bitmap struct {
	*ebiten.Image
}

func Run(title string, width int, height int, draw func(surface Surface)) {
	ebiten.Run(func(surface *ebiten.Image) error {

		if draw != nil {
			draw(Surface{surface})
		}
		return nil
	}, width, height, 1.0, title)

}

func DrawRect(surface Surface, x int, y int, width int, height int, color color.Color) {
	ebitenutil.DrawRect(surface.Image, float64(x), float64(y), float64(width), float64(height), color)
}

func DrawText(surface Surface, text string, x int, y int) {
	ebitenutil.DebugPrintAt(surface.Image, text, x, y)
}

func LoadBitmap(name string) (Bitmap, error) {
	ei, _, e := ebitenutil.NewImageFromFile(name, ebiten.FilterDefault)
	return Bitmap{ei}, e
}

func DrawBitmap(surface Surface, x int, y int, bitmap Bitmap) {
	surface.Image.DrawImage(bitmap.Image, &ebiten.DrawImageOptions{GeoM: ebiten.TranslateGeo(float64(x), float64(y))})
}
