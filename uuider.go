package main

import (
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

func main() {
	identifier := getUUID()

	println(identifier)
	println(strings.Replace(identifier, "-", "", 4))
}

func getUUID() string {
	if len(os.Args) > 1 {
		identifier, err := uuid.Parse(os.Args[1])
		if err != nil {
			log.Fatal("Format must be a valid UUID")
		}
		return identifier.String()
	}

	return uuid.New().String()
}