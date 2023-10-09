package main

import (
	"fmt"

	"github.com/neel004/cryptit/encrypt"

	"github.com/neel004/cryptit/decrypt"
)

func main() {
	original_string := "If you are reading this, then you'r hacker."
	fmt.Println("Original String : ", original_string)
	enc_string := encrypt.Nimbus(original_string)
	fmt.Println("Encrypted String : ", enc_string)
	dec_string := decrypt.Luffy(enc_string)

	fmt.Println("Decrypted String : ", dec_string)

	if original_string != dec_string {
		fmt.Println("Both String don't match up")
	} else {
		fmt.Println("Decoded Successfully")
	}
}
