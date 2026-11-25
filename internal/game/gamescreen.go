package game

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameScreen struct {
	game      *Game
	songLabel string
	back      MenuButton
}

func NewGameScreen(g *Game, songLabel string) *GameScreen {
	return &GameScreen{
		game:      g,
		songLabel: songLabel,
		back: MenuButton{
			label: "Назад",
			rect:  image.Rect(50, 520, 200, 570),
			action: func() Screen {
				return NewSongSelectScreen(g)
			},
		},
	}
}

func (gs *GameScreen) Update() (Screen, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if image.Pt(mx, my).In(gs.back.rect) {
			return gs.back.action(), nil
		}
	}

	return nil, nil
}

func (gs *GameScreen) Draw(screen *ebiten.Image) {
	drawText(screen, fmt.Sprintf("Игра: %s", gs.songLabel), 300, 200, 24, color.White)
	drawButton(screen, gs.back)
}
