package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ktttnv/twist-karaoke/internal/audio"
)

func main() {
	capture, err := audio.NewCapture()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Не удалось запустить микрофон:", err)
		os.Exit(1)
	}
	defer capture.Stop()

	fmt.Println("🎤 Микрофон запущен. Говорите/пойте! Ctrl+C для выхода.")

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	for {
		select {
		case <-ticker.C:
			fmt.Printf("\rГромкость (RMS): %.4f", capture.RMS())
		case <-stop:
			fmt.Println("\nВыход.")
			return
		}
	}
}
