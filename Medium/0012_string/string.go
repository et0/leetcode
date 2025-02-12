package l_m_0012_string

func intToRoman(num int) string {
	result := make([]byte, 0, 30)

	roman := map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}

	for i := 1000; i >= 1; i /= 10 {
		number := num / i % 10
		if i != 1000 && number == 0 {
			continue
		}

		if number == 4 || number == 9 {
			result = append(result, roman[1*i], roman[(number+1)*i])
		} else {
			if number >= 5 {
				result = append(result, roman[5*i])
				number -= 5
			}
			for one := 1; one <= number; one++ {
				result = append(result, roman[1*i])
			}
		}
	}

	return string(result[:])
}

func Wrapper(num int) string {
	return intToRoman(num)
}
