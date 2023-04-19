package helpers

import "strconv"

// StringToInt convert string to int
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
