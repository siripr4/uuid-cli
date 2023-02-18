package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("something went wrong, unable to generate uuid: %s", err)
		return
	}
	fmt.Println(uuid.String())
}
