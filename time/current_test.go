package time

import (
	"fmt"
	"testing"
)

func TestCurrent(t *testing.T) {
	s := GetCurrentTime()
	fmt.Println("s-->", s)
}
