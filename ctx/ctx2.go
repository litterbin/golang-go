package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
)

func main() {
	ctx0 := context.Background()
	ctx1, c1 := context.WithCancel(ctx0)

	var wg sync.WaitGroup

	f := func() {
		<-ctx1.Done()
		fmt.Println("Done")
		wg.Done()
	}

	wg.Add(2)
	go f()
	go f()

	c1()
	wg.Wait()
	fmt.Println(ctx1.Err())
}
