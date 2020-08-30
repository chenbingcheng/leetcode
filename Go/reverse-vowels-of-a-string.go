package Go


func reverseVowels(s string) string {
	left := 0
	right := len(s)-1
	bs := []byte(s)

	for true {
		for left < right {
			if isVowel(bs[left]){
				break
			}
			left++
		}

		for left < right {
			if isVowel(bs[right]){
				break
			}
			right--
		}

		if right < left {
			break
		}

		bs[left], bs[right] = bs[right],bs[left]
		left++
		right--

	}
	return string(bs)
}

func isVowel(s byte)bool {
	return s == byte('A') || s == byte('a') || s == byte('o') || s == byte('O') || s == byte('i') || s == byte('I') || s == byte('e') || s == byte('E') || s == byte('u') || s == byte('U')
}