package main

import (
	"fmt"
)

type AuthProvider interface {
	Authenticate() bool
	Serialize() string
}

type GoogleAuthProvider struct {
	token  string
	secret string
}

type FacebookAuthProvider struct {
	token  string
	secret string
}

func (fbProvider FacebookAuthProvider) Authenticate() bool {
	return true
}

func (g GoogleAuthProvider) Authenticate() bool {
	return false
}

func (fbProvider FacebookAuthProvider) Serialize() string {
	return "Facebook"
}

func (g GoogleAuthProvider) Serialize() string {
	return "Google"
}

func Login(provider AuthProvider) {
	fmt.Println(provider.Serialize())
	fmt.Println("Login successful:", provider.Authenticate())
}

func main() {
  Login(GoogleAuthProvider{
    token: "123",
    secret: "456",
  })

  Login(FacebookAuthProvider{
    token: "123",
    secret: "456",
  })
}
