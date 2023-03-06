package Mystring

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	wordsLen := len(words)
	currentLen := 0

	for left, right := 0, 0; left < wordsLen && right < wordsLen; {
		if currentLen+len(words[right])+right-left > maxWidth {
			line := layout(words, maxWidth, left, right-1, maxWidth-currentLen)
			res = append(res, line)
			left = right
			currentLen = 0
		} else if right == wordsLen-1 {
			line := layout(words, maxWidth, left, right, right-left)
			res = append(res, line)
			break
		} else {
			currentLen += len(words[right])
			right++
		}
	}

	return res
}

func layout(words []string, maxWidth int, left, right, space int) string {
	res := ""
	for i := left; i <= right; i++ {
		res += words[i]

		if right-i == 0 {
			continue
		}

		for j := (space + right - i - 1) / (right - i); j > 0; j-- {
			res += " "
			space--
		}
	}

	for k := len(res); k < maxWidth; k++ {
		res += " "
	}

	return res
}
