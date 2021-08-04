package finder

func FindLongestPrefix(words []string) string {
	letterIndex := 0
	prefix := ""

	outerLoop:
	for true {
		currentLetter := ""
		for i := 0; i < len(words); i++ {
			if i >= len(words) {
				break outerLoop
			}
			word := words[i]
			letter := string(word[letterIndex])
			if currentLetter == "" {
				currentLetter = letter
				continue
			}
			if letter != currentLetter {
				break outerLoop
			}

			if i == len(words) -1 {
				prefix = prefix + currentLetter
				letterIndex++
			}

		}
	}

	return prefix
}
