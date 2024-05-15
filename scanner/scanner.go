package scanner

import (
	"bufio"
	"fmt"
	"strings"
)

func Scan() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text string
		n    int
	)
	for scanner.Scan() {
		n++
		/*
			The string text is not modified with the += operator,
			it is created anew for every line. This is a significant difference between strings and []byte slices —
			strings in Go are non-modifiable. If you need to modify a string, use a []byte slice.
		*/
		text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
	}
	fmt.Print(text)
}

// A better imlementation is to use []byte
// func Scan() {
// 	scanner := bufio.NewScanner(strings.NewReader(`one
// two
// three
// four
// `))
// 	var (
// 		text []byte
// 		n    int
// 	)
// 	for scanner.Scan() {
// 		n++
// 		/*
// 			The string text is not modified with the += operator,
// 			it is created anew for every line. This is a significant difference between strings and []byte slices —
// 			strings in Go are non-modifiable. If you need to modify a string, use a []byte slice.
// 		*/
// 		text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
// 	}
// 	os.Stdout.Write(text)
// }
