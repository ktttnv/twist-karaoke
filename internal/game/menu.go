package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuScreen struct {
	game    *Game
	buttons []MenuButton
}

type MenuButton struct {
	label  string
	rect   image.Rectangle
	action func() Screen
}

func NewMenuScreen(g *Game) *MenuScreen {
	return &MenuScreen{
		game: g,
		buttons: []MenuButton{
			{
				label: "Начать играть",
				rect:  image.Rect(300, 200, 500, 250),
				action: func() Screen {
					return NewSongSelectScreen(g)
				},
			},
			{
				label: "Индикатор громкости",
				rect:  image.Rect(300, 270, 500, 320),
				action: func() Screen {
					return NewVolumeIndicatorScreen(g)
				},
			},
		},
	}
}

func (m *MenuScreen) Update() (Screen, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		for _, b := range m.buttons {
			if image.Pt(mx, my).In(b.rect) {
				return b.action(), nil
			}
		}
	}

	return nil, nil
}

func (m *MenuScreen) Draw(screen *ebiten.Image) {
	for _, b := range m.buttons {
		drawButton(screen, b)
	}
}
