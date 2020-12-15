package hello

import "fmt"

func SayHello(lang string) {
	switch lang {
	case "tamil":
		fmt.Println("வணக்கம் (Vaṇakkam)!")

	default:
		fmt.Println("Hello!")
	}
}
