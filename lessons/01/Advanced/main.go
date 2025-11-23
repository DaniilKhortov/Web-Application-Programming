package main

// Підключення пакетів fmt (format), os (operating system), strconv (string conversion) та time.
// fmt використовується для форматованого вводу/виводу.
// time — для роботи з часом.
// os надає доступ до функцій операційної системи, зокрема до аргументів командного рядка.
// strconv використовується для конвертації між рядками та числовими типами.
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Точка входу програми.
func main() {
	// Тут виконується основний код програми.
	// Отримання поточного часу з системи пристрою.
	currentTime := time.Now()

	// Перевірка на кількість введених аргументів.
	// Має бути рівно 6 аргументів (включно з іменем програми).
	if len(os.Args) != 6 {
		// Обробка виключення.
		fmt.Printf("Cannot process! Number of arguments are not equal to 5!")

		// Зупинка роботи програми
		// Як і будь-яка функція, при повернені значення (return), функція main завершує роботую.
		return
	}

	// Конвертація тарифу з тексту у число.
	// У Go багато функцій повертають помилку, яку потрібно обробляти через змінну err.
	tariff, err := strconv.ParseFloat(os.Args[1], 64)

	// Перевірка чи отримало err значення обробленої помилки
	// Для позначення відсутності значення використовується nil
	if err != nil {
		fmt.Println("Invalid tariff:", err)
		return
	}

	// Конвертація денного коефіцієнту у текст.
	kDay, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid k of daytime:", err)
		return
	}

	// Конвертація використаної електрики за день у текст.
	energyDay, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid used energy amount during daytime:", err)
		return
	}

	// Конвертація нічного коефіцієнту у текст.
	kNight, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Println("Invalid k of nighttime:", err)
		return
	}

	// Конвертація використаної електрики за ніч у текст.
	energyNight, err := strconv.ParseFloat(os.Args[5], 64)
	if err != nil {
		fmt.Println("Invalid used energy amount during nighttime:", err)
		return
	}

	// Виклик функція calculateEnergy.
	// Його результат записується у змінну n.
	//calculateEnergy імплементовано у файлі Calc.go
	n := calculateEnergy(tariff, kDay, energyDay, kNight, energyNight)

	//Вивід результатів.
	fmt.Printf("Welcome to our app!\n")

	// Форматований вивід результатів у консоль:
	// %s — рядок, %.2f — число з плаваючою комою з двома знаками після неї.
	fmt.Printf("Today is %s\n", currentTime.Format("02.01.2006"))
	fmt.Printf("Power service bill: %.2f UAH\n", n)

}
