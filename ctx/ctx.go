package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main() {
	ctx0 := context.Background()

	ctx1, c1 := context.WithCancel(ctx0)
	ctx2, _ := context.WithCancel(ctx1)

	go func() {
		time.Sleep(1 * time.Second)
		c1()
	}()

	<-ctx2.Done()

	fmt.Println(ctx1.Err())

}
