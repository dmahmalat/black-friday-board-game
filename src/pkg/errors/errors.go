package errors

import (
	"fmt"
	"log"

	color "github.com/TwiN/go-color"
)

func ErrorCheck(e error, str string) {
	if e != nil {
		err := fmt.Errorf("%s %w", str, e)
		log.Fatalln(color.Ize(color.Red, err.Error()))
	}
}
