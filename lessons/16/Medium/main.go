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
	//Ініціалізація застосунку fyne
	myApp := app.New()

	//Створення вікна графічного інтерфейсу
	myWindow := myApp.NewWindow("Застосунок 2")

	//Створення першого поля вводу
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Введіть перше число")

	//Створення другого поля вводу
	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Введіть друге число")

	//Створення напису з результатом
	resultLabel := widget.NewLabel("Місце для результату")

	//Створення кнопки
	calcButton := widget.NewButton("Обчислити суму", func() {
		//Тіло функції-тригера
		num1Text := entry1.Text
		num2Text := entry2.Text

		num1, err1 := strconv.ParseFloat(num1Text, 64)
		num2, err2 := strconv.ParseFloat(num2Text, 64)

		if err1 != nil || err2 != nil {
			resultLabel.SetText(fmt.Sprintf("[%s] Помилка: введіть лише числа!", time.Now().Format("15:04:05")))
			return
		}

		sum := num1 + num2

		//Задання значення напису
		resultLabel.SetText(fmt.Sprintf("[%s] Сума: %.2f", time.Now().Format("15:04:05"), sum))
	})

	//Розміщення вмісту
	content := container.NewVBox(
		widget.NewLabel("Введіть два числа для обчислення суми:"),
		entry1,
		entry2,
		calcButton,
		resultLabel,
	)
	myWindow.SetContent(content)

	//Задання розміру вікна
	myWindow.Resize(fyne.NewSize(400, 300))

	//Запуск графічного інтерфейсу
	myWindow.ShowAndRun()
}
