package goixc

import "regexp"

var notNumbersRE = regexp.MustCompile(`[^0-9]+`)
