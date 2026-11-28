package personaldata

import "fmt"

// структура с персональными данными: имя, вес и рост пользователя
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// метод для вывода данных из структуры Personal
func (p Personal) Print() {
	// тесты не принимают так как print без f отбрасывает нули
	//fmt.Print("Имя: ", p.Name, "\nВес: ", p.Weight, "\nРост: ", p.Height, "\n")
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
