package functions

func FindIndex(str string, strArr []string) int {
	for index, element := range strArr {
		if str == element {
			return index
		}
	}
	return -1
}
