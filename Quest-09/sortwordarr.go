package piscine

func SortWordArr(array []string) {
	counter := 0
	for range array {
		counter++
	}
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
