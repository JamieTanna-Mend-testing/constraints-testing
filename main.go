package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	u := uuid.MustParse(uuid.NewString())
	fmt.Printf("u: %v\n", u)
}
