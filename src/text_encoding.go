package main

import (
	"fmt"
	"unicode"
	"unsafe"
	"strconv"
)

func textEncoding(text string) []byte {
	var result []byte
	for i := 0; i < len(text); i++ {
		if !unicode.IsDigit(rune(text[i])) {
			result = append(result, byte(text[i]))
		} else {
			number, err := strconv.Atoi(string(text[i]))
			if err == nil {
				result = append(result, byte(number))
			}
		}
	}

	return result
}

func textDecoding(encodedText []int) string {
	var result string
	for _, value := range encodedText {
		result += string(rune(value))
	}

	return result
}

func main() {
	text := "12345678910Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras erat mauris, suscipit a dapibus sed, gravida rutrum risus. Quisque venenatis sit amet massa nec lobortis. Vestibulum diam mi, feugiat eu est in, maximus varius ligula. Ut fringilla turpis quis est dapibus, vel pulvinar tellus dapibus. Duis elementum, libero ac accumsan porta, leo mi tincidunt metus, a molestie nunc lectus et odio. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec tempus quam."
	encodedText := textEncoding(text)
	fmt.Println(encodedText)
	fmt.Println(unsafe.Sizeof(encodedText))

}
