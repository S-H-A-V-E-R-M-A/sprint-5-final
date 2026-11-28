package actioninfo

import (
	"fmt"
	"log"
)

// интерфейс в который входят два метода из пакетов trainings и daysteps
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// функция в которую можно передавать структуры из пакетов trainings и daysteps
func Info(dataset []string, dp DataParser) {
	for _, znachenie := range dataset { // перебор всех значений слайса
		err := dp.Parse(znachenie) // прогоняю из через метод для парсинга строк, вынимаю ошибку или nil
		if err != nil {
			log.Println(err) // если есть ошибка то логгирую ее
		}
		res, err := dp.ActionInfo() // через метод для получения строки с данными о прогулке получаю строку с данными и ошибку/nil
		if err != nil {
			log.Println(err) // если есть ошибка то логгирую ее
		}
		fmt.Println(res) // вывел строку с данными о прогулке
	}
}
