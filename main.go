package main

import (
	"fmt"
	"github.com/ZhangSetSail/yamltempl/oty"
)

type People struct {
	Name string
	Age  int
}

func main() {
	test := People{
		Name: "ZhangSetSail",
		Age:  24,
	}
	s, _ := oty.ObjectTOYaml(test)
	fmt.Println(s)
}
