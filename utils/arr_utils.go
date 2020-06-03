package utils

// helper function to find the index of an array element
func IndexOfArr(element string, data []string) int {
	for key, value := range data {
		if element == value {
			return key
		}
	}
	return -1 //not found.
}
