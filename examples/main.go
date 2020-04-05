package main

import (
	"fmt"
	"github.com/ndthuan/go-vi-wordseg-client/pkg/apiv1"
)

func main() {
	c := apiv1.NewClient("http://segmenterv1:8080")

	tagged, err := c.Tag("Việt Nam tổng tấn công COVID!")

	if err != nil {
		panic(err)
	}

	println("Word-segmented text with tagging:")

	for _, sentence := range tagged.Sentences {
		for _, word := range sentence {
			fmt.Printf("form=%s pos=%s ner=%s dep=%s\n", word.Form, word.Pos, word.Ner, word.Dep)
		}
	}
}
