package main

import (
	"fmt"

	"github.com/alanmxll/gpdf/htmlParser"
)

func main() {
	h := htmlParser.New("tmp")

	htmlGenerated, err := h.Create("templates/example.html", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generated HTML:", htmlGenerated)
}
