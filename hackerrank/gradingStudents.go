//input
//73
//67
//38
//33

//output
//75
//67
//40
//33

package main

func gradingStudents(grades []int32) []int32 {
	result := []int32{}
	for _, value := range grades {
		remain := value / 5
		nextValue := (remain + 1) * 5
		if nextValue-value < 3 && 40-value < 3 {
			result = append(result, nextValue)
		} else {
			result = append(result, value)
		}
	}
	return result
}
