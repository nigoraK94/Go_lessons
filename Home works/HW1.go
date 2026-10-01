package main

import "fmt"

// func main() {

// Задача 1.
// var (
// 	i    int
// 	i8   int8
// 	i16  int16
// 	i32  int32
// 	i64  int64
// 	u    uint
// 	u8   uint8
// 	u16  uint16
// 	u32  uint32
// 	u64  uint64
// 	f32  float32
// 	f64  float64
// 	b    bool
// 	r    rune
// 	by   byte
// 	s    string
// 	c64  complex64
// 	c128 complex128
// )

// fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n", i, i8, i16, i32, i64)
// fmt.Printf("uint: %d, uint8: %d, uint16: %d, uint32: %d, uint64: %d\n", u, u8, u16, u32, u64)
// fmt.Printf("float32: %f, float64: %f\n", f32, f64)
// fmt.Printf("bool: %t\n", b)
// fmt.Printf("rune: %d (символ: %c), byte: %d\n", r, r, by)
// fmt.Printf("string: %q\n", s)
// fmt.Printf("complex64: %v, complex128: %v\n\n", c64, c128)

// i = 5
// i8 = 10
// i16 = 20
// i32 = 30
// i64 = 40
// u = 50
// u8 = 60
// u16 = 70
// u32 = 80
// u64 = 90
// f32 = 1.5
// f64 = 5.5
// b = true
// r = 'М'
// by = 'B'
// s = "GO!"
// c64 = 1 + 1i
// c128 = 2 + 2i

// fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n", i, i8, i16, i32, i64)
// fmt.Printf("uint: %d, uint8: %d, uint16: %d, uint32: %d, uint64: %d\n", u, u8, u16, u32, u64)
// fmt.Printf("float32: %f, float64: %f\n", f32, f64)
// fmt.Printf("bool: %t\n", b)
// fmt.Printf("rune: %c, byte: %c\n", r, by)
// fmt.Printf("string: %s\n", s)
// fmt.Printf("complex64: %v, complex128: %v\n", c64, c128)

// Задача2.
// var a int = 45
// var b int = 23

// fmt.Println("Сложение:", a+b)
// fmt.Println("Вычитание:", a-b)
// fmt.Println("Умножение:", a*b)
// fmt.Println("Деление:", a/b)
// fmt.Println("Остаток от деления:", a%b)

// Задача3

// const(
// 	PI = 3.14159
// 	CompanyName = "Alif"
// 	CurrentYear = 2026
// 	DaysInWeek = 7
// )
//     fmt.Println("Значение PI:", PI)
// 	fmt.Println("Название компании:", CompanyName)
// 	fmt.Println("Текущий год:", CurrentYear)
// 	fmt.Println("Дней в неделе:", DaysInWeek)
//

// Задача4

// var name string
// 	var age int
// 	var height float64
// 	var weight float64
// 	var isEmployed bool

// 	fmt.Print("Введите ваше имя: ")
// 	fmt.Scan(&name)

// 	fmt.Print("Введите ваш возраст: ")
// 	fmt.Scan(&age)

// 	fmt.Print("Введите ваш рост (в см): ")
// 	fmt.Scan(&height)

// 	fmt.Print("Введите ваш вес (в кг): ")
// 	fmt.Scan(&weight)

// 	fmt.Print("Вы работаете? (введите true или false): ")
// 	fmt.Scan(&isEmployed)

// 	fmt.Println("==============================")
// 	fmt.Printf("Имя:         %s\n", name)
// 	fmt.Printf("Возраст:     %d лет\n", age)
// 	fmt.Printf("Рост:        %.1f см\n", height)
// 	fmt.Printf("Вес:         %.1f кг\n", weight)
// 	fmt.Printf("Работает:    %t\n", isEmployed)
// 	fmt.Println("==============================")

// Задача5

// var num1 int
// var num2 int

// fmt.Print("Введите цнлое число:")
// fmt.Scan(&num1)

// fmt.Print("Введите второе целое число:")
// fmt.Scan(&num2)

// fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
// fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
// fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
// fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
// fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)

// Задача 6
// var firstName string = "Нигора"
// var lastName string = "Разработчик"
// var favoriteLanguage string = "Go"

// var age int = 32
// var experienceYears int = 1
// var projectsCount uint = 12

// var salary float64 = 5500.50

// var hasCommercialExperience bool = true
// var knowsDocker bool = true

// var nameFirstLetter rune = 'Н'
// var favoriteAscii byte = '$'

// fmt.Printf("Имя и Фамилия:         %s %s\n", firstName, lastName)
// fmt.Printf("Первая буква имени:    %c\n", nameFirstLetter) //
// fmt.Printf("Возраст:               %d лет\n", age)
// fmt.Printf("Стаж работы:           %d г.\n", experienceYears)
// fmt.Printf("Завершенных проектов:  %d\n", projectsCount)
// fmt.Printf("Оклад ($):             %.2f\n", salary)
// fmt.Printf("Любимый язык:          %s\n", favoriteLanguage)
// fmt.Printf("Опыт коммерческий:     %t\n", hasCommercialExperience)
// fmt.Printf("Знание Docker:         %t\n", knowsDocker)
// fmt.Printf("Любимый ASCII-символ:  %c (код: %d)\n", favoriteAscii, favoriteAscii)

// Задача8

const (
	Version = "v1.27"
	Border  = "----------------------------------------"
)

func main() {

	var name string
	var age int
	var salary float64
	var isReady bool

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("Желаемая зарплата ($): ")
	fmt.Scan(&salary)

	fmt.Print("Готовы к переезду? (true/false): ")
	fmt.Scan(&isReady)

	fmt.Println("\n" + Border)
	fmt.Printf("      ВИЗИТКА РАЗРАБОТЧИКА (%s)\n", Version)
	fmt.Println(Border)

	fmt.Printf("Имя:                   %s\n", name)
	fmt.Printf("Возраст:               %d лет\n", age)
	fmt.Printf("Желаемый оклад:        %.2f $\n", salary)
	fmt.Printf("Готовность к переезду:  %t\n", isReady)

	fmt.Println(Border)
}
