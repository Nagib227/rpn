package application

import (
	// "bufio"
	// "log"
	// "os"
	// "strings"

	"fmt"
	"github.com/Nagib227/rpn/pkg/rpn"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

// Функция запуска приложения
// тут будем читать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() {
	err := rpn.Run()
	if err != nil {
        fmt.Println("Ошибка:", err)
    }
}