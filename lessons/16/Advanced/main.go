package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()

	myWindow := myApp.NewWindow("Застосунок 3")

	timeLabel := widget.NewLabel("Час оновлюється...")

	titleLabel := widget.NewLabel("Поточний час системи:")

	content := container.NewVBox(
		titleLabel,
		timeLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.Show()

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {

			currentTime := t.Format("15:04:05")

			timeLabel.SetText(fmt.Sprintf(" %s", currentTime))
		}
	}()

	myApp.Run()
}
