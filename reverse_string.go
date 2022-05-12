package reverse_string

func ReverseString(input string) (output string) {
	outputArr := []rune{}
	runeArr := []rune(input)
	runeArrLen := len(runeArr)
	for idx := runeArrLen - 1; idx >= 0; idx-- {
		outputArr = append(outputArr, runeArr[idx])
	}
	output = string(outputArr)
	return output
}
