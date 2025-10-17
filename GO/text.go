package main

import ("fmt" 
		"sort")



func main() {
	var larva int
	fmt.Println("1-овечковый проверятор наличия оных в заведении вар2")
	fmt.Println("2-своеобразная прогрессия чисел вар4")
	fmt.Println("3-сортировка нечетных возрастаний var3")
	fmt.Scanln(&larva)
	switch (larva) {
	case 1:
		sheeps := []bool{true, false, false, false, true, true, false, true, true, false}
		fmt.Println("овцы тут->",sheeps)
		counter := countSheeps(sheeps)
		fmt.Println("Количество овец:", counter)
		
	case 2:
		nums := []int{4, 7, 9}
		fmt.Println("массив до преобразований->",nums)
		fmt.Println(incrementer(nums))
		fmt.Println("массив после преобразований->",nums)

	case 3:
		array := []int{12, 54, 11, 17, 1, 0, 3, 7, 4}
		fmt.Println("массив до тудасюда->",array)
		fmt.Println(array)
		sortirovachka := sortArray(array)
		fmt.Println("массив после тудасюда", sortirovachka)
	default:
		fmt.Println("ну и куда мы лезем?")
	}


}



func incrementer(nums []int) []int{
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

func sortArray(array []int)[]int{
	newArray := []int {}
	for _, value := range array{
		if value %2!=0 {
			newArray= append(newArray, value)
		}
	}
	sort.Ints(newArray)
	obratka :=0
	for i := 0; i < len(array); i++ {
		if array[i]%2!=0 {
			array[i]= newArray[obratka]
			obratka++			
		}
	}
	return array
}