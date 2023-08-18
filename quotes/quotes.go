package quotes

import (
	"fmt"

	"rsc.io/quote"

  "math/rand"
  
)

func Hello(name string) string {
	return fmt.Sprintf("%s, %s", quote.Hello(), name)
}

func RandomGreeting(name string) string {
  data := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
  }
  return fmt.Sprintf(data[rand.Intn(len(data))], name)
}
