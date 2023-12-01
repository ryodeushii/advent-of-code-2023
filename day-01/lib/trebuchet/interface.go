package trebuchet

type ITrebuchet interface {
	GetCalibrationNumber(value string) int
	GetCalibrationSum(value []string) int
    getFirstDigit(value string) int
    getLastDigit(value string) int
    reverse(input string) string
}
