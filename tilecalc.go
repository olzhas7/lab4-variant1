// Package tilecalc предназначен для расчетов количества строительной плитки,
// площади помещений и формирования отчетов.
package tilecalc

import (
	"fmt"
	"math"
)

// CalcArea вычисляет площадь помещения по длине и ширине.
// Возвращает ошибку, если входные данные меньше или равны нулю.
func CalcArea(length, width float64) (float64, error) {
	if length <= 0 || width <= 0 {
		return 0, fmt.Errorf("ошибка: длина и ширина должны быть больше нуля (получено: %.2f, %.2f)", length, width)
	}
	return length * width, nil
}

// CalcTilesNeeded рассчитывает количество плиток с учетом запаса в процентах (F1).
func CalcTilesNeeded(roomArea, tileArea, reservePercent float64) (int, error) {
	if tileArea <= 0 {
		return 0, fmt.Errorf("ошибка: площадь плитки не может быть равна нулю или меньше")
	}
	if roomArea <= 0 {
		return 0, fmt.Errorf("ошибка: площадь помещения не может быть отрицательной или нулевой")
	}
	if reservePercent < 0 {
		return 0, fmt.Errorf("ошибка: процент запаса не может быть отрицательным")
	}

	baseCount := roomArea / tileArea
	reserveMultiplier := 1.0 + (reservePercent / 100.0)
	totalWithReserve := baseCount * reserveMultiplier
	return int(math.Ceil(totalWithReserve)), nil
}

// ApplyReserveWithPointer изменяет значение площади с учетом дополнительного коэффициента через указатель (F2).
func ApplyReserveWithPointer(area *float64, extraPercent float64) error {
	if area == nil {
		return fmt.Errorf("ошибка: передан нулевой указатель")
	}
	if extraPercent < 0 {
		return fmt.Errorf("ошибка: процент не может быть отрицательным")
	}
	*area = *area * (1.0 + extraPercent/100.0)
	return nil
}

// GenerateReport формирует строку текстового отчета с использованием fmt.Sprintf (F3).
func GenerateReport(roomName string, roomArea, tileArea float64, tilesCount int) string {
	return fmt.Sprintf("Отчет для помещения: %s\nПлощадь: %.2f кв.м\nНеобходимое количество плитки (с запасом): %d шт.", roomName, roomArea, tilesCount)
}
