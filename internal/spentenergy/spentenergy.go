package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// по росту/весу, кол-ву шагов и времени бега считаю сколько клорий сжег при ходьбе
// оптимизировал через обращение к RunningSpentCalories
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	caloriesRun, err := RunningSpentCalories(steps, weight, height, duration) // получаю калогии бега и ошибки если есть                                          // получаю скорость
	calories := caloriesRun * walkingCaloriesCoefficient                      // получаю калории для шага

	return calories, err
}

// по росту/весу, кол-ву шагов и времени бега считаю сколько клорий сжег при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { // если где-то ноль то ошибка
		return 0, errors.New("кривые параметры")
	}

	speed := MeanSpeed(steps, height, duration)                // получаю скорость
	calories := (weight * speed * duration.Minutes()) / minInH //получаю калории для бега

	return calories, nil
}

// теперь экспортируемая
// по росту, кол-ву шагов и времени ходьбы считаю среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// требование TODO из тз учел в прошлом спритне
	if steps <= 0 || height <= 0 || duration <= 0 { //если шаги\рост\время меньше или равно нулю то перну ноль
		return 0
	}

	dist := Distance(steps, height) // считаю дистанцию
	return dist / duration.Hours()  // возвращаю скорость
}

// теперь экспортируемая
// по росту и кол-ву шагов считаю пройденнуб дистанцию в км
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 { // если шагов или рост меньше или равен нулю то верну 0
		return 0
	}

	lengthStep := height * stepLengthCoefficient // считаю длинну шага
	return (lengthStep * float64(steps)) / mInKm // возвращаю пройденную дистанцибю в км
}
