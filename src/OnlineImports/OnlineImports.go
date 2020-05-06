package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)

	// There is ways to lock to a specific version of a github library. But won't include that in the basics.
}
