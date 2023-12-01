package trebuchet

type Trebuchet struct {
	input []string
}

func NewTrebuchet(input []string) Trebuchet {
	return Trebuchet{input}
}

func (t Trebuchet) GetCalibrationNumber(value string) int {
	return t.getFirstDigit(value)*10 + t.getLastDigit(value)
}

func (t Trebuchet) GetCalibrationSum() int {
	sum := 0
	for _, value := range t.input {
		sum += t.GetCalibrationNumber(value)
	}
	return sum

}

func (t Trebuchet) getFirstDigit(value string) int {
	for _, char := range value {
		// if char is a digit - return it
		if char >= '0' && char <= '9' {
			return int(char - '0')
		}
	}
	return 0
}

func (t Trebuchet) getLastDigit(value string) int {
	reversed := t.reverse(value)
	return t.getFirstDigit(reversed)
}

func (t Trebuchet) reverse(input string) string {
	result := string("")
	for _, char := range input {
		result = string(char) + result
	}

	return result
}
