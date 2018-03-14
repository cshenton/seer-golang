package seer_test

import (
	"testing"

	"github.com/cshenton/seer-golang/seer"
)

func TestNew(t *testing.T) {
	_, err := seer.New("localhost:8080")
	if err != nil {
		t.Error("unexpected error in New:", err)
	}
}
