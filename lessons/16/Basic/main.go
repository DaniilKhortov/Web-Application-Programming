package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()

	myWindow := myApp.NewWindow("Застосунок 1")

	button := widget.NewButton("Натисни мене!", func() {
		fmt.Println("Кнопку натиснуто!")
	})

	centered := container.NewCenter(button)

	myWindow.SetContent(centered)

	myWindow.Resize(fyne.NewSize(300, 200))

	myWindow.ShowAndRun()
}
