package funkybrain

import (
	"testing"
)

func TestMood(t *testing.T) {
	if Mood() != "happy" {
		t.Fatal("Not happy, now I am sad")
	}
}
