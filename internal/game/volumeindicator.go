package game

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type VolumeIndicatorScreen struct {
	game *Game
	back MenuButton
}

func NewVolumeIndicatorScreen(g *Game) *VolumeIndicatorScreen {
	return &VolumeIndicatorScreen{
		game: g,
		back: MenuButton{
			label: "Назад",
			rect:  image.Rect(50, 520, 200, 570),
			action: func() Screen {
				return NewMenuScreen(g)
			},
		},
	}
}

func (v *VolumeIndicatorScreen) Update() (Screen, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if image.Pt(mx, my).In(v.back.rect) {
			return v.back.action(), nil
		}
	}

	return nil, nil
}

func (v *VolumeIndicatorScreen) Draw(screen *ebiten.Image) {
	rawVol := v.game.capture.RMS()
	smoothVol := v.game.capture.SmoothedRMS()

	drawVolumeBar(screen, 260, 100, 80, 450, float64(rawVol), color.RGBA{255, 100, 100, 255})
	drawVolumeBar(screen, 460, 100, 80, 450, float64(smoothVol), color.RGBA{100, 255, 100, 255})

	drawText(screen, fmt.Sprintf("Volume raw: %.4f", rawVol), 260, 80, 14, color.White)
	drawText(screen, fmt.Sprintf("Smooth raw: %.4f", smoothVol), 460, 80, 14, color.White)

	drawButton(screen, v.back)
}
