package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	// Подключаемся к Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	// Проверяем подключение
	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Println("Ошибка Redis:", err)
		return
	}

	// Тестируем функции
	sheeps := []bool{true, false, false, true}
	countSheeps(sheeps)

	nums := []int{4, 7, 9}
	incrementer(nums)

	array := []int{12, 54, 11, 17, 1, 0, 3, 7, 4}
	sortArray(array)
}

func countSheeps(sheeps []bool) int {
	// Новый формат ключа: task1_метка_времени
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("task1_%d", timestamp)
	
	// Пробуем взять из Redis
	if count, err := rdb.Get(ctx, key).Int(); err == nil {
		return count
	}
	// Вычисляем если нет в кэше
	counter := 0
	for _, value := range sheeps {
		if value {
			counter++
		}
	}
	// Сохраняем в Redis
	rdb.Set(ctx, key, counter, time.Hour)
	return counter
}

func incrementer(nums []int) []int {
	// Новый формат ключа: task2_метка_времени
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("task2_%d", timestamp)
	
	// Пробуем взять из Redis
	if result, err := rdb.Get(ctx, key).Bytes(); err == nil {
		var data []int
		json.Unmarshal(result, &data)
		return data
	}
	// Вычисляем если нет в кэше
	result := make([]int, len(nums))
	for i, value := range nums {
		sum := value + (i + 1)
		result[i] = sum % 10
	}
	// Сохраняем в Redis
	jsonData, _ := json.Marshal(result)
	rdb.Set(ctx, key, jsonData, time.Hour)
	return result
}

func sortArray(array []int) []int {
	// Новый формат ключа: task3_метка_времени
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("task3_%d", timestamp)
	
	// Пробуем взять из Redis
	if result, err := rdb.Get(ctx, key).Bytes(); err == nil {
		var data []int
		json.Unmarshal(result, &data)
		return data
	}
	// Вычисляем если нет в кэше
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
	// Сохраняем в Redis
	jsonData, _ := json.Marshal(array)
	rdb.Set(ctx, key, jsonData, time.Hour)
	return array
}