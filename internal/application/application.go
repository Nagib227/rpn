package application

import (
	// "bufio"
	"log"
	// "os"
	// "strings"

	"github.com/Nagib_227/rpn/pkg/rpn"
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
	log.Println(rpn.NewWorld(20, 30))
}