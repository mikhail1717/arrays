package util

import (
	"fmt"
)

func ReadNumber() int {
	var a int
	_, _ = fmt.Scanln(&a)
	return a
}
