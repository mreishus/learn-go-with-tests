// Package iteration provides ...
package iteration

func Repeat(input string) string {
	var repeated string
	for index := 0; index < 5; index++ {
		repeated += input
	}
	return repeated
}
