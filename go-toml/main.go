package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func main() {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
		return
	}

	buildTree := config.Get("build").(*toml.TomlTree)
	fmt.Printf("build %s", buildTree.Keys())
	fmt.Println()
}
