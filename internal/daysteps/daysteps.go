package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// структура с необходимыми данными
// в которую встроена структура с персональными данными
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// метод парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps
func (ds *DaySteps) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",") // разбил строку на слайс строк по запятой
	if len(values) != 2 {                    // если меньше 2 строк в салйсе то ошибка
		return fmt.Errorf("неверный формат данных: %s", datastring)
	}

	counter, err := strconv.Atoi(values[0]) //конверчу первую строку из слайса
	if err != nil || counter <= 0 {         // если есть ошибка то нули и err
		return errors.New("ошибка парсинга")
	}
	ds.Steps = counter // заполняю поле структуры Steps

	duration, err := time.ParseDuration(values[1]) // конверчу строку вторую в формат time
	if err != nil || duration <= 0 {               // если есть ошибка то нули и err
		return errors.New("ошибка конвертации")
	}
	ds.Duration = duration // заполняю поле структуры Duration

	return nil
}

// метод для формирования строки с информацией о прогулках
func (ds DaySteps) ActionInfo() (string, error) {
	// получаю дистанцию в ки из ф-ции из пакета spentenergy
	distKm := spentenergy.Distance(ds.Steps, ds.Height)

	// получаю калории из ф-ции из пакета spentenergy
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil || calories <= 0 { // если ошибка или каллорий меньшу или равно нулю то ошибка и пустая строка
		fmt.Println("ошибка ", err)
		return "", err
	}

	// результат
	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distKm, calories)

	return res, nil
}
