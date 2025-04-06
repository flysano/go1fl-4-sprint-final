package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go1fl-4-sprint-final/internal/spentcalories"
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) < 2 {
		return 0, 0, fmt.Errorf("invalid input format: expected 2 values, got %d: %s", len(dataSlice), data)
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("step count must be greater than zero")
	}

	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("failed to parse package: %v", err)
	}

	distance := float64(steps) * spentcalories.LenStep / spentcalories.MInKm
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.", steps, distance, calories)

}
