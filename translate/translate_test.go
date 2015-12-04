package translate

import (
	"fmt"
	"testing"
)

func testGoogle(t *testing.T) {
	s, err := Translate("hello", "mn")
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("%s", s)
	}

}
