package aes

import (
    "testing"
    "fmt"
)

func TestAES(t *testing.T){

    message := "hellohowareoyoumynameis"
    key := "testkey"

    encrypt([]byte(message), []byte(key))
    fmt.Print("Hello")
}
