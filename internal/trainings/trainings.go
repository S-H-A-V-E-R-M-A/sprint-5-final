package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// структура с данные о тренировке: количество шагов, тип тренировки,
// длительность тренировки, а также данные из структуры personaldata.Personal,
// то есть имя, вес и рост пользователя
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// теперь экспортируемая, пример получаемой строки: "3456,Ходьба,3h00m"
// парсинг строки и запись их в поля структуры Training
func (t *Training) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",") // разбил строку на слайс строк по запятой
	if len(values) != 3 {                    // если длинна слайса меньше 3 то нули, ошибка и пустая строка
		return errors.New("длинна слайса меньше 3")
	}

	counter, err := strconv.Atoi(values[0]) //конверчу первую строку из слайса
	if err != nil || counter <= 0 {         // если ошибка или ко-во шагов меньше или равно нулю то ошибка
		return errors.New("ошибка конвертации")
	}
	t.Steps = counter // пишу кол-во шагов в поле Steps структуры

	t.TrainingType = values[1] // пишу тип тренировки в поле TrainingType структуры

	duration, err := time.ParseDuration(values[2]) // конверчу строку третью в формат time
	if err != nil || duration <= 0 {               // если есть ошибка или время меньше или равно нулю то ошибка
		return errors.New("ошибка парсинга time")
	}
	t.Duration = duration // пишу время нагрузки в поле Duration структуры

	return nil
}

// метод для формирования строки с информацией о проведённой тренировке по ее типу
func (t Training) ActionInfo() (string, error) {
	distKm := spentenergy.Distance(t.Steps, t.Height)             // считаю дистанцию в км используя ф-цию из пакета spentenergy
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) // скорость

	var calories float64 // объявил переменную калории
	var err error        // объявил переменую для ошибок

	switch { // перебор типов нагрузки и подсчет по типу нагрузки калории
	case t.TrainingType == "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case t.TrainingType == "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = errors.New("неизвестный тип тренировки") //если непонятный тип нагрузки - то ошибка
	}

	if err != nil { // если ошибка то пустая строка и ошибка
		return "", err
	}

	// собираю в кучу результат и возвращаю его
	// МУСОР, УДАЛТЬ лол х2, тесты не принимают эту шляпу так как sprint как и print без f отбрасывает нули
	//res := fmt.Sprint("Тип тренировки: ", t.TrainingType, "\nДлительность: ", t.Duration.Hours(), " ч.\nДистанция: ", distKm, " км.\nСкорость: ", speed, " км/ч\nСожгли калорий: ", calories)
	res := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distKm,
		speed,
		calories,
	)

	return res, nil
}
