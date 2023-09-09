package quotes

import (
	"fmt"
	"math/rand"

	"rsc.io/quote"
)

func Hello(name string) string {
	return fmt.Sprintf("%s, %s", quote.Hello(), name)
}

func RandomGreeting(name string) string {
	data := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
    "Hello, %v!",
    "Howdy, %v!",
    "Salutations, %v!",
    "Greetings, %v!",
	}
	return fmt.Sprintf(data[rand.Intn(len(data))], name)
}
