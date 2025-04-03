package security

import (
	"log"
	"os"
)

var Secret string

func ReadSecret() {
	Secret = os.Getenv("SECRET")
	if Secret == "" {
		log.Fatalln("SECRET is not specified")
	}
}
