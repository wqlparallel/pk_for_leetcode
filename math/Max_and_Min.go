package math


func Min(numbers ...int)int{
	min := numbers[0]
	for _, number := range numbers{
		if number < min{
			min = number
		}
	}
	return min
}

func Max(numbers ...int)int{
	max := numbers[0]
	for _, number := range numbers{
		if number > max{
			max = number
		}
	}
	return max
}
