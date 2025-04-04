package spentcalories

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if dataSlice < 3 {
		return 0, "", fmt.Errorf("Неверный формат ввода данных: %s", data)
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше нуля")
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
	fullDistance := (float64(steps) * lenStep) / float64(mInKm)

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
	
	averageSpeed := distance(steps) / duration.Hours()

	return averageSpeed //средняя скорость км/ч
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	steps, activity, duration, err := parseTraining(data)
		if err != nil {
			return ""
		}
		durationHours := duration.Hours()
		distance := distance(steps)
		speed := meanSpeed(steps, duration)
		caloriesWalk := WalkingSpentCalories(steps, weight, height, duration)
		caloriesRun := RunningSpentCalories(steps, weight, duration)
	switch activity {
	case "Бег":
		fmt.Printf("Тип тренировки: %s\n
		Длительность: %.1f ч.\n
		Дистанция: %.2f км.\n
		Скорость: %.2f км/ч\n
		Сожгли калорий: %.2f",
		activity, durationHours, distance, speed, calcaloriesRun)

	case "Хотьба":
		fmt.Printf("Тип тренировки: %s\n
		Длительность: %.1f ч.\n
		Дистанция: %.2f км.\n
		Скорость: %.2f км/ч\n
		Сожгли калорий: %.2f",
		activity, durationHours, distance, speed, caloriesWalk)
	default:
		fmt.Println("неизвестный тип тренировки")
	}
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
	calories := ((runningCaloriesMeanSpeedMultiplier*speed)-runningCaloriesMeanSpeedShift) * weight
	
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



