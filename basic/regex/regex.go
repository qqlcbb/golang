package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com
My email1 is ccmouse1@gmail.com
My email2 is ccmouse2@gmail.com
My email3 is ccmouse3@gmail.com
My email3 is ccmouse3@gmail.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
