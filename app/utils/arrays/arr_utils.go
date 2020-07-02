package arrays

// IndexOfArr is a helper function to find the index of an array element
func IndexOfArr(element string, data []string) int {
	for key, value := range data {
		if element == value {
			return key
		}
	}
	return -1 //not found.
}

// RemoveElementFromStringArr removes an element from a string array
func RemoveElementFromStringArr(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
