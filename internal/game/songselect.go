package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SongSelectScreen struct {
	game *Game
	back MenuButton
}

func NewSongSelectScreen(g *Game) *SongSelectScreen {
	return &SongSelectScreen{
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

func (s *SongSelectScreen) Update() (Screen, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if image.Pt(mx, my).In(s.back.rect) {
			return s.back.action(), nil
		}
	}

	return nil, nil
}

func (s *SongSelectScreen) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Выбор песни", 250, 200)
	drawButton(screen, s.back)
}
