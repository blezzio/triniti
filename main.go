package main

import (
	"fmt"

	"github.com/blezzio/tini/utils"
)

func main() {
	call()
}

func call() {
	fmt.Printf("%v", utils.Trace(throw1(), "failed %d", 3))
}
func throw1() error {
	return utils.Trace(throw2(), "failed 2")
}

func throw2() error {
	return utils.Trace(throw3(), "failed 1")
}

func throw3() error {
	return fmt.Errorf("root error")
}
