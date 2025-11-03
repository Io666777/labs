package main

import (
	"fmt"
	"sort"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Модель для БД
type Result struct {
	ID        uint   `gorm:"primaryKey"`
	Task      string // "sheep", "incrementer", "sort"
	Input     string 
	Output    string
	CreatedAt time.Time
}

var db *gorm.DB

func main() {
	// Подключаемся к БД
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	} else {
		db.AutoMigrate(&Result{})
		fmt.Println("OK!")
	}

	// 1. Подсчет овец
	sheeps := []bool{true, false, false, false, true, true, false, true, true, false}
	fmt.Println("овцы тут->", sheeps)
	counter := countSheeps(sheeps)
	fmt.Println("Количество овец:", counter)
	saveResult("sheep", fmt.Sprintf("%v", sheeps), fmt.Sprintf("%d", counter))

	// 2. Увеличение массива
	nums := []int{4, 7, 9}
	fmt.Println("массив до преобразований->", nums)
	result := incrementer(nums)
	fmt.Println("результат->", result)
	fmt.Println("оригинальный массив после->", nums)
	saveResult("incrementer", fmt.Sprintf("%v", nums), fmt.Sprintf("%v", result))

	// 3. Сортировка массива
	array := []int{12, 54, 11, 17, 1, 0, 3, 7, 4}
	fmt.Println("массив до тудасюда->", array)
	sorted := sortArray(array)
	fmt.Println("массив после тудасюда", sorted)
	saveResult("sort", fmt.Sprintf("%v", array), fmt.Sprintf("%v", sorted))
}

func saveResult(task, input, output string) {
	if db != nil {
		db.Create(&Result{
			Task:      task,
			Input:     input,
			Output:    output,
			CreatedAt: time.Now(),
		})
	}
}

// Ваши оригинальные функции
func incrementer(nums []int) []int {
	result := make([]int, len(nums))
	for i, value := range nums {
		sum := value + (i + 1)
		result[i] = sum % 10
	}
	return result
}

func countSheeps(sheeps []bool) int {
	counter := 0
	for _, value := range sheeps {
		if value {
			counter++
		}
	}
	return counter
}

func sortArray(array []int) []int {
	newArray := []int{}
	for _, value := range array {
		if value%2 != 0 {
			newArray = append(newArray, value)
		}
	}
	sort.Ints(newArray)
	obratka := 0
	for i := 0; i < len(array); i++ {
		if array[i]%2 != 0 {
			array[i] = newArray[obratka]
			obratka++
		}
	}
	return array
}