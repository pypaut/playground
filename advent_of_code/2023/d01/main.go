package main

import (
	"d01/internal/compute"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	Check(err)

	lines := strings.Split(string(content), "\n")
	sum, err := compute.Compute(lines)
	Check(err)

	fmt.Println(sum)
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
