package aes

import (
    "testing"
    "fmt"
)

func TestAES(t *testing.T){

    message := "hellohowareoyoumynameis"
    key := "testkeyhelhalkfd"

    encrypt([]byte(message), []byte(key))
    fmt.Print("Hello")
}
