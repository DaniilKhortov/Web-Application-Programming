// Для допоміжних файлів необхідно вказати такий самий пакет, що й основна програма.
package main

// Підключення пакету math (mathematic).
//math відповідає за здійснення складних арифметчних операцій.
import (
	"math"
)

// Функція calculateEnergy. Імпортує 5 значень у форматі float64 та повертає число у float64
// float64 - це число з плавуючою комою
func calculateEnergy(tariff, kDay, energyDay, kNight, energyNight float64) float64 {
	//Повернення результатів обчислення.
	return tariff*kDay*math.Round(energyDay) + tariff*kNight*math.Round(energyNight)
}
