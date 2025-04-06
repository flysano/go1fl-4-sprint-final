package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	LenStep = 0.65 // средняя длина шага.
	MInKm   = 1000 // количество метров в километре.
	MinInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) < 3 {
		return 0, "", fmt.Errorf("invalid input format: expected 3 values, got %d: %s", len(dataSlice), data)
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("step count must be greater than zero")
	}

	activity := dataSlice[1]

	duration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return 0, 0, err
	}

	return steps, activity, duration, nil

}


// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	fullDistance := (float64(steps) * LenStep) / MInKm

	return fullDistance //расстояние в км
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return distance(steps) / duration.Hours() //средняя скорость км/ч

}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(activity string, durationHours, distance, speed, calories float64) {
	fmt.Printf("Тип тренировки: %s\nДлительность: %.1f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity, durationHours, distance, speed, calories)

}

// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return err.Error()
	}

	var calories float64, data string

	durationHours := duration.Hours()
	distance := distance(steps)
	speed := meanSpeed(steps, duration)

	switch activity {
	case "Бег":
		calories = RunningSpentCalories(steps, weight, duration)
		data = ShowTrainingInfo(activity, durationHours, distance, speed, calories)
	case "Хотьба":
		calories = WalkingSpentCalories(steps, weight, height, duration)
		data = ShowTrainingInfo(activity, durationHours, distance, speed, calories)
	default:
		data = "Неизвестный тип тренировки"
	}
	return data
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {

	speed := meanSpeed(steps, duration)
	calories := ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight

	return calories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	speed := meanSpeed(steps, duration)
	durationHours := duration.Hours()
	calories := ((walkingCaloriesWeightMultiplier * weight) + ((speed * speed / height) * walkingSpeedHeightMultiplier)) * durationHours * minInH

	return calories

}
