package rpn

import (
    "fmt"
    "net/http"
    "sync"
    // "time"
    // "strconv"
)

var (
    fibs = []int{0, 1} // Начальные числа Фибоначчи
    mu sync.Mutex    // Мьютекс для безопасного доступа к fibs
    current int          // Индекс текущего возвращаемого числа
    requestCount int
)

// fibonacciHandler обрабатывает запросы и возвращает следующее число Фибоначчи
func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock() // Блокируем доступ к текущему индексу
    defer mu.Unlock()

    // Проверяем, нужно ли вычислить следующее число
    if current >= len(fibs) {
        next := fibs[current-1] + fibs[current-2]
        fibs = append(fibs, next)
    }

    // Возвращаем текущее число и увеличиваем индекс
    fmt.Fprint(w, fibs[current])
    current++
}


func Metrics(next http.HandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        mu.Lock() // Блокируем доступ к текущему индексу
        requestCount++
        mu.Unlock()

        next(w, r)
    })
}


func MetricsHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock() // Блокируем доступ к текущему индексу
    defer mu.Unlock()

    // Возвращаем текущее число
    fmt.Fprintf(w, "rpc_duration_milliseconds_count %v", requestCount)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", FibonacciHandler)
    mux.Handle("/metrics", Metrics(MetricsHandler)) // Связываем путь с обработчиком
    fmt.Println("Сервер работает на http://localhost:8080/")
    http.ListenAndServe(":8080", mux) // Запуск сервера
}   