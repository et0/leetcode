package le0541 // https://leetcode.com/problems/reverse-string-ii/

func reverseStr(s string, k int) string {
	newS := []byte(s)
	length := len(s)

	// смещение
	step := 0
	// правый указатель на K элементов относительно левого или до конца строки
	right := min(0+k-1, length-1)
	for left := 0; left < length; left, right = left+1, right-1 {
		// когда прошли половину подотредка, выполняем смещение и устанавливаем границу
		if left >= right {
			step++

			// проверка, если новая граница выходит за длину строки
			if step*k*2 >= length {
				break
			}

			left = step * k * 2
			right = min(left+k-1, length-1)
		}

		newS[left], newS[right] = newS[right], newS[left]
	}

	return string(newS)
}
