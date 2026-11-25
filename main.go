package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ktttnv/twist-karaoke/internal/audio"
	"github.com/ktttnv/twist-karaoke/internal/game"
)

func main() {
	capture, err := audio.NewCapture()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Не удалось запустить микрофон:", err)
		os.Exit(1)
	}
	defer capture.Stop()

	g := game.New(capture)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Твист караоке - Индикатор громкости")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
