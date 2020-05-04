package main

import (
	"fmt"
	"html/template"

	"starbob.com/go/test/morestrings"
)

func main() {
	// out string

	morestrings.ReverseRunes("babak")

	var t = template.New(`<p>Hello {{.}}.</p>`)
	var err = t.Execute(out, template.HTML(`<b>World</b>`))

	fmt.Printf("hello, world\n")
}
