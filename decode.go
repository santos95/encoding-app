package main

import (
	b64 "encoding/base64"
	"fmt"
)

func DecodeString(text string) string {

	fmt.Println("Decodign string ...")
	decStr, _ := b64.StdEncoding.DecodeString(text)

	return string(decStr)
}
