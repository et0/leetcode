package l_e_0168_string

func convertToTitle(columnNumber int) string {
	columnName := make([]byte, 0, 10)

	for columnNumber > 0 {
		// fmt.Println(columnNumber)
		number := columnNumber / 26

		// columnNumber < 26
		if number == 0 {
			columnName = append(columnName, byte(64+columnNumber%26))
			break
		}

		// columnNumber >= 26
		if columnNumber-number*26 == 0 {
			columnName = append(columnName, byte(64+26))
			columnNumber = number - 1
		} else {
			columnName = append(columnName, byte(64+columnNumber-number*26))
			columnNumber = number
		}
	}
	// fmt.Println(columnName)

	// Reverse
	size := len(columnName)
	for i := 0; i < size/2; i++ {
		columnName[i], columnName[size-1-i] = columnName[size-1-i], columnName[i]
	}

	return string(columnName)
}

func Wrapper(columnNumber int) string {
	return convertToTitle(columnNumber)
}
