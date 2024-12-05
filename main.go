package main

import (
	"fmt"
)

func main() {

	option := InputOption()

	switch option {

	case 1:

		text := InputText()

		// encode the string
		sEnc := EncodeString(text)
		fmt.Println("Encoded string: ", sEnc)

	case 2:

		text := InputText()

		// decode the string
		sDec := DecodeString(text)
		fmt.Println("Decoded String: ", sDec)

	}
}
