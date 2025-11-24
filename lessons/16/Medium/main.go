package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()

	myWindow := myApp.NewWindow("Застосунок 2")

	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Введіть перше число")

	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Введіть друге число")

	resultLabel := widget.NewLabel("Місце для результату")

	calcButton := widget.NewButton("Обчислити суму", func() {

		num1Text := entry1.Text
		num2Text := entry2.Text

		num1, err1 := strconv.ParseFloat(num1Text, 64)
		num2, err2 := strconv.ParseFloat(num2Text, 64)

		if err1 != nil || err2 != nil {
			resultLabel.SetText(fmt.Sprintf("[%s] Помилка: введіть лише числа!", time.Now().Format("15:04:05")))
			return
		}

		sum := num1 + num2

		resultLabel.SetText(fmt.Sprintf("[%s] Сума: %.2f", time.Now().Format("15:04:05"), sum))
	})

	content := container.NewVBox(
		widget.NewLabel("Введіть два числа для обчислення суми:"),
		entry1,
		entry2,
		calcButton,
		resultLabel,
	)

	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(400, 300))

	myWindow.ShowAndRun()
}
