package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched1, err := regexp.MatchString(".*贷款", "银行审评贷款")
	fmt.Println(matched1, err)
}
