package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/olzhas/lab4-variant1/pkg/tilecalc"
)

func main() {
	// Генерация уникального ID (внешний пакет uuid)
	sessionID := uuid.New().String()

	color.Green("=== Лабораторная работа №4 (Вариант 1) ===")
	color.Cyan("Идентификатор сессии: %s\n", sessionID)

	roomLength := 5.0
	roomWidth := 4.0
	tileLength := 0.3
	tileWidth := 0.3
	reserve := 10.0

	// Вызов функции F1 с обработкой ошибок
	roomArea, err := tilecalc.CalcArea(roomLength, roomWidth)
	if err != nil {
		color.Red("Ошибка расчета площади: %v", err)
		return
	}

	tileArea, err := tilecalc.CalcArea(tileLength, tileWidth)
	if err != nil {
		color.Red("Ошибка расчета плитки: %v", err)
		return
	}

	// Демонстрация работы функции F2 с указателем
	mutableArea := roomArea
	err = tilecalc.ApplyReserveWithPointer(&mutableArea, 5.0)
	if err != nil {
		color.Red("Ошибка указателя: %v", err)
		return
	}

	// Расчет общего количества плиток
	tilesNeeded, err := tilecalc.CalcTilesNeeded(roomArea, tileArea, reserve)
	if err != nil {
		color.Red("Ошибка расчета количества: %v", err)
		return
	}

	// Вызов функции F3 (формирование отчета)
	report := tilecalc.GenerateReport("Гостиная", roomArea, tileArea, tilesNeeded)

	color.Yellow("\nРезультаты расчетов:")
	fmt.Println(report)
	fmt.Printf("Площадь после применения указателя (F2) с запасом 5%%: %.3f кв.м\n", mutableArea)

	// Использование функции из импортированного чужого готового пакета (Пункт 10)
	formattedCount := humanize.Comma(int64(tilesNeeded))
	color.Magenta("Демонстрация стороннего пакета go-humanize: отформатированное число плиток -> %s шт.", formattedCount)
}
