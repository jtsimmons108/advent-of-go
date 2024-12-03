package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func DayInput(year int, day int) string {
	bytes, err := os.ReadFile(fmt.Sprintf(`inputs/y%d/day%d.txt`, year, day))
	CheckError(err, fmt.Sprintf(`Unable to read input for %d Day %d`, year, day))
	return string(bytes)
}

func ConvertInputToStringSlice(input string, delimeter string) []string {
	return strings.Split(input, delimeter)
}

func ConvertInputToIntSlice(input string, delimeter string) []int {
	strs := ConvertInputToStringSlice(input, delimeter)
	ints := []int{}

	for _, str := range strs {
		n, err := strconv.Atoi(str)
		CheckError(err, fmt.Sprintf(`Unable to convert [%s] to int`, str))
		ints = append(ints, n)
	}

	return ints
}

func ConvertInputToBigIntSlice(input string, delimeter string) []int64 {
	strs := ConvertInputToStringSlice(input, delimeter)
	ints := []int64{}

	for _, str := range strs {
		n, err := strconv.ParseInt(str, 10, 64)
		CheckError(err, fmt.Sprintf(`Unable to convert [%s] to int`, str))
		ints = append(ints, n)
	}

	return ints
}

func ExtractInts(input string, includeNegatives bool) []int {

	pattern := `\d+`

	if includeNegatives {
		pattern = `-?` + pattern
	}

	regex := regexp.MustCompile(pattern)
	matches := regex.FindAll([]byte(input), -1)

	ints := []int{}

	for _, match := range matches {
		n, err := strconv.Atoi(string(match))
		CheckError(err, `Unable to convert string to int`)
		ints = append(ints, n)
	}

	return ints
}

func FindAllOccurrences(input string, pattern string) []string {
	regex := regexp.MustCompile(pattern)
	return regex.FindAllString(input, -1)
}
