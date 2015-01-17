package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type T struct{}

func main() {
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	t := T{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		fmt.Println(err)
	}
}
