package pkg

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	New()
	log.Println("ddd", 3)
	t.Error()
}
