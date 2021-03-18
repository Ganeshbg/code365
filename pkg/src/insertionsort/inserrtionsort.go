package insertionsort

func InsertionSort(list []int) []int{
	for i := 1; i< len(list); i++ {
		j := i-1
		temp := list[i]
		for j>=0 && list[j] > temp {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = temp
	}
	return list
}
