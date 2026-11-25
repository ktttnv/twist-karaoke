package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SongSelectScreen struct {
	game  *Game
	back  MenuButton
	songs []MenuButton
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
		songs: []MenuButton{
			{
				label: "Песня 1",
				rect:  image.Rect(200, 200, 600, 260),
				action: func() Screen {
					return NewGameScreen(g, "Песня 1")
				},
			},
			{
				label: "Песня 2",
				rect:  image.Rect(200, 280, 600, 340),
				action: func() Screen {
					return NewGameScreen(g, "Песня 2")
				},
			},
		},
	}
}

func (s *SongSelectScreen) Update() (Screen, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		pt := image.Pt(mx, my)

		if pt.In(s.back.rect) {
			return s.back.action(), nil
		}

		for _, song := range s.songs {
			if pt.In(song.rect) {
				return song.action(), nil
			}
		}
	}

	return nil, nil
}

func (s *SongSelectScreen) Draw(screen *ebiten.Image) {
	drawText(screen, "Выбор песни", 300, 80, 24, color.White)

	for _, song := range s.songs {
		drawButton(screen, song)
	}

	drawButton(screen, s.back)
}
