package funkybrain

import (
	"fmt"

	"github.com/lziest/mood"
)

func Mood() string {
	fmt.Print(mood.Status)
	return mood.Status
}
