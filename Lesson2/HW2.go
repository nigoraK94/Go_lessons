package main

import "fmt"

func main() {
var a, b float64
    var op string

    fmt.Print("Введите первое число, оператор (+, -, *, /) и второе число: ")
    fmt.Scan(&a, &op, &b)

    switch op {
    case "+":
        fmt.Printf("Результат: %.2f\n", a+b)
    case "-":
        fmt.Printf("Результат: %.2f\n", a-b)
    case "*":
        fmt.Printf("Результат: %.2f\n", a*b)
    case "/":
        if b == 0 {
            fmt.Println("Ошибка: деление на ноль!")
        } else {
            fmt.Printf("Результат: %.2f\n", a/b)
        }
    default:
        fmt.Println("Неизвестная операция")
    }
}
