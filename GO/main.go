package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var choice int
    fmt.Println("Выберите задание:")
    fmt.Println("1 - Буферизированный vs небуферизированный канал")
    fmt.Println("2 - Вычисление факториала через горутины")
    fmt.Print("Ваш выбор: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        demoChannels()
    case 2:
        calcFactorial()
    default:
        fmt.Println("Неверный выбор")
    }
}

func demoChannels() {
    fmt.Println("\n=== Небуферизированный канал ===")
    unbuffered := make(chan int)

    go func() {
        fmt.Println("Горутина: отправка 1...")
        unbuffered <- 1
        fmt.Println("Горутина: отправка завершена")
    }()

    time.Sleep(2 * time.Second)
    fmt.Println("Основная программа: получение...")
    fmt.Println("Получено:", <-unbuffered)
    time.Sleep(500 * time.Millisecond)

    fmt.Println("\n=== Буферизированный канал ===")
    buffered := make(chan int, 2)

    buffered <- 10
    buffered <- 20
    fmt.Println("Отправлено 10 и 20 без блокировки")

    fmt.Println("Получено:", <-buffered)
    fmt.Println("Получено:", <-buffered)
}

func calcFactorial() {
    n := 10
    parts := 3
    results := make(chan int, parts)

    var wg sync.WaitGroup

    for i := 0; i < parts; i++ {
        wg.Add(1)
        start := i*n/parts + 1
        end := (i + 1) * n / parts
        if i == parts-1 {
            end = n
        }

        go func(s, e int) {
            defer wg.Done()
            product := 1
            for j := s; j <= e; j++ {
                product *= j
            }
            fmt.Printf("Горутина %d-%d: %d\n", s, e, product)
            results <- product
        }(start, end)
    }

    wg.Wait()
    close(results)

    final := 1
    for res := range results {
        final *= res
    }

    fmt.Printf("\nФакториал %d = %d\n", n, final)
}