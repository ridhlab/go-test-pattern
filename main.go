package main

import (
	"fmt"

	"github.com/ridhlab/go-math-pattern/pattern"
)

func main() {
	num, arr := pattern.CollatzConjecture(999, []int{})
	fmt.Println(num, arr)
}
