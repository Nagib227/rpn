package rpn

import (
	"fmt"
	// "github.com/Nagib227/rpn/config"
	"os"
	"strconv"
)

type World struct {
	Height int 
    Width int 
    Cells [][]bool 
}

 func NewWorld(height, width int) (*World, error){
 	if height <= 0 || width <= 0 {
 		return nil, fmt.Errorf("err")
 	}
 	var c [][]bool
 	for range height {
 		c = append(c, make([]bool, width))
 	}
 	return &World{Height: height, Width: width, Cells: c}, nil
 }

func run() error {
    // Проверяем количество аргументов
    if len(os.Args) < 4 {
        return fmt.Errorf("недостаточно аргументов: требуется 3 (имя файла, ширина, высота, процент)")
    }

    // Извлекаем ширину, высоту и процент заполнения
    width, _ := strconv.Atoi(os.Args[1])
    height, _ := strconv.Atoi(os.Args[2])
    // percent := os.Args[3]

    // Формируем строку для записи в файл
    config, err := NewWorld(height, width)
    if err != nil {
        return fmt.Errorf("ошибка: %v", err)
    }

    var c string
    for i := range config.Cells {
    	for j := range i {
    		fmt.Errorf("%v", j)
    		// c += strconv.FormatBool(j) + " "
    	}
    	c += "\n"
    }
    

    // Создаем (или открываем) файл для записи
    file, err := os.Create("config.txt")
    if err != nil {
        return fmt.Errorf("ошибка создания файла: %v", err)
    }
    defer file.Close() // Отложенное закрытие файла

    // Записываем данные в файл
    _, err = file.WriteString(c)
    if err != nil {
        return fmt.Errorf("ошибка записи в файл: %v", err)
    }

    return nil
}
