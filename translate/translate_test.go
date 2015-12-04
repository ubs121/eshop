package translate

import (
	"log"
	"testing"
)

func testGoogle(t *testing.T) {
	s, err := Translate("hello", "mn")

	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%s", s)
	}

}
