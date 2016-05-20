package funkybrain

import (
	"testing"
)

func TestMood(t *testing.T) {
	if Mood() != "sad" {
		t.Fatal("Not sad, now I am sad")
	}
}
