// Package iteration provides ...
package iteration

func Repeat(input string, count int) string {
	var repeated string
	for index := 0; index < count; index++ {
		repeated += input
	}
	return repeated
}
