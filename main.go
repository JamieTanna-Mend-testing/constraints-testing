package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/onsi/ginkgo"
)

func main() {
	u := uuid.MustParse(uuid.NewString())
	fmt.Printf("u: %v\n", u)

	fmt.Println(ginkgo.Context("", func() {}))
}
