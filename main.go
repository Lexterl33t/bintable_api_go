package main

import (
	"log"

	"fmt"
	"github.com/lexterl33t/bintable_api_go/bintable"
)

func main() {
	bin := bintable.New("969f6f251d10568e28887c5fe3287c28e2b6f4b1")

	res, err := bin.Lookup("403244")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
